package tracelog

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc"
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
		return
	}
}
