package service

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"github.com/kevinmichaelchen/api-dispatch/internal/service/money"
	"github.com/kevinmichaelchen/api-dispatch/internal/service/ranking"
	"github.com/kevinmichaelchen/api-dispatch/pkg/maps"
	"github.com/kevinmichaelchen/api-dispatch/pkg/maps/distance"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"math/rand"
	"time"
)

const (
	// TODO probably a way to get the limit programmatically; https://go.dev/blog/protobuf-apiv2
	maxResults = 1000
)

func (s *Service) GetNearestDrivers(
	ctx context.Context,
	req *v1beta1.GetNearestDriversRequest,
) (*v1beta1.GetNearestDriversResponse, error) {
	err := validate(req, req)
	if err != nil {
		return nil, err
	}

	// Query database
	nearby, err := s.dataStore.GetNearbyDriverLocations(ctx, req.GetPickupLocation())
	if err != nil {
		return nil, err
	}

	// Merge results
	results := MergeDrivers(
		req.GetPickupLocation(),
		MergeDriversInput{Drivers: nearby.R7K1, Res: 7, KValue: 1},
		MergeDriversInput{Drivers: nearby.R8K1, Res: 8, KValue: 1},
		MergeDriversInput{Drivers: nearby.R8K2, Res: 8, KValue: 2},
		MergeDriversInput{Drivers: nearby.R9K1, Res: 9, KValue: 1},
		MergeDriversInput{Drivers: nearby.R9K2, Res: 9, KValue: 2},
		MergeDriversInput{Drivers: nearby.R10K1, Res: 10, KValue: 1},
		MergeDriversInput{Drivers: nearby.R10K2, Res: 10, KValue: 2},
	)

	// Apply server-side results limit
	if len(results) > int(maxResults) {
		results = results[:maxResults]
	}

	// The initial sort will be based on H3 resolutions and k-rings
	results = ranking.SortResultsByKRing(results)

	// Enrich results (e.g., with distance/duration info, among other things)
	var driverLocations []*v1beta1.LatLng
	for _, result := range results {
		driverLocations = append(driverLocations, result.GetLocation())
	}
	matrixOut, err := s.enrichNearbyDrivers(ctx, results, driverLocations, req.GetPickupLocation())
	if err != nil {
		return nil, err
	}

	// Final ranking/sorting pass
	results = ranking.RankDrivers(results)

	// Apply client-side limits
	// TODO do not let client exceed server-side max limit
	if len(results) > int(req.GetLimit()) {
		results = results[:req.GetLimit()]
	}

	return &v1beta1.GetNearestDriversResponse{
		Results:       results,
		PickupAddress: matrixOut.DestinationAddresses[0],
	}, nil
}

func (s *Service) GetNearestTrips(
	ctx context.Context,
	req *v1beta1.GetNearestTripsRequest,
) (*v1beta1.GetNearestTripsResponse, error) {

	err := validate(req, req)
	if err != nil {
		return nil, err
	}

	// Query database
	nearby, err := s.dataStore.GetNearbyTrips(ctx, req.GetDriverLocation())
	if err != nil {
		return nil, err
	}

	// Merge results
	results := MergeTrips(
		req.GetDriverLocation(),
		// TODO these should be fed in reverse order
		MergeTripsInput{trips: nearby.R7K1, res: 7, kValue: 1},
		MergeTripsInput{trips: nearby.R8K1, res: 8, kValue: 1},
		MergeTripsInput{trips: nearby.R8K2, res: 8, kValue: 2},
		MergeTripsInput{trips: nearby.R9K1, res: 9, kValue: 1},
		MergeTripsInput{trips: nearby.R9K2, res: 9, kValue: 2},
		MergeTripsInput{trips: nearby.R10K1, res: 10, kValue: 1},
		MergeTripsInput{trips: nearby.R10K2, res: 10, kValue: 2},
	)

	// Apply server-side results limit
	if len(results) > int(maxResults) {
		results = results[:maxResults]
	}

	// The initial sort will be based on H3 resolutions and k-rings
	results = ranking.SortResultsByKRing(results)

	// Enrich results (e.g., with distance/duration info, among other things)
	var pickupLocations []*v1beta1.LatLng
	for _, result := range results {
		pickupLocations = append(pickupLocations, result.GetLocation())
	}
	_, err = s.enrichNearbyTrips(ctx, results, req.GetDriverLocation(), pickupLocations)
	if err != nil {
		return nil, err
	}

	// Final ranking/sorting pass
	results = ranking.RankTrips(results)

	// Apply client-side limits
	// TODO do not let client exceed server-side max limit
	if len(results) > int(req.GetLimit()) {
		results = results[:req.GetLimit()]
	}

	return &v1beta1.GetNearestTripsResponse{
		Results: results,
	}, nil
}

func (s *Service) enrichNearbyDrivers(
	ctx context.Context,
	results []*v1beta1.SearchResult,
	driverLocations []*v1beta1.LatLng,
	pickupLocation *v1beta1.LatLng,
) (*distance.MatrixResponse, error) {
	logger := ctxzap.Extract(ctx)

	out, err := s.distanceSvc.BetweenPoints(ctx, distance.BetweenPointsInput{
		// the driver location(s) is/are always the origin(s)
		Origins:      toLatLngs(driverLocations),
		Destinations: toLatLngs([]*v1beta1.LatLng{pickupLocation}),
	})
	if err != nil {
		return nil, err
	}

	for i, row := range out.Rows {
		for _, elem := range row.Elements {
			logger.Info("Got Distance Matrix element", zap.Any("elem", elem))
			results[i].Duration = durationpb.New(elem.Duration)
			results[i].DistanceMeters = float64(elem.Distance)
			results[i].Address = out.OriginAddresses[i]
		}
	}

	return out, nil
}

func (s *Service) enrichNearbyTrips(
	ctx context.Context,
	results []*v1beta1.SearchResult,
	driverLocation *v1beta1.LatLng,
	pickupLocations []*v1beta1.LatLng,
) (*distance.MatrixResponse, error) {
	logger := ctxzap.Extract(ctx)

	out, err := s.distanceSvc.BetweenPoints(ctx, distance.BetweenPointsInput{
		// the driver location(s) is/are always the origin(s)
		Origins:      toLatLngs([]*v1beta1.LatLng{driverLocation}),
		Destinations: toLatLngs(pickupLocations),
	})
	if err != nil {
		return nil, err
	}

	for idx := range results {
		e := results[idx]
		t := e.GetTrip()
		t.ScheduledFor = timestamppb.New(randomTime())
		t.ExpectedPayment = randomMoney()
	}

	for _, row := range out.Rows {
		for i, elem := range row.Elements {
			logger.Info("Got Distance Matrix element", zap.Any("elem", elem))
			results[i].Duration = durationpb.New(elem.Duration)
			results[i].DistanceMeters = float64(elem.Distance)
			results[i].Address = out.DestinationAddresses[i]
		}
	}

	return out, nil
}

func randomTime() time.Time {
	minutes := time.Duration(rand.Intn(20)) * time.Minute
	seconds := time.Duration(rand.Intn(60)) * time.Second
	return time.Now().Add(minutes + seconds)
}

func randomMoney() *v1beta1.Money {
	randomUnits := 4 + rand.Intn(25)
	randomCents := rand.Intn(100)
	f := float64(randomUnits) + (float64(randomCents) / float64(100))
	return money.ConvertFloatToMoney(f)
}

func toLatLngs(in []*v1beta1.LatLng) []maps.LatLng {
	var out []maps.LatLng
	for _, e := range in {
		out = append(out, maps.LatLng{
			Lat: e.GetLatitude(),
			Lng: e.GetLongitude(),
		})
	}
	return out
}
