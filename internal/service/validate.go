package service

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"google.golang.org/genproto/googleapis/type/latlng"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type Validater interface {
	Validate() error
}

func validateLatLng(value interface{}) error {
	l, ok := value.(*latlng.LatLng)
	if !ok {
		return fmt.Errorf("expected *latlng.LatLng, but received %T", l)
	}
	return validation.ValidateStruct(l,
		validation.Field(&l.Latitude,
			validation.Min(float64(-90)),
			validation.Max(float64(90)),
		),
		validation.Field(&l.Longitude,
			validation.Min(float64(-180)),
			validation.Max(float64(180)),
		),
	)
}

func validateGetNearestDriversRequest(r *v1beta1.GetNearestDriversRequest) error {
	err := r.Validate()
	if err != nil {
		return err
	}
	return validation.Errors{
		"PickupLocation": validation.Validate(
			r.GetPickupLocation(),
			validation.By(validateLatLng),
		),
	}.Filter()
}

func validate(m proto.Message, r Validater) error {
	//name := m.ProtoReflect().Type().Descriptor().Name()
	err := r.Validate()
	if err != nil {
		return status.Error(codes.InvalidArgument, err.Error())
	}
	return nil
}
