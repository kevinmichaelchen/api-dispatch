package ranking

import (
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"sort"
	"time"
)

func RankTrips(in []*v1beta1.SearchResult) []*v1beta1.SearchResult {
	results := CopySearchResults(in)

	sort.SliceStable(results, func(i, j int) bool {
		a := results[i]
		b := results[j]
		return score(a) > score(b)
	})

	return results
}

func score(in *v1beta1.SearchResult) float64 {
	//timeUntilA := time.Until(a.GetTrip().GetScheduledFor().AsTime())
	//timeUntilB := time.Until(b.GetTrip().GetScheduledFor().AsTime())
	// descending payment
	//return a.GetTrip().GetExpectedPayment() > b.GetTrip().GetExpectedPayment()
	// ascending duration
	u := time.Until(in.GetTrip().GetScheduledFor().AsTime())
	p := in.GetTrip().GetExpectedPayment()
	d := in.GetDuration().AsDuration()
	return 100 + float64(moneyToCents(p)) - float64(u/time.Minute) - float64(d/time.Minute)
}

func nanosToCents(nanos int32) int32 {
	return nanos / 10_000_000
}

func moneyToCents(in *v1beta1.Money) int32 {
	return int32(in.GetUnits()*100) + int32(in.GetNanos()/10_000_000)
}
