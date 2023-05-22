package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/denistakeda/mpass/internal/config"
	"github.com/denistakeda/mpass/internal/logging"
	"github.com/denistakeda/mpass/internal/server"
)

// Generate protobuf specification:
// go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/mpass.proto
func main() {
	logService := logging.New()
	logger := logService.ComponentLogger("main")

	conf, err := config.ParseServerCfg()
	if err != nil {
		logger.Fatal().Msg("failed to read the configuration")
	}
	logger.Info().Msgf("Configuration: %v", conf)

	interruptChan := handleInterrupt()

	srv := server.New(server.NewServerParams{
		Host:       conf.Host,
		LogService: logService,
		// TODO: create auth service
		// AuthService: authService,
	})
	serverErrors := srv.Start()
	defer srv.Stop()

	select {
	case serverError := <-serverErrors:
		logger.Error().Err(serverError).Msg("server error")
	case <-interruptChan:
		logger.Info().Msg("Server was interrupted")
	}
}

func handleInterrupt() <-chan os.Signal {
	out := make(chan os.Signal, 2)
	signal.Notify(out, os.Interrupt)
	signal.Notify(out, syscall.SIGTERM)
	return out
}
