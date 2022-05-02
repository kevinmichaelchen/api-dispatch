package service

import (
	"context"
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type Service struct {
	logger *log.Logger
}

func NewService(logger *log.Logger) *Service {
	return &Service{logger: logger}
}

func (s *Service) Ingest(ctx context.Context, r *v1beta1.IngestRequest) (*v1beta1.IngestResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Unimplemented")
}

func (s *Service) Dispatch(ctx context.Context, r *v1beta1.DispatchRequest) (*v1beta1.DispatchResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Unimplemented")
}
