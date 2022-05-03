package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io/ioutil"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "dispatch",
	Short: "Dispatch is a tool to make gRPC requests",
	Long: `Dispatch is a tool to make gRPC requests built with
                love by The Drivers Coop.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var conn *grpc.ClientConn
var latitude, longitude float64
var limit int32
var ingestPath string

func init() {
	rootCmd.AddCommand(dispatchCmd)
	rootCmd.AddCommand(ingestCmd)

	dispatchCmd.Flags().Int32VarP(&limit, "limit", "", 10, "Max number of results to return")
	dispatchCmd.Flags().Float64VarP(&latitude, "latitude", "", 40.7110694, "Latitude of pickup location")
	dispatchCmd.Flags().Float64VarP(&longitude, "longitude", "", -73.9514453, "Longitude of pickup location")

	ingestCmd.Flags().StringVarP(&ingestPath, "file", "", "seed.json", "path of JSON import file")

	log.Println("Initializing gRPC connection...")
	var err error
	conn, err = grpc.Dial("localhost:8080",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to dial gRPC connection: %v", err)
	}
	log.Println("Initialized gRPC connection.")
}

type SeedSchema struct {
	Locations []struct {
		DriverID string `json:"driver_id"`
		LatLng   struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"lat_lng"`
	} `json:"locations"`
}

var ingestCmd = &cobra.Command{
	Use:   "ingest",
	Short: "Hits the DispatchService/Ingest endpoint",
	Long:  `Hits the DispatchService/Ingest endpoint`,
	Run: func(cmd *cobra.Command, args []string) {
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

		var data SeedSchema
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
	},
}

func marshalProto(m proto.Message) (string, error) {
	b, err := protojson.MarshalOptions{
		Multiline: true,
		Indent:    "  ",
	}.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

var dispatchCmd = &cobra.Command{
	Use:   "dispatch",
	Short: "Hits the DispatchService/Dispatch endpoint",
	Long:  `Hits the DispatchService/Dispatch endpoint`,
	Run: func(cmd *cobra.Command, args []string) {
		// Create request
		req := &v1beta1.GetNearestDriversRequest{
			PickupLocation: &v1beta1.LatLng{
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
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
