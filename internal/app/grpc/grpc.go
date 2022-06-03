package grpc

import (
	"context"
	"fmt"
	"github.com/bufbuild/connect-go"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1/v1beta1connect"
	"github.com/kevinmichaelchen/api-dispatch/internal/service"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
)

var Module = fx.Module("grpc",
	fx.Provide(NewGRPCServer, NewConnectGoServer),
	fx.Invoke(Register),
)

func Register(
	logger *zap.Logger,
	svc *service.Service,
	connectSvc *service.ConnectWrapper,
	server *grpc.Server,
	mux *http.ServeMux,
) {
	// Register our gRPC server
	v1beta1.RegisterDispatchServiceServer(server, svc)
	grpc_health_v1.RegisterHealthServer(server, svc)
	reflection.Register(server)

	// Register our Connect-Go server
	path, handler := v1beta1connect.NewDispatchServiceHandler(
		connectSvc,
		// TODO interceptors
		connect.WithInterceptors(getUnaryInterceptorsForConnect(logger)...),
	)
	mux.Handle(path, handler)
}

func NewConnectGoServer(lc fx.Lifecycle) *http.ServeMux {
	mux := http.NewServeMux()
	// TODO make configurable
	address := fmt.Sprintf(":%d", 8081)
	srv := &http.Server{
		Addr: address,
		// Use h2c so we can serve HTTP/2 without TLS.
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return srv.ListenAndServe()
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return mux
}

func NewGRPCServer(lc fx.Lifecycle, logger *zap.Logger) (*grpc.Server, error) {
	// TODO configure options here
	//var opts grpc.ServerOption
	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			getUnaryInterceptors(logger)...,
		),
		grpc.ChainStreamInterceptor(
			otelgrpc.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(logger),
		),
	)
	lc.Append(fx.Hook{
		// To mitigate the impact of deadlocks in application startup and
		// shutdown, Fx imposes a time limit on OnStart and OnStop hooks. By
		// default, hooks have a total of 15 seconds to complete. Timeouts are
		// passed via Go's usual context.Context.
		OnStart: func(context.Context) error {
			logger.Info("Starting gRPC server.")
			// TODO make configurable
			address := fmt.Sprintf(":%d", 8080)
			lis, err := net.Listen("tcp", address)
			if err != nil {
				logger.Sugar().Fatal("Failed to listen on address \"%s\"", address, zap.Error(err))
			}
			go s.Serve(lis)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Stopping gRPC server.")
			// GracefulStop stops the gRPC server gracefully. It stops the server from
			// accepting new connections and RPCs and blocks until all the pending RPCs are
			// finished.
			s.GracefulStop()
			return nil
		},
	})
	return s, nil
}
