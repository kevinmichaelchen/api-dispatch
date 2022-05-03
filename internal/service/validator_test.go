package service

import (
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_validateDispatchRequest(t *testing.T) {
	buildValid := func() *v1beta1.DispatchRequest {
		return &v1beta1.DispatchRequest{
			Location: &v1beta1.LatLng{
				Latitude:  40.2,
				Longitude: -73.3,
			},
			Limit: 5,
		}
	}
	cases := map[string]struct {
		build  func() *v1beta1.DispatchRequest
		expect func(t *testing.T, err error)
	}{
		"Valid": {
			build: buildValid,
			expect: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		//"Negative limit": {
		//	build: func() *v1beta1.DispatchRequest {
		//		p := buildValid()
		//		p.Limit = -5
		//		return p
		//	},
		//	expect: func(t *testing.T, err error) {
		//		require.EqualError(t, err, `foobar`)
		//	},
		//},
	}
	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			err := validateDispatchRequest(tc.build())
			tc.expect(t, err)
		})
	}
}
