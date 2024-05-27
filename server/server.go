package server

import (
	"github.com/ArenDjango/golang-test-task/config"
	logger "github.com/ArenDjango/golang-test-task/logger"
	"github.com/ArenDjango/golang-test-task/pkg/csd/logger/log"
	"github.com/ArenDjango/golang-test-task/pkg/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	logger            *logger.Logger
	GRPCServerAPI     *grpc.Server
	cfg               *config.Config
	registry          *prometheus.Registry
	prometheusMetrics *metrics.Metrics
}

func NewServer(
	cfg *config.Config,
	logger *logger.Logger,
	prometheusMetric *metrics.Metrics,
	registry *prometheus.Registry,
) *Server {
	return &Server{
		cfg:               cfg,
		logger:            logger,
		registry:          registry,
		prometheusMetrics: prometheusMetric,
	}
}

var instanceStartTime string

func (s *Server) Run() error {
	instanceStartTime = time.Now().Format(time.RFC1123)

	apiListener, err := net.Listen("tcp", s.cfg.GRPCServerAPI.Host)
	if err != nil {
		return err
	}

	if err = s.CreateGPRCServers(); err != nil {
		log.Fatalf("failed to initialize grpc servers: %v", err)
	}

	if err := s.MapHandlers(s.cfg); err != nil {
		log.Fatalf("Cannot map handlers: %v", err) //nolint:gocritic // ??
	}

	go func() {
		log.Infof("GRPCServerAPI server is listening on %v", s.cfg.GRPCServerAPI.Host)

		if err := s.GRPCServerAPI.Serve(apiListener); err != nil {
			log.Fatalf("Error starting GRPC Server: %v", err)
		}
	}()

	log.Infof("GracefulStop. App is running. Wait for stop instance %s", instanceStartTime)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)
	<-quit
	log.Infof("GracefulStop. Read signal for stop instance %s", instanceStartTime)
	s.GRPCServerAPI.GracefulStop()
	log.Infof("GracefulStop. GRPC stopped for instance %s", instanceStartTime)

	return nil
}
