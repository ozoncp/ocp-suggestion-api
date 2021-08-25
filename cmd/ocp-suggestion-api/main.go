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

	"github.com/caarlos0/env/v6"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	"github.com/ozoncp/ocp-suggestion-api/internal/api"
	"github.com/ozoncp/ocp-suggestion-api/internal/repo"
	desc "github.com/ozoncp/ocp-suggestion-api/pkg/ocp-suggestion-api"
)

type config struct {
	GrpcPort   string `env:"GRPC_PORT" envDefault:"8082"`
	HttpPort   string `env:"HTTP_PORT" envDefault:"8081"`
	DBName     string `env:"POSTGRES_DB,unset,notEmpty"`
	DBUser     string `env:"POSTGRES_USER,unset,notEmpty"`
	DBPassword string `env:"POSTGRES_PASSWORD,unset,notEmpty"`
	DBHost     string `env:"POSTGRES_HOST,unset,notEmpty"`
	DBPort     uint   `env:"POSTGRES_PORT,unset,notEmpty"`
}

func connectDB(cfg *config) (*sqlx.DB, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s connect_timeout=10 sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
		cfg.DBUser,
		cfg.DBPassword,
	)
	return sqlx.Connect("pgx", dataSourceName)
}

func runGRPC(ctx context.Context, cfg *config, db *sqlx.DB) error {
	listen, err := net.Listen("tcp", ":"+cfg.GrpcPort)
	if err != nil {
		return fmt.Errorf("failed to Listen: %w", err)
	}

	s := grpc.NewServer()
	repoDB := repo.NewRepo(db)
	desc.RegisterOcpSuggestionApiServer(s, api.NewSuggestionAPI(repoDB))

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

func runHTTP(ctx context.Context, cfg *config) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := desc.RegisterOcpSuggestionApiHandlerFromEndpoint(ctx, mux, ":"+cfg.GrpcPort, opts)
	if err != nil {
		return fmt.Errorf("failed to RegisterOcpSuggestionApiHandler: %w", err)
	}

	srv := &http.Server{
		Addr:    ":" + cfg.HttpPort,
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

	cfg := &config{}
	if err := env.Parse(cfg); err != nil {
		log.Fatal().Msgf("failed to read required environment variables: %+v", err)
	}

	db, err := connectDB(cfg)
	if err != nil {
		log.Fatal().Msgf("failed to connect to database: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Error().Err(err).Msgf("failed to close database")
		}
	}()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := runGRPC(ctx, cfg, db); err != nil {
			log.Fatal().Msgf("runGRPC: %v", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := runHTTP(ctx, cfg); err != nil {
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
