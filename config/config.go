package config

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"sync"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	GARANTEX_API_URL string `envconfig:"GARANTEX_API_URL"`
	LogLevel         string `envconfig:"LOG_LEVEL"`
	PgURL            string `envconfig:"PG_URL"`
	PgMigrationsPath string `envconfig:"PG_MIGRATIONS_PATH"`
	HTTPAddr         string `envconfig:"HTTP_ADDR"`
	GCBucket         string `envconfig:"GC_BUCKET"`
	FilePath         string `envconfig:"FILE_PATH"`
	GRPCServerAPI    struct {
		Host              string        `envconfig:"GRPC_SERVER_API_HOST"`
		Port              int           `envconfig:"GRPC_SERVER_API_PORT"`
		MaxConnectionIdle time.Duration `envconfig:"GRPC_SERVER_API_MAX_CONNECTION_IDLE"`
		Timeout           time.Duration `envconfig:"GRPC_SERVER_API_TIMEOUT"`
		MaxConnectionAge  time.Duration `envconfig:"GRPC_SERVER_API_MAX_CONNECTION_AGE"`
		Time              time.Duration `envconfig:"GRPC_SERVER_API_TIME"`
	}
	OpenTelemetry struct {
		URL         string `envconfig:"OPENTELEMETRY_URL"`
		Password    string `envconfig:"OPENTELEMETRY_PASSWORD"`
		Username    string `envconfig:"OPENTELEMETRY_USERNAME"`
		ServiceName string `envconfig:"OPENTELEMETRY_SERVICE_NAME"`
	}
	Sentry struct {
		DNS         string `envconfig:"SENTRY_DSN"`
		IgnoreFrame string `envconfig:"SENTRY_IGNORE_FRAME"`
	}
	Metrics struct {
		Host      string `envconfig:"METRICS_HOST"`
		Namespace string `envconfig:"METRICS_NAMESPACE"`
	}
}

var (
	config Config
	once   sync.Once
)

func Get(envFilePath ...string) *Config {
	once.Do(func() {
		envPath := ".env"
		if len(envFilePath) > 0 {
			envPath = envFilePath[0]
		}
		err := godotenv.Load(envPath)
		if err != nil {
			log.Printf("No .env file found or error loading .env file: %v", err)
		}

		err = envconfig.Process("", &config)
		if err != nil {
			log.Fatal(err)
		}
		configBytes, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Configuration:", string(configBytes))
	})
	return &config
}
