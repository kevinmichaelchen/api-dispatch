package main

import (
	"context"
	"encoding/json"
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type SeedDriversSchema struct {
	Locations []struct {
		DriverID string `json:"driver_id"`
		LatLng   struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"lat_lng"`
	} `json:"locations"`
}

type SeedTripsSchema struct {
	Trips []struct {
		Id           string    `json:"id"`
		ScheduledFor time.Time `json:"scheduled_for"`
		LatLng       struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"lat_lng"`
		ExpectedPayment float64 `json:"expected_payment"`
	} `json:"trips"`
}

func ingestDrivers(cmd *cobra.Command, args []string) {
	// Open file
	f, err := os.Open(ingestPath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}

	// Read file bytes
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("Failed to read file bytes: %v", err)
	}
	//log.Println(string(b))

	var data SeedDriversSchema
	err = json.Unmarshal(b, &data)
	if err != nil {
		log.Fatalf("Failed to unmarshal bytes: %v", err)
	}

	//log.Println(data)

	// Create request
	var locations []*v1beta1.DriverLocation
	for _, e := range data.Locations {
		locations = append(locations, &v1beta1.DriverLocation{
			DriverId:  e.DriverID,
			Timestamp: timestamppb.Now(),
			LatLng: &v1beta1.LatLng{
				Latitude:  e.LatLng.Latitude,
				Longitude: e.LatLng.Longitude,
			},
		})
	}

	log.Printf("Seeding %d driver locations\n", len(locations))
	req := &v1beta1.UpdateDriverLocationsRequest{Locations: locations}
	s, err := marshalProto(req)
	if err != nil {
		log.Fatalf("Failed to marshal request: %v", err)
	}
	log.Println(s)

	log.Println("connection", conn)

	// Execute request
	client := v1beta1.NewDispatchServiceClient(conn)
	res, err := client.UpdateDriverLocations(context.Background(), req)
	if err != nil {
		log.Fatalf("gRPC request failed: %v", err)
	}

	// Print response
	s, err = marshalProto(res)
	if err != nil {
		log.Fatalf("Failed to marshal response: %v", err)
	}
	log.Println(s)
}

func ingestTrips(cmd *cobra.Command, args []string) {
	// Open file
	f, err := os.Open(ingestPath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}

	// Read file bytes
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("Failed to read file bytes: %v", err)
	}
	//log.Println(string(b))

	var data SeedTripsSchema
	err = json.Unmarshal(b, &data)
	if err != nil {
		log.Fatalf("Failed to unmarshal bytes: %v", err)
	}

	//log.Println(data)

	// Create request
	var trips []*v1beta1.Trip
	for _, e := range data.Trips {
		trips = append(trips, &v1beta1.Trip{
			Id:           e.Id,
			ScheduledFor: timestamppb.New(e.ScheduledFor),
			PickupLocation: &v1beta1.LatLng{
				Latitude:  e.LatLng.Latitude,
				Longitude: e.LatLng.Longitude,
			},
			ExpectedPayment: e.ExpectedPayment,
		})
	}

	log.Printf("Seeding %d driver locations\n", len(trips))
	req := &v1beta1.CreateTripsRequest{Trips: trips}
	s, err := marshalProto(req)
	if err != nil {
		log.Fatalf("Failed to marshal request: %v", err)
	}
	log.Println(s)

	log.Println("connection", conn)

	// Execute request
	client := v1beta1.NewDispatchServiceClient(conn)
	res, err := client.CreateTrips(context.Background(), req)
	if err != nil {
		log.Fatalf("gRPC request failed: %v", err)
	}

	// Print response
	s, err = marshalProto(res)
	if err != nil {
		log.Fatalf("Failed to marshal response: %v", err)
	}
	log.Println(s)
}
