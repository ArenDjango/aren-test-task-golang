package grpcsrv

import (
	"github.com/ArenDjango/golang-test-task/pkg/csd/logger"
	"github.com/ArenDjango/golang-test-task/pkg/csd/logger/log"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	grpcprom "github.com/ArenDjango/golang-test-task/pkg/go-grpc-middleware/providers/prometheus"
	"github.com/prometheus/client_golang/prometheus"
)

var histogramRequestDurationBuckets = []float64{0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10, 50}

type Options struct {
	MaxConnectionIdle time.Duration
	Timeout           time.Duration
	MaxConnectionAge  time.Duration
	Time              time.Duration
}

type MetricsOptions struct {
	Registerer prometheus.Registerer
	Subsystem  string
}

func New(opts Options, metricsOpts *MetricsOptions) (*grpc.Server, error) {
	var chainInterceptor grpc.ServerOption

	if metricsOpts == nil {
		chainInterceptor = grpc.ChainUnaryInterceptor(
			logger.NewUnaryServerInterceptorWithLogger(log.Logger),
		)
	} else {
		srvMetrics := grpcprom.NewServerMetrics(
			grpcprom.WithServerCounterOptions(grpcprom.WithSubsystem(metricsOpts.Subsystem)),
			grpcprom.WithServerHandlingTimeHistogram(
				grpcprom.WithHistogramBuckets(histogramRequestDurationBuckets),
				grpcprom.WithHistogramSubsystem(metricsOpts.Subsystem),
			),
		)

		if err := metricsOpts.Registerer.Register(srvMetrics); err != nil {
			return nil, err
		}

		chainInterceptor = grpc.ChainUnaryInterceptor(
			srvMetrics.UnaryServerInterceptor(),
			logger.NewUnaryServerInterceptorWithLogger(log.Logger),
		)
	}

	srv := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: opts.MaxConnectionIdle,
			Timeout:           opts.Timeout,
			MaxConnectionAge:  opts.MaxConnectionAge,
			Time:              opts.Time,
		}),
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
		chainInterceptor,
	)
	return srv, nil
}
