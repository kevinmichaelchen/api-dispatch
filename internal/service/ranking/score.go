package ranking

import (
	"math"
	"time"
)

func scoreDurationToPickup(d time.Duration) float64 {
	// The negative multiplier ensures we de-prioritize trips that are far from
	// the driver's current position.
	// TODO 1m15s should rank higher than 1m45s
	return -1 * math.Pow(float64(d/time.Minute), 1.5)
}
