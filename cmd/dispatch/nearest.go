package main

import (
	"context"
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"github.com/spf13/cobra"
	"google.golang.org/genproto/googleapis/type/latlng"
	"log"
)

func getNearestDrivers(cmd *cobra.Command, args []string) {
	// Create request
	req := &v1beta1.GetNearestDriversRequest{
		PickupLocation: &latlng.LatLng{
			Latitude:  latitude,
			Longitude: longitude,
		},
		Limit: limit,
	}

	// Execute request
	client := v1beta1.NewDispatchServiceClient(conn)
	res, err := client.GetNearestDrivers(context.Background(), req)
	if err != nil {
		log.Fatalf("gRPC request failed: %v", err)
	}

	// Print response
	s, err := marshalProto(res)
	if err != nil {
		log.Fatalf("Failed to marshal response: %v", err)
	}
	log.Println(s)
}

func getNearestTrips(cmd *cobra.Command, args []string) {
	// Create request
	req := &v1beta1.GetNearestTripsRequest{
		Driver: &v1beta1.GetNearestTripsRequest_DriverLocation{
			DriverLocation: &latlng.LatLng{
				Latitude:  latitude,
				Longitude: longitude,
			},
		},
		Limit: limit,
	}

	// Execute request
	client := v1beta1.NewDispatchServiceClient(conn)
	res, err := client.GetNearestTrips(context.Background(), req)
	if err != nil {
		log.Fatalf("gRPC request failed: %v", err)
	}

	// Print response
	s, err := marshalProto(res)
	if err != nil {
		log.Fatalf("Failed to marshal response: %v", err)
	}
	log.Println(s)
}
