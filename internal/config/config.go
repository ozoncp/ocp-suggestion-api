package config

import (
	"sync"

	"github.com/caarlos0/env/v6"
	"github.com/rs/zerolog/log"
)

type config struct {
	GrpcPort    string   `env:"GRPC_PORT" envDefault:"8082"`
	HttpPort    string   `env:"HTTP_PORT" envDefault:"8081"`
	MetricsPort string   `env:"METRICS_PORT" envDefault:"9100"`
	DBName      string   `env:"POSTGRES_DB,unset,notEmpty"`
	DBUser      string   `env:"POSTGRES_USER,unset,notEmpty"`
	DBPassword  string   `env:"POSTGRES_PASSWORD,unset,notEmpty"`
	DBHost      string   `env:"POSTGRES_HOST,unset,notEmpty"`
	DBPort      uint     `env:"POSTGRES_PORT,unset,notEmpty"`
	BatchSize   uint     `env:"BATCH_SIZE" envDefault:"100"`
	Brokers     []string `env:"KAFKA_BROKERS,unset,notEmpty" envSeparator:","`
	TracerAddr  string   `env:"TRACER_ADDR" envDefault:"localhost:6831"`
}

// Global config
var Config config
var doOnce sync.Once

func init() {
	doOnce.Do(func() {
		if err := ReadConfig(); err != nil {
			log.Fatal().Msgf("InitConfig error: %+v", err)
		}
	})
}

// ReadConfig читает конфиг из окружения и обновляет его
func ReadConfig() error {
	if err := env.Parse(&Config); err != nil {
		log.Error().Msgf("failed to read required environment variables: %+v", err)
		return err
	}

	return nil
}
