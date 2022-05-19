package service

import (
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"github.com/kevinmichaelchen/api-dispatch/internal/models"
	"github.com/kevinmichaelchen/api-dispatch/internal/service/h3"
	"github.com/kevinmichaelchen/api-dispatch/internal/service/money"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type MergeDriversInput struct {
	Drivers models.DriverLocationSlice
	Res     int
	KValue  int
}

func MergeDrivers(location *v1beta1.LatLng, in ...MergeDriversInput) []*v1beta1.SearchResult {
	cache := make(map[string]*v1beta1.SearchResult)
	for _, mi := range in {
		for _, dl := range mi.Drivers {
			extant, exists := cache[dl.DriverID]
			// if we've already recorded the driver appearing in a
			// higher-resolution neighbor, skip
			if exists {
				// prefer higher-res immediate neighbors
				if extant.GetResolution() > int32(mi.Res) {
					continue
				}
				// if they're the same res, prefer those in 1-ring over the
				// 2-ring
				if extant.GetResolution() == int32(mi.Res) {
					if extant.GetKValue() >= int32(mi.KValue) {
						continue
					}
				}
			}
			latLng := &v1beta1.LatLng{
				Latitude:  dl.Latitude,
				Longitude: dl.Longitude,
			}
			cache[dl.DriverID] = &v1beta1.SearchResult{
				Payload: &v1beta1.SearchResult_Driver{
					Driver: &v1beta1.DriverLocation{
						DriverId:            dl.DriverID,
						MostRecentHeartbeat: timestamppb.New(dl.CreatedAt),
						CurrentLocation:     latLng,
					},
				},
				DistanceMeters: h3.PointDistance(location, latLng),
				Location:       latLng,
				Resolution:     int32(mi.Res),
				KValue:         int32(mi.KValue),
			}
		}
	}
	var out []*v1beta1.SearchResult
	for _, e := range cache {
		out = append(out, e)
	}
	return out
}

type MergeTripsInput struct {
	trips  models.TripSlice
	res    int
	kValue int
}

func MergeTrips(location *v1beta1.LatLng, in ...MergeTripsInput) []*v1beta1.SearchResult {
	cache := make(map[string]*v1beta1.SearchResult)
	for _, mi := range in {
		for _, e := range mi.trips {
			extant, exists := cache[e.ID]
			// if we've already recorded the driver appearing in a
			// higher-resolution neighbor, skip
			if exists {
				// prefer higher-res immediate neighbors
				if extant.GetResolution() > int32(mi.res) {
					continue
				}
				// if they're the same res, prefer those in 1-ring over the
				// 2-ring
				if extant.GetResolution() == int32(mi.res) {
					if extant.GetKValue() >= int32(mi.kValue) {
						continue
					}
				}
			}
			latLng := &v1beta1.LatLng{
				Latitude:  e.Latitude,
				Longitude: e.Longitude,
			}
			cache[e.ID] = &v1beta1.SearchResult{
				Payload: &v1beta1.SearchResult_Trip{
					Trip: &v1beta1.Trip{
						Id:              e.ID,
						PickupLocation:  latLng,
						ScheduledFor:    timestamppb.New(e.ScheduledFor),
						ExpectedPayment: money.ConvertFloatToMoney(e.ExpectedPay),
					},
				},
				DistanceMeters: h3.PointDistance(location, latLng),
				Location:       latLng,
				Resolution:     int32(mi.res),
				KValue:         int32(mi.kValue),
			}
		}
	}
	var out []*v1beta1.SearchResult
	for _, e := range cache {
		out = append(out, e)
	}
	return out
}
