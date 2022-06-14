package ranking

import (
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"github.com/stretchr/testify/require"
	"google.golang.org/genproto/googleapis/type/latlng"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"
)

func TestScoreTrip(t *testing.T) {
	tests := map[string]struct {
		timeUntilTripStart    time.Duration
		timeForDriverToArrive time.Duration
		payment               *v1beta1.Money
		expected              float64
	}{
		"OK": {
			timeUntilTripStart:    5 * time.Minute,
			timeForDriverToArrive: time.Minute + (15 * time.Second),
			payment:               newMoney(15, 50),
			expected:              109.5,
		},
	}
	for testName, tc := range tests {
		t.Run(testName, func(t *testing.T) {
			sr := newSearchResult(tc.timeUntilTripStart, tc.timeForDriverToArrive, tc.payment)
			actual := scoreTrip(sr)
			require.Equal(t, tc.expected, actual)
		})
	}

}

func newSearchResult(timeUntilTripStart time.Duration, timeForDriverToArrive time.Duration, payment *v1beta1.Money) *v1beta1.SearchResult {
	return &v1beta1.SearchResult{
		Payload: &v1beta1.SearchResult_Trip{
			Trip: &v1beta1.Trip{
				Id:              "trip",
				ScheduledFor:    timestamppb.New(time.Now().Add(timeUntilTripStart)),
				ExpectedPayment: payment,
			},
		},
		DistanceMeters: 100,
		Duration:       durationpb.New(timeForDriverToArrive),
		Location: &latlng.LatLng{
			Latitude:  10,
			Longitude: 10,
		},
		Address:    "Address",
		Resolution: 8,
		KValue:     2,
	}
}

func newMoney(units int64, cents int32) *v1beta1.Money {
	return &v1beta1.Money{
		CurrencyCode: "USD",
		Units:        units,
		Nanos:        cents * 10_000_000,
	}
}
