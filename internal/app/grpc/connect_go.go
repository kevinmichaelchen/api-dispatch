package grpc

import (
	"context"
	"fmt"
	"github.com/bufbuild/connect-go"
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1/v1beta1connect"
	"github.com/kevinmichaelchen/api-dispatch/internal/service"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"net/http"
)

func RegisterConnectGoServer(
	logger *zap.Logger,
	connectSvc *service.ConnectWrapper,
	mux *http.ServeMux,
) {
	// Register our Connect-Go server
	path, handler := v1beta1connect.NewDispatchServiceHandler(
		connectSvc,
		connect.WithInterceptors(getUnaryInterceptorsForConnect(logger)...),
	)
	mux.Handle(path, handler)
}

func NewConnectWrapper(s *service.Service) *service.ConnectWrapper {
	return service.NewConnectWrapper(s)
}

func NewConnectGoServer(lc fx.Lifecycle, logger *zap.Logger) *http.ServeMux {
	mux := http.NewServeMux()
	// TODO make configurable
	address := fmt.Sprintf("localhost:%d", 8081)
	srv := &http.Server{
		Addr: address,
		// Use h2c so we can serve HTTP/2 without TLS.
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// In production, we'd want to separate the Listen and Serve phases for
			// better error-handling.
			go func() {
				err := srv.ListenAndServe()
				if err != nil {
					logger.Error("connect-go ListenAndServe failed", zap.Error(err))
				}
			}()
			logger.Sugar().Infof("Listing for connect-go on: %s", address)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return mux
}
