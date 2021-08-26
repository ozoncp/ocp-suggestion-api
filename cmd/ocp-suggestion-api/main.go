package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	"github.com/ozoncp/ocp-suggestion-api/internal/api"
	cfg "github.com/ozoncp/ocp-suggestion-api/internal/config"
	"github.com/ozoncp/ocp-suggestion-api/internal/metrics"
	"github.com/ozoncp/ocp-suggestion-api/internal/producer"
	"github.com/ozoncp/ocp-suggestion-api/internal/repo"
	"github.com/ozoncp/ocp-suggestion-api/internal/tracer"
	desc "github.com/ozoncp/ocp-suggestion-api/pkg/ocp-suggestion-api"
)

func runMetrics(ctx context.Context) error {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	srv := &http.Server{
		Addr:    ":" + cfg.Config.MetricsPort,
		Handler: mux,
	}

	metrics.RegisterMetrics()

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Msgf("metrics: listen: %+s", err)
		}
	}()

	log.Info().Msg("metrics started")
	defer log.Info().Msg("metrics stopped")
	<-ctx.Done()

	ctxShutDown, cancelShutDown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutDown()
	if err := srv.Shutdown(ctxShutDown); err != nil {
		return fmt.Errorf("metrics Shutdown failed: %w", err)
	}

	return nil
}

func connectDB() (*sqlx.DB, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s connect_timeout=10 sslmode=disable",
		cfg.Config.DBHost,
		cfg.Config.DBPort,
		cfg.Config.DBName,
		cfg.Config.DBUser,
		cfg.Config.DBPassword,
	)
	return sqlx.Connect("pgx", dataSourceName)
}

func runGRPC(ctx context.Context, db *sqlx.DB, prod producer.Producer) error {
	listen, err := net.Listen("tcp", ":"+cfg.Config.GrpcPort)
	if err != nil {
		return fmt.Errorf("failed to Listen: %w", err)
	}

	s := grpc.NewServer()
	repoDB := repo.NewRepo(db)

	desc.RegisterOcpSuggestionApiServer(s, api.NewSuggestionAPI(repoDB, cfg.Config.BatchSize, prod))

	go func() {
		if err = s.Serve(listen); err != nil {
			log.Fatal().Msgf("grpc: failed to Serve: %+s", err)
		}
	}()

	log.Info().Msg("grpc started")
	defer log.Info().Msg("grpc stopped")
	<-ctx.Done()

	s.GracefulStop()

	return nil
}

func runHTTP(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := desc.RegisterOcpSuggestionApiHandlerFromEndpoint(ctx, mux, ":"+cfg.Config.GrpcPort, opts)
	if err != nil {
		return fmt.Errorf("failed to RegisterOcpSuggestionApiHandler: %w", err)
	}

	srv := &http.Server{
		Addr:    ":" + cfg.Config.HttpPort,
		Handler: mux,
	}
	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Msgf("grpc-gateway: listen: %+s", err)
		}
	}()

	log.Info().Msg("grpc-gateway started")
	defer log.Info().Msg("grpc-gateway stopped")
	<-ctx.Done()

	ctxShutDown, cancelShutDown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutDown()
	if err = srv.Shutdown(ctxShutDown); err != nil {
		return fmt.Errorf("grpc-gateway Shutdown failed: %w", err)
	}

	return nil
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := connectDB()
	if err != nil {
		log.Fatal().Msgf("failed to connect to database: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Error().Err(err).Msgf("failed to close database")
		}
	}()

	prod, err := producer.NewProducer()
	if err != nil {
		log.Fatal().Msgf("failed to connect to broker: %v", err)
	}
	defer func() {
		if err := prod.Close(); err != nil {
			log.Error().Err(err).Msgf("failed to close connection to broker")
		}
	}()

	traceCloser, err := tracer.InitTracing()
	if err != nil {
		log.Fatal().Msgf("failed to initTracing: %v", err)
	}
	defer func() {
		if err := traceCloser.Close(); err != nil {
			log.Error().Err(err).Msgf("failed to close tracer")
		}
	}()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := runMetrics(ctx); err != nil {
			log.Fatal().Msgf("runMetrics: %v", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := runGRPC(ctx, db, prod); err != nil {
			log.Fatal().Msgf("runGRPC: %v", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := runHTTP(ctx); err != nil {
			log.Fatal().Msgf("runHTTP: %v", err)
		}
	}()

	log.Info().Msg("ocp-suggestion-api started")
	defer log.Info().Msg("ocp-suggestion-api stopped")

	// Graceful shutdown
	interruptCh := make(chan os.Signal, 1)
	signal.Notify(interruptCh, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(interruptCh)
	oscall := <-interruptCh
	log.Printf("got signal %+v, attempting graceful shutdown...", oscall)
	cancel()

	wg.Wait()
}
