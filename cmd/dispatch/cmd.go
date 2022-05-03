package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "dispatch",
	Short: "Dispatch is a tool to make gRPC requests",
	Long: `Dispatch is a tool to make gRPC requests built with
                love by The Drivers Coop.`,
	Run: func(cmd *cobra.Command, args []string) {},
}

var conn *grpc.ClientConn

var latitude, longitude float64
var limit int32
var ingestPath string

func init() {
	rootCmd.AddCommand(nearestCmd)
	rootCmd.AddCommand(ingestCmd)

	nearestCmd.AddCommand(getNearestTripsCmd)
	nearestCmd.AddCommand(getNearestDriversCmd)

	ingestCmd.AddCommand(ingestDriversCmd)
	ingestCmd.AddCommand(ingestTripsCmd)

	nearestCmd.PersistentFlags().Int32VarP(&limit, "limit", "", 10, "Max number of results to return")
	nearestCmd.PersistentFlags().Float64VarP(&latitude, "latitude", "", 40.7110694, "Latitude of pickup location")
	nearestCmd.PersistentFlags().Float64VarP(&longitude, "longitude", "", -73.9514453, "Longitude of pickup location")

	ingestCmd.PersistentFlags().StringVarP(&ingestPath, "file", "", "seed-drivers.json", "path of JSON import file")

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

var ingestCmd = &cobra.Command{
	Use:   "ingest",
	Short: "Ingest drivers or trips",
	Long:  `Ingest drivers or trips`,
	Run:   func(cmd *cobra.Command, args []string) {},
}

var ingestDriversCmd = &cobra.Command{
	Use:   "drivers",
	Short: "Ingest drivers",
	Long:  `Ingest drivers`,
	Run:   ingestDrivers,
}

var ingestTripsCmd = &cobra.Command{
	Use:   "trips",
	Short: "Ingest trips",
	Long:  `Ingest trips`,
	Run:   ingestTrips,
}

var getNearestDriversCmd = &cobra.Command{
	Use:   "drivers",
	Short: "Get nearest drivers",
	Long:  `Get nearest drivers`,
	Run:   getNearestDrivers,
}

var getNearestTripsCmd = &cobra.Command{
	Use:   "trips",
	Short: "Get nearest trips",
	Long:  `Get nearest trips`,
	Run:   getNearestTrips,
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

var nearestCmd = &cobra.Command{
	Use:   "nearest",
	Short: "Gets nearest trips or drivers",
	Long:  `Gets nearest trips or drivers`,
	Run:   func(cmd *cobra.Command, args []string) {},
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
