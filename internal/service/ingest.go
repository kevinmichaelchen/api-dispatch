package service

import (
	"context"
	"fmt"
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"github.com/kevinmichaelchen/api-dispatch/internal/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (s *Service) Ingest(ctx context.Context, r *v1beta1.IngestRequest) (*v1beta1.IngestResponse, error) {
	for _, l := range r.GetLocations() {
		dl := models.DriverLocation{
			CreatedAt:      l.GetTimestamp().AsTime(),
			DriverID:       l.GetDriverId(),
			Latitude:       l.GetLatLng().GetLatitude(),
			Longitude:      l.GetLatLng().GetLongitude(),
			R7Cell:         null.StringFrom(getCell(l.GetLatLng(), 7)),
			R8Cell:         null.StringFrom(getCell(l.GetLatLng(), 8)),
			R9Cell:         null.StringFrom(getCell(l.GetLatLng(), 9)),
			R10Cell:        null.StringFrom(getCell(l.GetLatLng(), 10)),
			R7K1Neighbors:  cellNeighbors(l.GetLatLng(), 7),
			R8K1Neighbors:  cellNeighbors(l.GetLatLng(), 8),
			R9K1Neighbors:  cellNeighbors(l.GetLatLng(), 9),
			R10K1Neighbors: cellNeighbors(l.GetLatLng(), 10),
		}
		err := dl.Insert(ctx, s.db, boil.Infer())
		if err != nil {
			return nil, fmt.Errorf("failed to insert location for driver: %s", l.GetDriverId())
		}
	}
	return &v1beta1.IngestResponse{}, nil
}
