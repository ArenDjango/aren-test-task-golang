// Copyright (c) The go-grpc-middleware Authors.
// Licensed under the Apache License 2.0.

package prometheus

import (
	"context"
	"github.com/ArenDjango/golang-test-task/pkg/go-grpc-middleware/v2/interceptors"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
)

type ClientMetrics struct {
	clientStartedCounter    *prometheus.CounterVec
	clientHandledCounter    *prometheus.CounterVec
	clientStreamMsgReceived *prometheus.CounterVec
	clientStreamMsgSent     *prometheus.CounterVec

	clientHandledHistogram    *prometheus.HistogramVec
	clientStreamRecvHistogram *prometheus.HistogramVec
	clientStreamSendHistogram *prometheus.HistogramVec
}

func NewClientMetrics(opts ...ClientMetricsOption) *ClientMetrics {
	var config clientMetricsConfig
	config.apply(opts)
	return &ClientMetrics{
		clientStartedCounter: prometheus.NewCounterVec(
			config.counterOpts.apply(prometheus.CounterOpts{
				Name: "grpc_client_started_total",
				Help: "Total number of RPCs started on the client.",
			}), []string{"grpc_type", "grpc_service", "grpc_method"}),

		clientHandledCounter: prometheus.NewCounterVec(
			config.counterOpts.apply(prometheus.CounterOpts{
				Name: "grpc_client_handled_total",
				Help: "Total number of RPCs completed by the client, regardless of success or failure.",
			}), []string{"grpc_type", "grpc_service", "grpc_method", "grpc_code"}),

		clientStreamMsgReceived: prometheus.NewCounterVec(
			config.counterOpts.apply(prometheus.CounterOpts{
				Name: "grpc_client_msg_received_total",
				Help: "Total number of RPC stream messages received by the client.",
			}), []string{"grpc_type", "grpc_service", "grpc_method"}),

		clientStreamMsgSent: prometheus.NewCounterVec(
			config.counterOpts.apply(prometheus.CounterOpts{
				Name: "grpc_client_msg_sent_total",
				Help: "Total number of gRPC stream messages sent by the client.",
			}), []string{"grpc_type", "grpc_service", "grpc_method"}),

		clientHandledHistogram:    config.clientHandledHistogram,
		clientStreamRecvHistogram: config.clientStreamRecvHistogram,
		clientStreamSendHistogram: config.clientStreamSendHistogram,
	}
}

func (m *ClientMetrics) Describe(ch chan<- *prometheus.Desc) {
	m.clientStartedCounter.Describe(ch)
	m.clientHandledCounter.Describe(ch)
	m.clientStreamMsgReceived.Describe(ch)
	m.clientStreamMsgSent.Describe(ch)
	if m.clientHandledHistogram != nil {
		m.clientHandledHistogram.Describe(ch)
	}
	if m.clientStreamRecvHistogram != nil {
		m.clientStreamRecvHistogram.Describe(ch)
	}
	if m.clientStreamSendHistogram != nil {
		m.clientStreamSendHistogram.Describe(ch)
	}
}

func (m *ClientMetrics) Collect(ch chan<- prometheus.Metric) {
	m.clientStartedCounter.Collect(ch)
	m.clientHandledCounter.Collect(ch)
	m.clientStreamMsgReceived.Collect(ch)
	m.clientStreamMsgSent.Collect(ch)
	if m.clientHandledHistogram != nil {
		m.clientHandledHistogram.Collect(ch)
	}
	if m.clientStreamRecvHistogram != nil {
		m.clientStreamRecvHistogram.Collect(ch)
	}
	if m.clientStreamSendHistogram != nil {
		m.clientStreamSendHistogram.Collect(ch)
	}
}

func (r *reportable) ClientReporter(ctx context.Context, meta interceptors.CallMeta) (interceptors.Reporter, context.Context) {
	// Реализуйте логику создания Reporter здесь
	return nil, ctx
}

func (m *ClientMetrics) UnaryClientInterceptor(opts ...Option) grpc.UnaryClientInterceptor {
	return interceptors.UnaryClientInterceptor(&reportable{
		opts:          opts,
		clientMetrics: m,
	})
}

func (m *ClientMetrics) StreamClientInterceptor(opts ...Option) grpc.StreamClientInterceptor {
	return interceptors.StreamClientInterceptor(&reportable{
		opts:          opts,
		clientMetrics: m,
	})
}
