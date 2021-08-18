package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	"github.com/ozoncp/ocp-suggestion-api/internal/api"
	desc "github.com/ozoncp/ocp-suggestion-api/pkg/ocp-suggestion-api"
)

const (
	grpcPort     = ":8082"
	grpcEndpoint = "localhost:8082"
	httpPort     = ":8080"
)

func runGRPC() error {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		return fmt.Errorf("failed to Listen: %w", err)
	}

	s := grpc.NewServer()
	desc.RegisterOcpSuggestionApiServer(s, api.NewSuggestionAPI())

	if err = s.Serve(listen); err != nil {
		return fmt.Errorf("failed to Serve: %w", err)
	}

	return nil
}

func runHTTP() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := desc.RegisterOcpSuggestionApiHandlerFromEndpoint(ctx, mux, grpcEndpoint, opts)
	if err != nil {
		log.Fatal().Msgf("runHTTP: failed to RegisterOcpSuggestionApiHandler: %v", err)
	}

	err = http.ListenAndServe(httpPort, mux)
	if err != nil {
		log.Fatal().Msgf("runHTTP: failed to ListenAndServe: %v", err)
	}
}

func main() {
	log.Printf("ocp-suggestion-api started")

	go runHTTP()

	if err := runGRPC(); err != nil {
		log.Fatal().Msgf("runGRPC: %v", err)
	}
}
