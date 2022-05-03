package service

import (
	"context"
	"fmt"
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"github.com/kevinmichaelchen/api-dispatch/internal/models"
	"github.com/rs/xid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (s *Service) CreateTrips(ctx context.Context, r *v1beta1.CreateTripsRequest) (*v1beta1.CreateTripsResponse, error) {
	for _, tripPB := range r.GetTrips() {
		trip := models.Trip{
			ID:             tripPB.GetId(),
			ScheduledFor:   tripPB.GetScheduledFor().AsTime(),
			ExpectedPay:    tripPB.GetExpectedPayment(),
			Latitude:       tripPB.GetPickupLocation().GetLatitude(),
			Longitude:      tripPB.GetPickupLocation().GetLongitude(),
			R7Cell:         null.StringFrom(getCell(tripPB.GetPickupLocation(), 7)),
			R8Cell:         null.StringFrom(getCell(tripPB.GetPickupLocation(), 8)),
			R9Cell:         null.StringFrom(getCell(tripPB.GetPickupLocation(), 9)),
			R10Cell:        null.StringFrom(getCell(tripPB.GetPickupLocation(), 10)),
			R7K1Neighbors:  cellNeighbors(tripPB.GetPickupLocation(), 7, 1),
			R8K1Neighbors:  cellNeighbors(tripPB.GetPickupLocation(), 8, 1),
			R8K2Neighbors:  cellNeighbors(tripPB.GetPickupLocation(), 8, 2),
			R9K1Neighbors:  cellNeighbors(tripPB.GetPickupLocation(), 9, 1),
			R9K2Neighbors:  cellNeighbors(tripPB.GetPickupLocation(), 9, 2),
			R10K1Neighbors: cellNeighbors(tripPB.GetPickupLocation(), 10, 1),
			R10K2Neighbors: cellNeighbors(tripPB.GetPickupLocation(), 10, 2),
		}
		err := trip.Insert(ctx, s.db, boil.Infer())
		if err != nil {
			return nil, fmt.Errorf("failed to insert location for trip: %s: %v", tripPB.GetId(), err)
		}
	}
	return &v1beta1.CreateTripsResponse{}, nil
}

func (s *Service) UpdateDriverLocations(ctx context.Context, r *v1beta1.UpdateDriverLocationsRequest) (*v1beta1.UpdateDriverLocationsResponse, error) {
	for _, l := range r.GetLocations() {
		dl := models.DriverLocation{
			ID:             xid.New().String(),
			CreatedAt:      l.GetMostRecentHeartbeat().AsTime(),
			DriverID:       l.GetDriverId(),
			Latitude:       l.GetCurrentLocation().GetLatitude(),
			Longitude:      l.GetCurrentLocation().GetLongitude(),
			R7Cell:         null.StringFrom(getCell(l.GetCurrentLocation(), 7)),
			R8Cell:         null.StringFrom(getCell(l.GetCurrentLocation(), 8)),
			R9Cell:         null.StringFrom(getCell(l.GetCurrentLocation(), 9)),
			R10Cell:        null.StringFrom(getCell(l.GetCurrentLocation(), 10)),
			R7K1Neighbors:  cellNeighbors(l.GetCurrentLocation(), 7, 1),
			R8K1Neighbors:  cellNeighbors(l.GetCurrentLocation(), 8, 1),
			R8K2Neighbors:  cellNeighbors(l.GetCurrentLocation(), 8, 2),
			R9K1Neighbors:  cellNeighbors(l.GetCurrentLocation(), 9, 1),
			R9K2Neighbors:  cellNeighbors(l.GetCurrentLocation(), 9, 2),
			R10K1Neighbors: cellNeighbors(l.GetCurrentLocation(), 10, 1),
			R10K2Neighbors: cellNeighbors(l.GetCurrentLocation(), 10, 2),
		}
		err := dl.Insert(ctx, s.db, boil.Infer())
		if err != nil {
			return nil, fmt.Errorf("failed to insert location for driver: %s: %v", l.GetDriverId(), err)
		}
	}
	return &v1beta1.UpdateDriverLocationsResponse{}, nil
}
