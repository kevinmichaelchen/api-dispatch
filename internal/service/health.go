package service

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

func (s *Service) Check(ctx context.Context, in *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	//_, span := otel.Tracer("dispatch-service").Start(ctx, "workHard",
	//	trace.WithAttributes(attribute.String("extra.key", "extra.value")))
	//defer span.End()

	logger := ctxzap.Extract(ctx)
	logger.Info("Health check endpoint hit 🎉") //zap.String("traceid", span.SpanContext().TraceID().String()),

	return &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}, nil
}

func (s *Service) Watch(in *grpc_health_v1.HealthCheckRequest, _ grpc_health_v1.Health_WatchServer) error {
	// Example of how to register both methods but only implement the Check method.
	return status.Error(codes.Unimplemented, "unimplemented")
}
