package service

import (
	"database/sql"
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"github.com/uber/h3-go"
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

func distance(l1, l2 *v1beta1.LatLng) float64 {
	// TODO can't find this on the h3-go SDK
	// https://h3geo.org/docs/api/misc/#pointdistm
	return 0
}
