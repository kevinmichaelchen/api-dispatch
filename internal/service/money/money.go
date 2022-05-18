package money

import (
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"math"
	"math/big"
)

const (
	nanosInCent = 10_000_000
)

// ConvertFloatToMoney converts a float like 3.14 to
// &v1beta1.Money{Units: 3, Nanos: 14*10_000_000}
func ConvertFloatToMoney(f float64) *v1beta1.Money {
	ff := big.NewFloat(f)
	bigCents := ff.Sub(ff, big.NewFloat(math.Floor(f)))
	cents, _ := bigCents.Float64()
	return &v1beta1.Money{
		CurrencyCode: "USD",
		Units:        int64(math.Floor(f)),
		Nanos:        int32(cents * 100 * nanosInCent),
	}
}

// ConvertMoneyToFloat converts a Money object like
// &v1beta1.Money{Units: 3, Nanos: 14*10_000_000}
// to float64(3.14)
func ConvertMoneyToFloat(in *v1beta1.Money) float64 {
	ff := big.NewFloat(float64(in.GetUnits()))
	cents := big.NewFloat(float64(in.GetNanos() / nanosInCent))
	add := cents.Quo(cents, big.NewFloat(100))
	ff = ff.Add(ff, add)
	out, _ := ff.Float64()
	return out
}
