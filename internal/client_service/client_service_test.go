package client_service

import (
	"context"
	"testing"
	"time"

	"github.com/denistakeda/mpass/internal/auth_service"
	"github.com/denistakeda/mpass/internal/client_storage"
	"github.com/denistakeda/mpass/internal/db"
	"github.com/denistakeda/mpass/internal/domain/record"
	"github.com/denistakeda/mpass/internal/grpc_client"
	"github.com/denistakeda/mpass/internal/logging"
	"github.com/denistakeda/mpass/internal/record_service"
	"github.com/denistakeda/mpass/internal/record_store"
	"github.com/denistakeda/mpass/internal/server"
	"github.com/denistakeda/mpass/internal/user_store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestTest(t *testing.T) {
	// -- SETUP --

	// initiate postgres container
	container, err := postgres.RunContainer(context.Background(),
		testcontainers.WithImage("docker.io/postgres:15.2-alpine"),
		postgres.WithDatabase("postgres"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(wait.ForLog("database system is ready to accept connections").WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	require.NoError(t, err, "failed to start postgres")

	container.Start(context.Background())
	stopTime := time.Second
	defer container.Stop(context.Background(), &stopTime)

	databaseURI, err := container.ConnectionString(context.Background(), "sslmode=disable")

	require.NoError(t, err, "failed to get database connections string")

	// Run server
	logService := logging.New()

	// Stores
	db, err := db.NewDB(databaseURI, "file://../../migrations")
	require.NoError(t, err, "failed to initiate database")

	userStore := user_store.NewWithDB(db)
	recordStore := record_store.NewWithDb(db)

	// Services
	authService := auth_service.New(auth_service.NewAuthServiceParams{
		Secret: "secret",

		LogService: logService,
		UserStore:  userStore,
	})

	recordService := record_service.New(logService, recordStore)

	s := server.New(server.NewServerParams{
		Host:          ":3200",
		LogService:    logService,
		AuthService:   authService,
		RecordService: recordService,
	})
	s.Start()
	defer s.Stop()

	// setup client service
	grpcClient := grpc_client.New(":3200")
	defer grpcClient.Close()

	clientStorage := client_storage.NewInMemory("~/.mpass/state.gob")
	defer clientStorage.Close()

	clientService := New(clientStorage, grpcClient)

	// -- TEST DATA --

	defaultUserLogin, defaultUserPassword := "login", "password"

	// -- TESTS --

	t.Run("create a user", func(t *testing.T) {
		err := clientService.RegisterUser(defaultUserLogin, defaultUserPassword)
		assert.NoError(t, err, "failed to create a user")
	})

	t.Run("login", func(t *testing.T) {
		err := clientService.LoginUser(defaultUserLogin, defaultUserPassword)
		assert.NoError(t, err, "failed to login")
	})

	t.Run("create login-password record", func(t *testing.T) {
		err := clientService.SetRecord(record.NewLoginPasswordRecord("test-login", "test-password"))
		assert.NoError(t, err, "failed to create a login-password record")
	})

	t.Run("create bank card record", func(t *testing.T) {
		err := clientService.SetRecord(record.NewBankCardRecord("1234123412341234", 1, 1, 123))
		assert.NoError(t, err, "failed to create a bank card record")
	})

	t.Run("create text record", func(t *testing.T) {
		err := clientService.SetRecord(record.NewTextRecord("just-a-text", "text text text"))
		assert.NoError(t, err, "failed to create a text record")
	})

	t.Run("sync with the server", func(t *testing.T) {
		err := clientService.Sync()
		assert.NoError(t, err, "failed to create a login-password record")
	})
}
