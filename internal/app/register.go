package app

import (
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"github.com/kevinmichaelchen/api-dispatch/internal/service"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric/global"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"net/http"
)

// Register is a typical top-level application function: it takes a generic
// type like *grpc.Server, which typically comes from a third-party library,
// and introduces it to a type that contains our application logic. In this
// case, that introduction consists of registering a gRPC handler. Other typical
// examples include registering HTTP servers and starting queue consumers.
//
// Fx calls these functions invocations, and they're treated differently from
// the constructor functions above. Their arguments are still supplied via
// dependency injection, and they may still return an error to indicate
// failure, but any other return values are ignored.
//
// Unlike constructors, invocations are called eagerly. See the main function
// below for details.
func Register(
	server *grpc.Server,
	svc *service.Service,
	tp *tracesdk.TracerProvider,
	exporter *prometheus.Exporter,
	mux *http.ServeMux,
) {
	v1beta1.RegisterDispatchServiceServer(server, svc)
	grpc_health_v1.RegisterHealthServer(server, svc)
	reflection.Register(server)

	// Register our TracerProvider as the global so any imported
	// instrumentation in the future will default to using it.
	otel.SetTracerProvider(tp)

	// Set global meter provider
	global.SetMeterProvider(exporter.MeterProvider())

	// Register the Prometheus export handler on our Mux HTTP Server.
	mux.HandleFunc("/", exporter.ServeHTTP)
}
