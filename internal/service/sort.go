package service

import (
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"google.golang.org/protobuf/proto"
	"sort"
)

func sortResultsByKRing(in []*v1beta1.SearchResult) []*v1beta1.SearchResult {
	results := copySearchResults(in)

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

func sortDrivers(in []*v1beta1.SearchResult) []*v1beta1.SearchResult {
	results := copySearchResults(in)

	// Sort by duration
	sort.SliceStable(results, func(i, j int) bool {
		a := results[i]
		b := results[j]
		return a.GetDuration().AsDuration() < b.GetDuration().AsDuration()
	})

	return results
}

func sortTrips(in []*v1beta1.SearchResult) []*v1beta1.SearchResult {
	results := copySearchResults(in)

	sort.SliceStable(results, func(i, j int) bool {
		a := results[i]
		b := results[j]
		// TODO factor in trip start time, payment
		//timeUntilA := time.Until(a.GetTrip().GetScheduledFor().AsTime())
		//timeUntilB := time.Until(b.GetTrip().GetScheduledFor().AsTime())
		// descending payment
		//return a.GetTrip().GetExpectedPayment() > b.GetTrip().GetExpectedPayment()
		// ascending duration
		return a.GetDuration().AsDuration() < b.GetDuration().AsDuration()
	})

	return results
}

func copySearchResults(results []*v1beta1.SearchResult) []*v1beta1.SearchResult {
	var out []*v1beta1.SearchResult
	for _, e := range results {
		out = append(out, proto.Clone(e).(*v1beta1.SearchResult))
	}
	return out
}
