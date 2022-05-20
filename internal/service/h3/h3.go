package h3

import (
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"github.com/uber/h3-go"
	"github.com/volatiletech/sqlboiler/v4/types"
)

// GetCell indexes the location at the specified resolution, returning the index
// of the cell containing the location. This buckets the geographic point into
// the H3 grid. See the algorithm description for more information:
// https://h3geo.org/docs/core-library/geoToH3desc
//
// https://h3geo.org/docs/api/indexing#geotoh3
func GetCell(l *v1beta1.LatLng, res int) string {
	i := h3.FromGeo(h3.GeoCoord{
		Latitude:  l.GetLatitude(),
		Longitude: l.GetLongitude(),
	}, res)
	return h3.ToString(i)
}

// CellNeighbors produces indices within k distance of the origin index.
// k-ring 0 is defined as the origin index.
// k-ring 1 is defined as k-ring 0 and all neighboring indices, and so on.
//
// https://h3geo.org/docs/api/traversal#kring
func CellNeighbors(l *v1beta1.LatLng, res int, k int) types.StringArray {
	i := h3.FromGeo(h3.GeoCoord{
		Latitude:  l.GetLatitude(),
		Longitude: l.GetLongitude(),
	}, res)
	indexes := h3.KRing(i, k)
	var out []string
	for _, idx := range indexes {
		out = append(out, h3.ToString(idx))
	}
	return out
}

// PointDistance returns the "great circle" or "haversine" distance between
// pairs of GeoCoord points (lat/lng pairs) in meters.
// https://h3geo.org/docs/api/misc/#pointdistm
func PointDistance(l1, l2 *v1beta1.LatLng) float64 {
	// TODO can't find this on the h3-go SDK
	return 0
}
