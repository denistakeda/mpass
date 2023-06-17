package ports

import (
	"context"

	"github.com/denistakeda/mpass/internal/domain"
	"github.com/denistakeda/mpass/internal/domain/record"
	"github.com/rs/zerolog"
)

type (
	LogService interface {
		ComponentLogger(component string) zerolog.Logger
	}

	UserStore interface {
		AddNewUser(ctx context.Context, login, passwordHash string) error
		GetUser(ctx context.Context, login string) (domain.User, error)
	}

	RecordStore interface {
		AddRecords(ctx context.Context, login string, records []record.Record) error
		AllRecords(ctx context.Context, login string) ([]record.Record, error)
	}
)
