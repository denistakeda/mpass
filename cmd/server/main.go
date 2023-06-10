package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/denistakeda/mpass/internal/auth_service"
	"github.com/denistakeda/mpass/internal/config"
	"github.com/denistakeda/mpass/internal/db"
	"github.com/denistakeda/mpass/internal/logging"
	"github.com/denistakeda/mpass/internal/ports"
	"github.com/denistakeda/mpass/internal/record_service"
	"github.com/denistakeda/mpass/internal/record_store"
	"github.com/denistakeda/mpass/internal/server"
	"github.com/denistakeda/mpass/internal/user_store"
	"github.com/rs/zerolog"
)

type (
	srv interface {
		Start() <-chan error
		Stop()
		Host() string
	}
)

func main() {
	logService := logging.New()
	logger := logService.ComponentLogger("main")

	conf, err := config.ParseServerCfg()
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to read the configuration")
	}
	logger.Info().Msgf("Configuration: %v", conf)

	interruptChan := handleInterrupt()

	srv := buildServer(buildParams{
		conf:       conf,
		logService: logService,
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

type buildParams struct {
	conf                config.Config
	logService          ports.LogService
	useInMemoryStorages bool
}

func buildServer(params buildParams) srv {
	logger := params.logService.ComponentLogger("buildServer")

	// Stores
	userStore, recordStore := makeStores(logger, params.conf.DatabaseURI, params.useInMemoryStorages)

	// Services
	authService := auth_service.New(auth_service.NewAuthServiceParams{
		Secret: params.conf.Secret,

		LogService: params.logService,
		UserStore:  userStore,
	})

	recordService := record_service.New(params.logService, recordStore)

	// One ring to rule them all
	s := server.New(server.NewServerParams{
		Host:          params.conf.Host,
		LogService:    params.logService,
		AuthService:   authService,
		RecordService: recordService,
	})

	return s
}

func makeStores(logger zerolog.Logger, databaseURI string, inMemory bool) (ports.UserStore, ports.RecordStore) {
	if inMemory {
		return user_store.NewInMemory(), record_store.NewInMemory()
	} else {
		db, err := db.NewDB(databaseURI)
		if err != nil {
			logger.Fatal().Err(err).Msg("failed to initiate database")
		}

		// TODO: create the record store with DB
		return user_store.NewWithDB(db), record_store.NewInMemory()
	}
}

func handleInterrupt() <-chan os.Signal {
	out := make(chan os.Signal, 2)
	signal.Notify(out, os.Interrupt)
	signal.Notify(out, syscall.SIGTERM)
	return out
}
