package geo

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/kevinmichaelchen/api-dispatch/pkg/maps"
	"github.com/kevinmichaelchen/api-dispatch/pkg/maps/distance"
	"github.com/kevinmichaelchen/api-dispatch/pkg/maps/distance/google"
	"github.com/kevinmichaelchen/api-dispatch/pkg/maps/distance/osrm"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	gMaps "googlemaps.github.io/maps"
	"net/http"
	"time"
)

type Service struct {
	googleClient *gMaps.Client
	httpClient   *http.Client
}

func NewService(client *gMaps.Client, httpClient *http.Client) *Service {
	return &Service{
		googleClient: client,
		httpClient:   httpClient,
	}
}

func (s *Service) BetweenPoints(ctx context.Context, in distance.BetweenPointsInput) (*distance.MatrixResponse, error) {
	tr := otel.Tracer("")
	ctx, span := tr.Start(ctx, "BetweenPoints")
	defer span.End()

	err := validate(in)
	if err != nil {
		return nil, err
	}

	if len(in.Origins) > 1 && len(in.Destinations) > 1 {
		return nil, status.Error(codes.Internal, "distance matrix function cannot support multiple origins and multiple destinations")
	}

	var responses []*distance.MatrixResponse
	var batches []distance.BetweenPointsInput
	multipleOrigins := len(in.Origins) > 1
	if len(in.Origins) > 25 || len(in.Destinations) > 25 || len(in.Origins)*len(in.Destinations) > 100 {
		batches = chunkSlice(in, multipleOrigins)
	}

	for _, batch := range batches {
		var res *distance.MatrixResponse
		var err error
		// Use Google Maps if there's an API key available
		if s.googleClient != nil {
			res, err = google.BetweenPoints(ctx, s.googleClient, batch)
		} else {
			// Otherwise we'll back to using Open Source Routing Machine (OSRM)
			res, err = osrm.BetweenPoints(ctx, s.httpClient, batch)
			// Comply with OSRM's Terms of Use (1 request per second)
			time.Sleep(time.Second)
		}
		if err != nil {
			return nil, err
		}
		responses = append(responses, res)
	}
	return merge(responses, multipleOrigins), nil
}

func chunkSlice(in distance.BetweenPointsInput, multipleOrigins bool) []distance.BetweenPointsInput {
	chunkSize := 25
	var slice []maps.LatLng
	if multipleOrigins {
		slice = in.Origins
	} else {
		slice = in.Destinations
	}

	var chunks [][]maps.LatLng
	var out []distance.BetweenPointsInput
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	for _, arr := range chunks {
		bpi := distance.BetweenPointsInput{}
		if multipleOrigins {
			bpi.Origins = arr
			bpi.Destinations = in.Destinations
		} else {
			bpi.Origins = in.Origins
			bpi.Destinations = arr
		}
		out = append(out, bpi)
	}
	return out
}

func merge(in []*distance.MatrixResponse, multipleOrigins bool) *distance.MatrixResponse {
	out := new(distance.MatrixResponse)
	for _, e := range in {
		out.Rows = append(out.Rows, e.Rows...)
		out.OriginAddresses = append(out.OriginAddresses, e.OriginAddresses...)
		out.DestinationAddresses = append(out.DestinationAddresses, e.DestinationAddresses...)
	}
	return out
}

func validate(i distance.BetweenPointsInput) error {
	return validation.ValidateStruct(&i,
		validation.Field(&i.Destinations,
			validation.Required,
			validation.Length(1, 0),
		),
		validation.Field(&i.Origins,
			validation.Required,
			validation.Length(1, 0),
		),
	)
}
