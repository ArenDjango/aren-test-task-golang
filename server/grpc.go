package server

import (
	"fmt"
	"github.com/ArenDjango/golang-test-task/pkg/utils/grpcsrv"
)

func (s *Server) CreateGPRCServers() error {
	a, err := grpcsrv.New(
		grpcsrv.Options{
			MaxConnectionIdle: s.cfg.GRPCServerAPI.MaxConnectionIdle,
			Timeout:           s.cfg.GRPCServerAPI.Timeout,
			MaxConnectionAge:  s.cfg.GRPCServerAPI.MaxConnectionAge,
			Time:              s.cfg.GRPCServerAPI.Time,
		},
		&grpcsrv.MetricsOptions{
			Registerer: s.prometheusMetrics.Registerer,
			//Subsystem:  fmt.Sprintf("%s_payment_api_server", s.cfg.Metrics.Namespace),
		},
	)
	if err != nil {
		return fmt.Errorf("failed to construct grpc a server: %w", err)
	}

	s.GRPCServerAPI = a

	return nil
}
