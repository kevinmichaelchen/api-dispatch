package service

import (
	"context"
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"github.com/kevinmichaelchen/api-dispatch/internal/models"
	"github.com/volatiletech/sqlboiler/v4/queries"
)

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

	var results []*v1beta1.SearchResult
	var obj models.DriverLocationSlice
	err := queries.Raw(`
SELECT driver_id, latitude, longitude
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
		results = append(results, &v1beta1.SearchResult{
			DriverId: e.DriverID,
			DistanceMiles: distance(r.GetLocation(), &v1beta1.LatLng{
				Latitude:  e.Latitude,
				Longitude: e.Longitude,
			}),
			// TODO do N separate queries, one for each resolution
			// this field should be the highest in which driver appears as a neighbor
			Resolution: 7,
		})
	}
	return &v1beta1.DispatchResponse{
		Results: results,
	}, nil
}
