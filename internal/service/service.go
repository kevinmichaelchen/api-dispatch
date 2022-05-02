package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"github.com/kevinmichaelchen/api-dispatch/internal/models"
	"github.com/uber/h3-go"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/types"
	"log"
)

type Service struct {
	logger *log.Logger
	db     *sql.DB
}

func NewService(logger *log.Logger, db *sql.DB) *Service {
	return &Service{logger: logger, db: db}
}

func getCell(l *v1beta1.LatLng, res int) string {
	i := h3.FromGeo(h3.GeoCoord{
		Latitude:  l.GetLatitude(),
		Longitude: l.GetLongitude(),
	}, res)
	return h3.ToString(i)
}

func cellNeighbors(l *v1beta1.LatLng, res int) types.StringArray {
	i := h3.FromGeo(h3.GeoCoord{
		Latitude:  l.GetLatitude(),
		Longitude: l.GetLongitude(),
	}, res)
	indexes := h3.KRing(i, 1)
	var out []string
	for _, idx := range indexes {
		out = append(out, h3.ToString(idx))
	}
	return out
}

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

func (s *Service) Dispatch(ctx context.Context, r *v1beta1.DispatchRequest) (*v1beta1.DispatchResponse, error) {
	//r7K1Neighbors := cellNeighbors(r.GetLocation(), 7)
	//r8K1Neighbors := cellNeighbors(r.GetLocation(), 8)
	//r9K1Neighbors := cellNeighbors(r.GetLocation(), 9)
	//r10K1Neighbors := cellNeighbors(r.GetLocation(), 10)

	// Get cells
	cell7 := getCell(r.GetLocation(), 7)
	cell8 := getCell(r.GetLocation(), 8)
	cell9 := getCell(r.GetLocation(), 9)
	cell10 := getCell(r.GetLocation(), 10)

	s.logger.Printf("Received request for trip location: (%v, %v)\n", r.GetLocation().GetLatitude(), r.GetLocation().GetLongitude())
	s.logger.Println("Res 7 Cell =", cell7)
	s.logger.Println("Res 8 Cell =", cell8)
	s.logger.Println("Res 9 Cell =", cell9)
	s.logger.Println("Res 10 Cell =", cell10)

	var driverIDs []string
	var obj models.DriverLocationSlice
	err := queries.Raw(`
SELECT driver_id 
FROM driver_location 
WHERE 
  $1 = ANY (r7_k1_neighbors)
  OR $2 = ANY (r8_k1_neighbors)
  OR $3 = ANY (r9_k1_neighbors)
  OR $4 = ANY (r10_k1_neighbors)
`, cell7, cell8, cell9, cell10).Bind(ctx, s.db, &obj)
	if err != nil {
		return nil, err
	}
	for _, e := range obj {
		driverIDs = append(driverIDs, e.DriverID)
	}
	return &v1beta1.DispatchResponse{
		DriverIds: driverIDs,
	}, nil
}
