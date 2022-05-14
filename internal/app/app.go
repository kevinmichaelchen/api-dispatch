package app

import (
	"github.com/kevinmichaelchen/api-dispatch/internal/app/config"
	"github.com/kevinmichaelchen/api-dispatch/internal/app/grpc"
	"github.com/kevinmichaelchen/api-dispatch/internal/app/logging"
	"github.com/kevinmichaelchen/api-dispatch/internal/app/metrics"
	"github.com/kevinmichaelchen/api-dispatch/internal/app/service"
	"github.com/kevinmichaelchen/api-dispatch/internal/app/sql"
	"github.com/kevinmichaelchen/api-dispatch/internal/app/tracing"
	"go.uber.org/fx"
)

var Module = fx.Options(
	config.Module,
	grpc.Module,
	logging.Module,
	metrics.Module,
	service.Module,
	sql.Module,
	tracing.Module,
)
