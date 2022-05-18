package service

import (
	"context"
	"github.com/kevinmichaelchen/api-dispatch/internal/distance"
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"github.com/kevinmichaelchen/api-dispatch/internal/service/ranking"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/durationpb"
)

const (
	maxResults = 100
)

func (s *Service) GetNearestDrivers(ctx context.Context, req *v1beta1.GetNearestDriversRequest) (*v1beta1.GetNearestDriversResponse, error) {
	err := validateGetNearestDriversRequest(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// Query database
	nearby, err := s.getNearbyDriverLocations(ctx, req.GetPickupLocation())
	if err != nil {
		return nil, err
	}

	// Merge results
	results := mergeDrivers(
		req.GetPickupLocation(),
		mergeDriversInput{drivers: nearby.r7k1Cells, res: 7, kValue: 1},
		mergeDriversInput{drivers: nearby.r8k1Cells, res: 8, kValue: 1},
		mergeDriversInput{drivers: nearby.r8k2Cells, res: 8, kValue: 2},
		mergeDriversInput{drivers: nearby.r9k1Cells, res: 9, kValue: 1},
		mergeDriversInput{drivers: nearby.r9k2Cells, res: 9, kValue: 2},
		mergeDriversInput{drivers: nearby.r10k1Cells, res: 10, kValue: 1},
		mergeDriversInput{drivers: nearby.r10k2Cells, res: 10, kValue: 2},
	)

	// Apply server-side results limit
	if len(results) > int(maxResults) {
		results = results[:maxResults]
	}

	// In the event we have no Google Maps client and are operating in a
	// degraded state, k-ring-sorting is still pretty good.
	results = sortResultsByKRing(results)

	// Enrich results with distance/duration info from Google Maps API
	var driverLocations []*v1beta1.LatLng
	for _, result := range results {
		driverLocations = append(driverLocations, result.GetLocation())
	}
	var pickupAddress string
	if s.distanceSvc != nil {
		out, err := s.distanceSvc.BetweenPoints(ctx, distance.BetweenPointsInput{
			PickupLocations: []*v1beta1.LatLng{req.GetPickupLocation()},
			DriverLocations: driverLocations,
		})
		if err != nil {
			return nil, err
		}
		for i, info := range out.Info {
			results[i].Duration = durationpb.New(info.Duration)
			results[i].DistanceMeters = float64(info.DistanceMeters)
			// the driver is always the origin
			results[i].Address = info.OriginAddress
			pickupAddress = info.DestinationAddress
		}
	}

	// Sort
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

func (s *Service) GetNearestTrips(ctx context.Context, req *v1beta1.GetNearestTripsRequest) (*v1beta1.GetNearestTripsResponse, error) {
	//err := validateGetNearestTripsRequest(req)
	//if err != nil {
	//	return nil, status.Error(codes.InvalidArgument, err.Error())
	//}

	// Query database
	nearby, err := s.getNearbyTrips(ctx, req.GetDriverLocation())
	if err != nil {
		return nil, err
	}

	// Merge results
	results := mergeTrips(
		req.GetDriverLocation(),
		// TODO these should be fed in reverse order
		mergeTripsInput{trips: nearby.r7k1Cells, res: 7, kValue: 1},
		mergeTripsInput{trips: nearby.r8k1Cells, res: 8, kValue: 1},
		mergeTripsInput{trips: nearby.r8k2Cells, res: 8, kValue: 2},
		mergeTripsInput{trips: nearby.r9k1Cells, res: 9, kValue: 1},
		mergeTripsInput{trips: nearby.r9k2Cells, res: 9, kValue: 2},
		mergeTripsInput{trips: nearby.r10k1Cells, res: 10, kValue: 1},
		mergeTripsInput{trips: nearby.r10k2Cells, res: 10, kValue: 2},
	)

	// Apply server-side results limit
	if len(results) > int(maxResults) {
		results = results[:maxResults]
	}

	// In the event we have no Google Maps client and are operating in a
	// degraded state, k-ring-sorting is still pretty good.
	results = sortResultsByKRing(results)

	// Enrich results with distance/duration info from Google Maps API
	var locations []*v1beta1.LatLng
	for _, result := range results {
		locations = append(locations, result.GetLocation())
	}
	if s.distanceSvc != nil {
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

	// Re-sort by duration
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
