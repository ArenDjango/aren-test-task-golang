package server

import (
	"context"
	"github.com/ArenDjango/golang-test-task/config"
	apiproto "github.com/ArenDjango/golang-test-task/protos/rates"
	"github.com/ArenDjango/golang-test-task/service"
	"github.com/ArenDjango/golang-test-task/store"
	"github.com/ArenDjango/golang-test-task/transport/delivery/apigrpc"
	"github.com/pkg/errors"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func (s *Server) MapHandlers(cfg *config.Config) error {
	ctx := context.Background()
	store, err := store.New(ctx)
	if err != nil {
		return errors.Wrap(err, "store.New failed")
	}
	serviceManager, err := service.NewManager(ctx, store, cfg)
	if err != nil {
		return errors.Wrap(err, "manager.New failed")
	}

	apiproto.RegisterRatesServiceServer(s.GRPCServerAPI, apigrpc.NewAPIService(s.cfg, serviceManager.Rates))
	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(s.GRPCServerAPI, healthServer)

	// Установка начального статуса проверки состояния
	healthServer.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)

	return nil
}
