package money

import (
	"github.com/google/go-cmp/cmp"
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/testing/protocmp"
	"testing"
)

func TestConvertFloatToMoney(t *testing.T) {
	tests := map[string]struct {
		f        float64
		expected *v1beta1.Money
	}{
		"$3.14": {
			f: 3.14,
			expected: &v1beta1.Money{
				CurrencyCode: "USD",
				Units:        3,
				Nanos:        140_000_000,
			},
		},
	}
	for testName, tc := range tests {
		t.Run(testName, func(t *testing.T) {
			actual := ConvertFloatToMoney(tc.f)
			require.Empty(t, cmp.Diff(actual, tc.expected, protocmp.Transform()))
		})
	}
}

func TestConvertMoneyToFloat(t *testing.T) {
	tests := map[string]struct {
		input    *v1beta1.Money
		expected float64
	}{
		"$3.14": {
			input: &v1beta1.Money{
				CurrencyCode: "USD",
				Units:        3,
				Nanos:        140_000_000,
			},
			expected: 3.14,
		},
	}
	for testName, tc := range tests {
		t.Run(testName, func(t *testing.T) {
			actual := ConvertMoneyToFloat(tc.input)
			require.Equal(t, tc.expected, actual)
		})
	}
}
