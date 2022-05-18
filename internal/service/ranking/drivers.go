package ranking

import (
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"sort"
)

func RankDrivers(in []*v1beta1.SearchResult) []*v1beta1.SearchResult {
	results := CopySearchResults(in)

	// TODO factor in driver seniority (whether they've done a lot of trips)

	// Sort by duration
	sort.SliceStable(results, func(i, j int) bool {
		a := results[i]
		b := results[j]
		return a.GetDuration().AsDuration() < b.GetDuration().AsDuration()
	})

	return results
}
