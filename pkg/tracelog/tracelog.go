package tracelog

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric/global"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"log"
)

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		logger := ctxzap.Extract(ctx)
		span := trace.SpanFromContext(ctx)
		traceID := span.SpanContext().TraceID().String()
		newCtx := ctxzap.ToContext(ctx, logger.With(
			zap.String("traceid", traceID),
		))

		resp, err = handler(newCtx, req)

		handleStatusMetrics(newCtx, err)

		return
	}
}

func handleStatusMetrics(ctx context.Context, err error) {
	meter := global.Meter("go.opentelemetry.io/otel/exporters/prometheus")
	counter, err := meter.SyncFloat64().Counter("ex.com.three")
	if err != nil {
		log.Panicf("failed to initialize instrument: %v", err)
	}

	counter.Add(ctx, 1, attribute.KeyValue{
		Key:   semconv.RPCGRPCStatusCodeKey,
		Value: attribute.StringValue(status.Code(err).String()),
	})
}
