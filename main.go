package main

import (
	app "github.com/kevinmichaelchen/api-dispatch/internal/app"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"os"
)

func main() {
	a := fx.New(
		fx.Provide(
			app.NewLogger,
			app.NewDatabase,
			app.NewGRPCServer,
			app.NewService,
			app.NewMapsClient,
			app.NewDistanceService,
			app.NewTracerProvider,
			app.NewConfig,
			app.NewPrometheusExporter,
			app.NewMux,
		),
		// Since constructors are called lazily, we need some invocations to
		// kick-start our application. In this case, we'll use Register. Since
		// it depends on a *grpc.Server, calling it requires Fx to build those
		// types using the constructors above. Since we call NewGRPCServer, we
		// also register Lifecycle hooks to start and stop an gRPC server.
		fx.Invoke(app.Register),

		// This is optional. With this, you can control where Fx logs its
		// events. In this case, we're using a NopLogger to keep our test
		// silent. Normally, you'll want to use an fxevent.ZapLogger or an
		// fxevent.ConsoleLogger.
		fx.WithLogger(
			func() fxevent.Logger {
				//return fxevent.NopLogger
				return &fxevent.ConsoleLogger{W: os.Stdout}
			},
		),
	)

	a.Run()
}
