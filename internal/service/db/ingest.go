package db

import (
	"context"
	"fmt"
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"github.com/kevinmichaelchen/api-dispatch/internal/models"
	"github.com/kevinmichaelchen/api-dispatch/internal/service/h3"
	"github.com/kevinmichaelchen/api-dispatch/internal/service/money"
	"github.com/rs/xid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (s *Store) CreateTrips(ctx context.Context, r *v1beta1.CreateTripsRequest) (*v1beta1.CreateTripsResponse, error) {
	for _, tripPB := range r.GetTrips() {
		trip := models.Trip{
			ID:             tripPB.GetId(),
			ScheduledFor:   tripPB.GetScheduledFor().AsTime(),
			ExpectedPay:    money.ConvertMoneyToFloat(tripPB.GetExpectedPayment()),
			Latitude:       tripPB.GetPickupLocation().GetLatitude(),
			Longitude:      tripPB.GetPickupLocation().GetLongitude(),
			R7Cell:         null.StringFrom(h3.GetCell(tripPB.GetPickupLocation(), 7)),
			R8Cell:         null.StringFrom(h3.GetCell(tripPB.GetPickupLocation(), 8)),
			R9Cell:         null.StringFrom(h3.GetCell(tripPB.GetPickupLocation(), 9)),
			R10Cell:        null.StringFrom(h3.GetCell(tripPB.GetPickupLocation(), 10)),
			R7K1Neighbors:  h3.CellNeighbors(tripPB.GetPickupLocation(), 7, 1),
			R8K1Neighbors:  h3.CellNeighbors(tripPB.GetPickupLocation(), 8, 1),
			R8K2Neighbors:  h3.CellNeighbors(tripPB.GetPickupLocation(), 8, 2),
			R9K1Neighbors:  h3.CellNeighbors(tripPB.GetPickupLocation(), 9, 1),
			R9K2Neighbors:  h3.CellNeighbors(tripPB.GetPickupLocation(), 9, 2),
			R10K1Neighbors: h3.CellNeighbors(tripPB.GetPickupLocation(), 10, 1),
			R10K2Neighbors: h3.CellNeighbors(tripPB.GetPickupLocation(), 10, 2),
		}
		err := trip.Upsert(ctx, s.db, true, []string{models.TripColumns.ID}, boil.Infer(), boil.Infer())
		if err != nil {
			return nil, fmt.Errorf("failed to insert location for trip: %s: %v", tripPB.GetId(), err)
		}
		err = cacheTrip(ctx, s.redisClient, trip)
		if err != nil {
			return nil, err
		}
	}
	return &v1beta1.CreateTripsResponse{}, nil
}

func (s *Store) UpdateDriverLocations(ctx context.Context, r *v1beta1.UpdateDriverLocationsRequest) (*v1beta1.UpdateDriverLocationsResponse, error) {
	for _, l := range r.GetLocations() {
		dl := models.DriverLocation{
			ID:             xid.New().String(),
			CreatedAt:      l.GetMostRecentHeartbeat().AsTime(),
			DriverID:       l.GetDriverId(),
			Latitude:       l.GetCurrentLocation().GetLatitude(),
			Longitude:      l.GetCurrentLocation().GetLongitude(),
			R7Cell:         null.StringFrom(h3.GetCell(l.GetCurrentLocation(), 7)),
			R8Cell:         null.StringFrom(h3.GetCell(l.GetCurrentLocation(), 8)),
			R9Cell:         null.StringFrom(h3.GetCell(l.GetCurrentLocation(), 9)),
			R10Cell:        null.StringFrom(h3.GetCell(l.GetCurrentLocation(), 10)),
			R7K1Neighbors:  h3.CellNeighbors(l.GetCurrentLocation(), 7, 1),
			R8K1Neighbors:  h3.CellNeighbors(l.GetCurrentLocation(), 8, 1),
			R8K2Neighbors:  h3.CellNeighbors(l.GetCurrentLocation(), 8, 2),
			R9K1Neighbors:  h3.CellNeighbors(l.GetCurrentLocation(), 9, 1),
			R9K2Neighbors:  h3.CellNeighbors(l.GetCurrentLocation(), 9, 2),
			R10K1Neighbors: h3.CellNeighbors(l.GetCurrentLocation(), 10, 1),
			R10K2Neighbors: h3.CellNeighbors(l.GetCurrentLocation(), 10, 2),
		}
		err := dl.Insert(ctx, s.db, boil.Infer())
		if err != nil {
			return nil, fmt.Errorf("failed to insert location for driver: %s: %v", l.GetDriverId(), err)
		}
	}
	return &v1beta1.UpdateDriverLocationsResponse{}, nil
}
