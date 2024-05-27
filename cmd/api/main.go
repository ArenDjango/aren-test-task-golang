package main

import (
	"flag"
	"github.com/ArenDjango/golang-test-task/config"
	"github.com/ArenDjango/golang-test-task/logger"
	"github.com/ArenDjango/golang-test-task/pkg/metrics"
	"github.com/ArenDjango/golang-test-task/server"
	"github.com/prometheus/client_golang/prometheus"

	"log"
)

var (
	databaseURL = flag.String("database_url", "", "Database URL")
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
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

	if err = s.Run(); err != nil {
		log.Panicf("Cannot start server: %v", err)
	}

	return nil
}
