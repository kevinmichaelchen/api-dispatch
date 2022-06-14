package service

import (
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"github.com/stretchr/testify/require"
	"google.golang.org/genproto/googleapis/type/latlng"
	"testing"
)

func Test_validateGetNearestDriversRequest(t *testing.T) {
	buildValid := func() *v1beta1.GetNearestDriversRequest {
		return &v1beta1.GetNearestDriversRequest{
			PickupLocation: &latlng.LatLng{
				Latitude:  40.2,
				Longitude: -73.3,
			},
			Limit: 5,
		}
	}
	cases := map[string]struct {
		build  func() *v1beta1.GetNearestDriversRequest
		expect func(t *testing.T, err error)
	}{
		"Valid": {
			build: buildValid,
			expect: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		"Negative limit": {
			build: func() *v1beta1.GetNearestDriversRequest {
				p := buildValid()
				p.Limit = -5
				return p
			},
			expect: func(t *testing.T, err error) {
				require.EqualError(t, err, `invalid GetNearestDriversRequest.Limit: value must be inside range (0, 1000]`)
			},
		},
		"Excessive limit": {
			build: func() *v1beta1.GetNearestDriversRequest {
				p := buildValid()
				p.Limit = 1001
				return p
			},
			expect: func(t *testing.T, err error) {
				require.EqualError(t, err, `invalid GetNearestDriversRequest.Limit: value must be inside range (0, 1000]`)
			},
		},
		"Latitude too low": {
			build: func() *v1beta1.GetNearestDriversRequest {
				p := buildValid()
				p.PickupLocation.Latitude = -91
				return p
			},
			expect: func(t *testing.T, err error) {
				require.EqualError(t, err, `invalid GetNearestDriversRequest.PickupLocation: embedded message failed validation | caused by: invalid LatLng.Latitude: value must be inside range [-90, 90]`)
			},
		},
		"Latitude too high": {
			build: func() *v1beta1.GetNearestDriversRequest {
				p := buildValid()
				p.PickupLocation.Latitude = 91
				return p
			},
			expect: func(t *testing.T, err error) {
				require.EqualError(t, err, `invalid GetNearestDriversRequest.PickupLocation: embedded message failed validation | caused by: invalid LatLng.Latitude: value must be inside range [-90, 90]`)
			},
		},
		"Longitude too low": {
			build: func() *v1beta1.GetNearestDriversRequest {
				p := buildValid()
				p.PickupLocation.Longitude = -181
				return p
			},
			expect: func(t *testing.T, err error) {
				require.EqualError(t, err, `invalid GetNearestDriversRequest.PickupLocation: embedded message failed validation | caused by: invalid LatLng.Longitude: value must be inside range [-180, 180]`)
			},
		},
		"Longitude too high": {
			build: func() *v1beta1.GetNearestDriversRequest {
				p := buildValid()
				p.PickupLocation.Longitude = 181
				return p
			},
			expect: func(t *testing.T, err error) {
				require.EqualError(t, err, `invalid GetNearestDriversRequest.PickupLocation: embedded message failed validation | caused by: invalid LatLng.Longitude: value must be inside range [-180, 180]`)
			},
		},
	}
	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			r := tc.build()
			err := r.Validate()
			tc.expect(t, err)
		})
	}
}
