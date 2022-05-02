package main

import (
	"context"
	"fmt"
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "dispatch",
	Short: "Dispatch is a tool to make gRPC requests",
	Long: `Dispatch is a tool to make gRPC requests built with
                love by The Drivers Coop.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("yay", args)
	},
}

var latitude, longitude float64

func init() {
	rootCmd.AddCommand(dispatchCmd)
	dispatchCmd.Flags().Float64VarP(&latitude, "latitude", "", 40.7110694, "Latitude of pickup location")
	dispatchCmd.Flags().Float64VarP(&longitude, "longitude", "", -73.9514453, "Longitude of pickup location")
}

var dispatchCmd = &cobra.Command{
	Use:   "dispatch",
	Short: "Hits the DispatchService/Dispatch endpoint",
	Long:  `Hits the DispatchService/Dispatch endpoint`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("lat =", latitude)
		fmt.Println("lng =", longitude)
		r := &v1beta1.DispatchRequest{Location: &v1beta1.LatLng{
			Latitude:  latitude,
			Longitude: longitude,
		}}
		conn, err := grpc.Dial("localhost:8080",
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			log.Fatalf("failed to dial gRPC connection: %v", err)
		}
		client := v1beta1.NewDispatchServiceClient(conn)
		res, err := client.Dispatch(context.Background(), r)
		if err != nil {
			log.Fatalf("gRPC request failed: %v", err)
		}
		for _, r := range res.GetResults() {
			log.Printf("Found neighboring driver %s at resolution %d\n", r.GetDriverId(), r.GetResolution())
		}
		b, err := protojson.MarshalOptions{
			Multiline: true,
			Indent:    "  ",
		}.Marshal(r)
		if err != nil {
			log.Fatalf("Failed to marshal response: %v", err)
		}
		log.Println(string(b))
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
