package service

import (
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"github.com/kevinmichaelchen/api-dispatch/internal/service/ranking"
	"sort"
)

func sortResultsByKRing(in []*v1beta1.SearchResult) []*v1beta1.SearchResult {
	results := ranking.CopySearchResults(in)

	sort.SliceStable(results, func(i, j int) bool {
		a := results[i]
		b := results[j]
		if a.GetResolution() == b.GetResolution() {
			return a.GetKValue() > b.GetKValue()
		}
		return a.GetResolution() > b.GetResolution()
	})

	return results
}
