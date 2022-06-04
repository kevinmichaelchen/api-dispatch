package db

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"github.com/kevinmichaelchen/api-dispatch/internal/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PageToken struct {
	ID string `json:"id"`
}

func (p PageToken) toRaw() (string, error) {
	b, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

func pageTokenFromRaw(raw string) (*PageToken, error) {
	if raw == "" {
		return &PageToken{}, nil
	}
	b, err := base64.StdEncoding.DecodeString(raw)
	if err != nil {
		return nil, err
	}
	var v PageToken
	err = json.Unmarshal(b, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

func (s *Store) ListDrivers(ctx context.Context, r *v1beta1.ListDriversRequest) (*v1beta1.ListDriversResponse, error) {
	pageToken, err := pageTokenFromRaw(r.GetPageToken())
	if err != nil {
		return nil, err
	}

	mods := []qm.QueryMod{
		qm.Select("*"),
	}

	if pageToken.ID != "" {
		// CURSOR PAGINATION
		mods = append(mods,
			qm.Where(models.DriverLocationColumns.ID+"> ?", pageToken.ID))
	}

	mods = append(mods, qm.OrderBy(models.DriverLocationColumns.ID))

	drivers, err := models.DriverLocations(mods...).All(ctx, s.db)
	if err != nil {
		return nil, err
	}

	pageTokenOut := PageToken{}
	if len(drivers) > 0 {
		lastID := drivers[len(drivers)-1].ID
		pageTokenOut.ID = lastID
	}
	nextPageToken, err := pageTokenOut.toRaw()
	if err != nil {
		return nil, err
	}

	return &v1beta1.ListDriversResponse{
		DriverLocations: driverLocationsToProtos(drivers),
		NextPageToken:   nextPageToken,
	}, nil
}

func driverLocationsToProtos(in models.DriverLocationSlice) []*v1beta1.DriverLocation {
	var out []*v1beta1.DriverLocation
	for _, e := range in {
		out = append(out, driverLocationToProto(e))
	}
	return out
}

func driverLocationToProto(in *models.DriverLocation) *v1beta1.DriverLocation {
	return &v1beta1.DriverLocation{
		Id:                  in.ID,
		DriverId:            in.DriverID,
		MostRecentHeartbeat: timestamppb.New(in.CreatedAt),
		CurrentLocation: &v1beta1.LatLng{
			Latitude:  in.Latitude,
			Longitude: in.Longitude,
		},
	}
}
