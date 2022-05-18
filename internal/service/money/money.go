package money

import "github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"

const (
	nanosInCent = 10_000_000
)

// ConvertFloatToMoney converts a float like 3.14 to
// &v1beta1.Money{Units: 3, Nanos: 14*10_000_000}
func ConvertFloatToMoney(f float64) *v1beta1.Money {
	return &v1beta1.Money{
		CurrencyCode: "USD",
		Units:        int64(f),
		// TODO use math/big; these floating ops aren't reliable
		Nanos: int32(f-float64(int64(f))) * nanosInCent,
	}
}

func ConvertMoneyToFloat(in *v1beta1.Money) float64 {
	return float64(in.GetUnits()) + float64(in.GetNanos()*nanosInCent)
}
