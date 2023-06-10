package record_service

import (
	"context"

	"github.com/denistakeda/mpass/internal/domain/record"
	"github.com/denistakeda/mpass/internal/ports"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type recordService struct {
	logger      zerolog.Logger
	recordStore ports.RecordStore
}

func New(logService ports.LogService, recordStore ports.RecordStore) *recordService {
	return &recordService{
		logger:      logService.ComponentLogger("recordService"),
		recordStore: recordStore,
	}
}

func (r *recordService) AddRecords(ctx context.Context, login string, records []record.Record) error {
	if err := r.recordStore.AddRecords(ctx, login, records); err != nil {
		return errors.Wrapf(err, "failed to store records for user %q", login)
	}

	r.logger.Info().Str("login", login).Msgf("%d records were sucessfully stored", len(records))

	return nil
}

func (r *recordService) AllRecords(ctx context.Context, login string) ([]record.Record, error) {
	records, err := r.recordStore.AllRecords(ctx, login)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch records")
	}

	return records, nil
}
