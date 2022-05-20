package ranking

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestScoreDurationToPickup(t *testing.T) {
	tests := map[string]struct {
		d        time.Duration
		expected float64
	}{
		"1m15s": {
			d:        time.Minute + (15 * time.Second),
			expected: -1,
		},
		"1m30s": {
			d:        time.Minute + (30 * time.Second),
			expected: -1,
		},
		"1m45s": {
			d:        time.Minute + (45 * time.Second),
			expected: -1,
		},
		"2m0s": {
			d:        2 * time.Minute,
			expected: -2.82842712474619,
		},
	}
	for testName, tc := range tests {
		t.Run(testName, func(t *testing.T) {
			actual := scoreDurationToPickup(tc.d)
			require.Equal(t, tc.expected, actual)
		})
	}
}
