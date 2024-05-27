package metrics

import (
	"github.com/ArenDjango/golang-test-task/config"
	"github.com/ArenDjango/golang-test-task/pkg/csd/logger/log"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

const (
	MetricEngineFinalizeDuration = "engine_finalize_duration"
	MetricInexactEngineTurnover  = "inexact_engine_turnover"
	MetricHTTPReqsTotal          = "http_requests_total"
	MetricHTTPReqsInFlight       = "http_request_in_flight"
	MetricHTTPReqDuration        = "http_request_duration"
	MetricEngineAPICallTime      = "engine_api_call_time"
	MetricEngineRejects          = "engine_rejects"
	MetricRejectedTurnover       = "rejected_turnover"
	MetricEngineErrors           = "engine_errors"
	MetricClientOrders           = "client_orders"
)

type Metrics struct {
	Registerer prometheus.Registerer
	cfg        *config.Config

	engineAPICallTime     *prometheus.SummaryVec
	engineFinalizeSeconds *prometheus.SummaryVec
	engineTurnover        *prometheus.CounterVec

	engineErrors     *prometheus.GaugeVec
	engineRejects    *prometheus.CounterVec
	rejectedTurnover *prometheus.CounterVec
	clientOrders     *prometheus.CounterVec

	requestsTotal   *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
	requestInFlight *prometheus.GaugeVec
}

func (m *Metrics) SetEngineFinalizeDuration(orderType, statusID, client, engine string, seconds float64) {
	m.engineFinalizeSeconds.WithLabelValues(orderType, statusID, client, engine).Observe(seconds)
}

func (m *Metrics) AddEngineSuccessPaid(orderType, client, engine, currencyID string, amount float64) {
	if amount < 0 {
		return
	}

	m.engineTurnover.WithLabelValues(orderType, client, engine, currencyID).Add(amount)
}

func (m *Metrics) SetEngineAPICallDuration(orderType, engine string, seconds float64) {
	m.engineAPICallTime.WithLabelValues(orderType, engine).Observe(seconds)
}

func (m *Metrics) IncEngineErrors(orderType, engine string) {
	m.engineErrors.WithLabelValues(orderType, engine).Inc()
}

func (m *Metrics) IncEngineRejects(orderType, engine string) {
	m.engineRejects.WithLabelValues(orderType, engine).Inc()
}

func (m *Metrics) AddRejectedTurnover(orderType, engineID, currencyID, bankName string, amount float64) {
	m.rejectedTurnover.WithLabelValues(orderType, engineID, currencyID, bankName).Add(amount)
}

func (m *Metrics) IncClientOrders(orderType, client, engine string) {
	m.clientOrders.WithLabelValues(orderType, client, engine).Inc()
}

func (m *Metrics) IncRequestInFlight(path, clientName string) {
	m.requestInFlight.WithLabelValues(path, clientName).Inc()
}

func (m *Metrics) DecRequestInFlight(path, clientName string) {
	m.requestInFlight.WithLabelValues(path, clientName).Dec()
}

func (m *Metrics) IncRequestTotal(statusCode, path, clientName string) {
	m.requestsTotal.WithLabelValues(statusCode, path, clientName).Inc()
}

func (m *Metrics) ObserveRequestDuration(statusCode, path, clientName string, elapsed float64) {
	m.requestDuration.WithLabelValues(statusCode, path, clientName).Observe(elapsed)
}

func New(cfg *config.Config, registry *prometheus.Registry) (*Metrics, error) {
	m := &Metrics{
		cfg:        cfg,
		Registerer: registry,
	}

	m.requestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: prometheus.BuildFQName(cfg.Metrics.Namespace, "", MetricHTTPReqsTotal),
		},
		[]string{"status_code", "path", "client_name"},
	)

	err := registry.Register(m.requestsTotal)
	if err != nil {
		return nil, errors.Wrapf(err, MetricHTTPReqsTotal)
	}

	// duration
	m.requestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: prometheus.BuildFQName(cfg.Metrics.Namespace, "", MetricHTTPReqDuration),
		Buckets: []float64{
			0.00001, // 10µs
			0.0001,  // 100µs
			0.001,   // 1ms
			0.01,    // 10ms
			0.1,     // 100 ms
			1.0,     // 1s
			10.0,    // 10s
			15.0,
			20.0,
			30.0,
		},
	},
		[]string{"status_code", "path", "client_name"},
	)

	err = registry.Register(m.requestDuration)
	if err != nil {
		return nil, errors.Wrapf(err, MetricHTTPReqDuration)
	}

	// http_request_in_flight
	m.requestInFlight = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: prometheus.BuildFQName(cfg.Metrics.Namespace, "", MetricHTTPReqsInFlight),
	},
		[]string{"path", "client_name"},
	)

	err = registry.Register(m.requestInFlight)
	if err != nil {
		return nil, errors.Wrapf(err, MetricHTTPReqsInFlight)
	}

	// engineAPICallTime

	m.engineAPICallTime = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name: prometheus.BuildFQName(cfg.Metrics.Namespace, "", MetricEngineAPICallTime),
	},
		[]string{"order_type", "engine"},
	)

	err = registry.Register(m.engineAPICallTime)
	if err != nil {
		return nil, errors.Wrapf(err, MetricEngineAPICallTime)
	}

	m.engineFinalizeSeconds = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name: prometheus.BuildFQName(cfg.Metrics.Namespace, "", MetricEngineFinalizeDuration),
	},
		[]string{"order_type", "status_id", "client", "engine"},
	)

	err = registry.Register(m.engineFinalizeSeconds)
	if err != nil {
		return nil, errors.Wrapf(err, MetricEngineFinalizeDuration)
	}

	m.engineTurnover = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: prometheus.BuildFQName(cfg.Metrics.Namespace, "", MetricInexactEngineTurnover),
	},
		[]string{"order_type", "client", "engine", "currency"},
	)

	err = registry.Register(m.engineTurnover)
	if err != nil {
		return nil, errors.Wrapf(err, MetricInexactEngineTurnover)
	}

	// engineErrors
	m.engineErrors = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: prometheus.BuildFQName(cfg.Metrics.Namespace, "", MetricEngineErrors),
	},
		[]string{"order_type", "engine"},
	)

	err = registry.Register(m.engineErrors)
	if err != nil {
		return nil, errors.Wrapf(err, MetricEngineErrors)
	}

	// engineRejects
	m.engineRejects = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: prometheus.BuildFQName(cfg.Metrics.Namespace, "", MetricEngineRejects),
	},
		[]string{"order_type", "engine"},
	)

	err = registry.Register(m.engineRejects)
	if err != nil {
		return nil, errors.Wrapf(err, MetricEngineRejects)
	}

	m.rejectedTurnover = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: prometheus.BuildFQName(cfg.Metrics.Namespace, "", MetricRejectedTurnover),
	},
		[]string{"order_type", "engine_id", "currency_id", "bank_name"},
	)

	err = registry.Register(m.rejectedTurnover)
	if err != nil {
		return nil, errors.Wrapf(err, MetricRejectedTurnover)
	}

	m.clientOrders = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: prometheus.BuildFQName(cfg.Metrics.Namespace, "", MetricClientOrders),
	},
		[]string{"order_type", "client", "engine"},
	)

	err = registry.Register(m.clientOrders)
	if err != nil {
		return nil, errors.Wrapf(err, MetricClientOrders)
	}

	go func() {
		mux := http.NewServeMux()
		mux.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{
			Registry: registry,
		}))

		err := http.ListenAndServe(cfg.Metrics.Host, mux)
		if err != nil {
			log.Fatalf("Metrics http.ListenAndServer: %v\n", err)
		}

		log.Infof("Metrics server is running on port: %s", cfg.Metrics.Host)
	}()

	return m, nil
}
