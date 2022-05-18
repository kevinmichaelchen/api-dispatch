package ranking

import (
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"sort"
	"time"
)

func RankDrivers(in []*v1beta1.SearchResult) []*v1beta1.SearchResult {
	results := CopySearchResults(in)

	// Sort by duration
	sort.SliceStable(results, func(i, j int) bool {
		a := results[i]
		b := results[j]
		// Sort results with the highest scores at the top
		return scoreDriver(a) > scoreDriver(b)
	})

	return results
}

func scoreDriver(in *v1beta1.SearchResult) float64 {
	// TODO factor in driver seniority (whether they've done a lot of trips)
	d := in.GetDuration().AsDuration()
	score := 100 - float64(d/time.Minute)
	in.Score = score
	return score
}
