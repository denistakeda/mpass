package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/denistakeda/mpass/internal/auth_service"
	"github.com/denistakeda/mpass/internal/config"
	"github.com/denistakeda/mpass/internal/logging"
	"github.com/denistakeda/mpass/internal/ports"
	"github.com/denistakeda/mpass/internal/server"
	"github.com/denistakeda/mpass/internal/user_store"
)

type (
	runnable interface {
		Start() <-chan error
		Stop()
	}
)

func main() {
	logService := logging.New()
	logger := logService.ComponentLogger("main")

	conf, err := config.ParseServerCfg()
	if err != nil {
		logger.Fatal().Msg("failed to read the configuration")
	}
	logger.Info().Msgf("Configuration: %v", conf)

	interruptChan := handleInterrupt()

	srv := buildServer(conf, logService)
	serverErrors := srv.Start()
	defer srv.Stop()

	select {
	case serverError := <-serverErrors:
		logger.Error().Err(serverError).Msg("server error")
	case <-interruptChan:
		logger.Info().Msg("Server was interrupted")
	}
}

func buildServer(conf config.Config, logService ports.LogService) runnable {
	userStore := user_store.NewInMemory()

	authService := auth_service.New(auth_service.NewAuthServiceParams{
		Secret: conf.Secret,

		LogService: logService,
		UserStore:  userStore,
	})

	s := server.New(server.NewServerParams{
		Host:        conf.Host,
		LogService:  logService,
		AuthService: authService,
	})

	return s
}

func handleInterrupt() <-chan os.Signal {
	out := make(chan os.Signal, 2)
	signal.Notify(out, os.Interrupt)
	signal.Notify(out, syscall.SIGTERM)
	return out
}
