package service

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"github.com/kevinmichaelchen/api-dispatch/internal/service/distance"
	"github.com/kevinmichaelchen/api-dispatch/internal/service/money"
	"github.com/kevinmichaelchen/api-dispatch/internal/service/ranking"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"math/rand"
	"time"
)

const (
	maxResults = 100
)

func (s *Service) GetNearestDrivers(
	ctx context.Context,
	req *v1beta1.GetNearestDriversRequest,
) (*v1beta1.GetNearestDriversResponse, error) {
	logger := ctxzap.Extract(ctx)
	err := validateGetNearestDriversRequest(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	trafficAware := s.distanceSvc != nil

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

	// In the event we have no Google Maps client and are operating in a
	// degraded state, k-ring-sorting is still pretty good.
	results = ranking.SortResultsByKRing(results)

	// Enrich results with distance/duration info from Google Maps API
	var driverLocations []*v1beta1.LatLng
	for _, result := range results {
		driverLocations = append(driverLocations, result.GetLocation())
	}
	var pickupAddress string
	if trafficAware {
		out, err := s.distanceSvc.BetweenPoints(ctx, distance.BetweenPointsInput{
			PickupLocations: []*v1beta1.LatLng{req.GetPickupLocation()},
			DriverLocations: driverLocations,
		})
		if err != nil {
			return nil, err
		}
		for i, info := range out.Info {
			logger.Info("received distance matrix info", zap.Any("info", info))
			results[i].Duration = durationpb.New(info.Duration)
			results[i].DistanceMeters = float64(info.DistanceMeters)
			// the driver is always the origin
			results[i].Address = info.OriginAddress
			pickupAddress = info.DestinationAddress
		}
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
		PickupAddress: pickupAddress,
	}, nil
}

func (s *Service) GetNearestTrips(
	ctx context.Context,
	req *v1beta1.GetNearestTripsRequest,
) (*v1beta1.GetNearestTripsResponse, error) {
	//err := validateGetNearestTripsRequest(req)
	//if err != nil {
	//	return nil, status.Error(codes.InvalidArgument, err.Error())
	//}
	err := req.GetDriverLocation().Validate()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	trafficAware := s.distanceSvc != nil

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

	// In the event we have no Google Maps client and are operating in a
	// degraded state, k-ring-sorting is still pretty good.
	results = ranking.SortResultsByKRing(results)

	// Enrich results with distance/duration info from Google Maps API
	var locations []*v1beta1.LatLng
	for _, result := range results {
		locations = append(locations, result.GetLocation())
	}
	if trafficAware {
		out, err := s.distanceSvc.BetweenPoints(ctx, distance.BetweenPointsInput{
			PickupLocations: locations,
			DriverLocations: []*v1beta1.LatLng{req.GetDriverLocation()},
		})
		if err != nil {
			return nil, err
		}
		for i, info := range out.Info {
			results[i].Duration = durationpb.New(info.Duration)
			results[i].DistanceMeters = float64(info.DistanceMeters)
			// the driver is always the origin, the pickup is the destination
			results[i].Address = info.DestinationAddress
		}
	}

	// Enrich results
	enrichTripsWithFakeData(results)

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

func enrichTripsWithFakeData(in []*v1beta1.SearchResult) {
	for idx := range in {
		e := in[idx]
		t := e.GetTrip()
		t.ScheduledFor = timestamppb.New(randomTime())
		t.ExpectedPayment = randomMoney()
	}
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
