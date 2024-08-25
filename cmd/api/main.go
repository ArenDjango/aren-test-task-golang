package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/ArenDjango/golang-test-task/config"
	"github.com/ArenDjango/golang-test-task/logger"
	"github.com/ArenDjango/golang-test-task/pkg/metrics"
	"github.com/ArenDjango/golang-test-task/pkg/middlewares"
	"github.com/ArenDjango/golang-test-task/server"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

var (
	databaseURL = flag.String("database_url", "", "Database URL")
)

func startTracing() (*trace.TracerProvider, error) {
	headers := map[string]string{
		"content-type": "application/json",
	}

	exporter, err := otlptrace.New(
		context.Background(),
		otlptracehttp.NewClient(
			otlptracehttp.WithEndpoint("localhost:14268"),
			otlptracehttp.WithHeaders(headers),
			otlptracehttp.WithInsecure(),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("creating new exporter: %w", err)
	}

	tracerProvider := trace.NewTracerProvider(
		trace.WithBatcher(
			exporter,
			trace.WithMaxExportBatchSize(trace.DefaultMaxExportBatchSize),
			trace.WithBatchTimeout(trace.DefaultScheduleDelay*time.Millisecond),
		),
		trace.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String("test-task-golang-aren"),
			),
		),
	)

	otel.SetTracerProvider(tracerProvider)

	return tracerProvider, nil
}

func main() {
	//f, erro := os.Create("cpu.prof")
	//if erro != nil {
	//	log.Fatal(erro)
	//}
	//defer f.Close()
	//
	//if err := pprof.StartCPUProfile(f); err != nil {
	//	log.Fatal(err)
	//}
	//defer pprof.StopCPUProfile()
	//
	//for i := 0; i <= 5; i++ {
	//	g := i / 2
	//	fmt.Println(g)
	//}

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	cfg := config.Get()

	if *databaseURL != "" {
		cfg.PgURL = *databaseURL
	}

	l := logger.Get()

	registry := prometheus.NewRegistry()

	prometheusMetrics, err := metrics.New(cfg, registry)
	if err != nil {
		log.Fatalf("prometheus metrics start: %s", err) //nolint:gocritic // ??
	}

	s := server.NewServer(cfg, l, prometheusMetrics, registry)

	tracerProvider, err := startTracing()
	if err != nil {
		fmt.Printf("Failed to start tracer: %v\n", err)
	}
	defer func() {
		if err := tracerProvider.Shutdown(context.Background()); err != nil {
			fmt.Printf("Failed to shutdown tracer: %v\n", err)
		}
	}()

	// Creating a named tracer for the main function
	tracer := otel.Tracer("main-service")

	_, span := tracer.Start(context.Background(), "main")
	defer span.End()

	mux := http.NewServeMux()

	mux.Handle("/metrics", promhttp.Handler())
	mux.Handle("/", middlewares.PrometheusMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second) // Симуляция длительного запроса
		w.Write([]byte("Hello, Prometheus!"))
	})))
	http.Handle("/", mux)

	if err = s.Run(); err != nil {
		log.Panicf("Cannot start server: %v", err)
	}

	return nil
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}
