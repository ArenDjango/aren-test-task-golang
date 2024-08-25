package middlewares

import (
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"time"
)

// Метрики
var (
	requestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status_code"},
	)

	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path", "status_code"},
	)

	inProgressRequests = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "http_in_progress_requests",
			Help: "Current number of in-progress HTTP requests",
		},
		[]string{"method", "path"},
	)
)

func init() {
	// Регистрация метрик в Prometheus
	prometheus.MustRegister(requestsTotal)
	prometheus.MustRegister(requestDuration)
	prometheus.MustRegister(inProgressRequests)
}

// Middleware для мониторинга
func PrometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Засекаем время начала обработки запроса
		start := time.Now()

		// Увеличиваем количество активных запросов
		inProgressRequests.WithLabelValues(r.Method, r.URL.Path).Inc()
		defer inProgressRequests.WithLabelValues(r.Method, r.URL.Path).Dec()

		// Перехватываем статус код ответа
		rr := &responseRecorder{w, http.StatusOK}
		next.ServeHTTP(rr, r)

		// Фиксируем метрики
		duration := time.Since(start).Seconds()
		requestsTotal.WithLabelValues(r.Method, r.URL.Path, http.StatusText(rr.statusCode)).Inc()
		requestDuration.WithLabelValues(r.Method, r.URL.Path, http.StatusText(rr.statusCode)).Observe(duration)
	})
}

// Обертка для перехвата статус кода
type responseRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (rr *responseRecorder) WriteHeader(statusCode int) {
	rr.statusCode = statusCode
	rr.ResponseWriter.WriteHeader(statusCode)
}
