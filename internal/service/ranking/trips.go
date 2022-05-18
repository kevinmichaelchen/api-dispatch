package ranking

import (
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"github.com/kevinmichaelchen/api-dispatch/internal/service/money"
	"sort"
	"time"
)

func RankTrips(in []*v1beta1.SearchResult) []*v1beta1.SearchResult {
	results := CopySearchResults(in)

	sort.SliceStable(results, func(i, j int) bool {
		a := results[i]
		b := results[j]
		// Sort results with the highest scores at the top
		return scoreTrip(a) > scoreTrip(b)
	})

	return results
}

func scoreTrip(in *v1beta1.SearchResult) float64 {
	//timeUntilA := time.Until(a.GetTrip().GetScheduledFor().AsTime())
	//timeUntilB := time.Until(b.GetTrip().GetScheduledFor().AsTime())
	// descending payment
	//return a.GetTrip().GetExpectedPayment() > b.GetTrip().GetExpectedPayment()
	// ascending duration
	u := time.Until(in.GetTrip().GetScheduledFor().AsTime())
	p := in.GetTrip().GetExpectedPayment()

	return 100 +
		float64(money.ConvertMoneyToFloat(p)) -
		float64(u/time.Minute) +
		scoreDurationToPickup(in.GetDuration().AsDuration())
}
