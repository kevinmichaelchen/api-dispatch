package service

import (
	"fmt"
	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
)

func validateGetNearestDriversRequest(r *v1beta1.GetNearestDriversRequest) error {
	return ozzo.Errors{
		"Limit": ozzo.Validate(
			r.GetLimit(),
			ozzo.Min(0),
			ozzo.Max(maxResults),
		),
		//"Location": ozzo.Validate(
		//	r.GetLocation(),
		//	ozzo.Required,
		//	ozzo.By(validateLatLngRuleFunc),
		//),
	}.Filter()
}

func validateLatLngRuleFunc(v interface{}) error {
	p, ok := v.(*v1beta1.LatLng)
	if !ok {
		return fmt.Errorf("expected v1beta1.LatLng but got %T", v)
	}
	return validateLatLng(p)
}

func validateLatLng(r *v1beta1.LatLng) error {
	// Null Island is considered invalid.
	return ozzo.Errors{
		// It must be in the range [-90.0, +90.0].
		"Latitude": ozzo.Validate(
			r.GetLatitude(),
			ozzo.Min(float64(-90)), ozzo.Max(float64(90)),
			ozzo.When(r.GetLongitude() == 0, ozzo.Required),
		),
		//It must be in the range [-180.0, +180.0].
		"Longitude": ozzo.Validate(
			r.GetLongitude(),
			ozzo.Min(float64(-180)), ozzo.Max(float64(180)),
			ozzo.When(r.GetLatitude() == 0, ozzo.Required),
		),
	}
}
