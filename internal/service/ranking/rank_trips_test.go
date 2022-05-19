package ranking

import (
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"
)

func TestScoreTrip(t *testing.T) {
	actual := scoreTrip(&v1beta1.SearchResult{
		Payload: &v1beta1.SearchResult_Trip{
			Trip: &v1beta1.Trip{
				Id:           "trip1",
				ScheduledFor: timestamppb.New(time.Now().Add(5 * time.Minute)),
				ExpectedPayment: &v1beta1.Money{
					CurrencyCode: "USD",
					Units:        15,
					Nanos:        50 * 10_000_000,
				},
			},
		},
		DistanceMeters: 100,
		Duration:       durationpb.New(time.Minute + (15 * time.Second)),
		Location: &v1beta1.LatLng{
			Latitude:  10,
			Longitude: 10,
		},
		Address:    "Address",
		Resolution: 8,
		KValue:     2,
	})
	require.Equal(t, float64(109.5), actual)
}
