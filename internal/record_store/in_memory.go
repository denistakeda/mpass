package record_store

import (
	"context"

	"github.com/denistakeda/mpass/internal/domain/record"
)

type recordStore struct {
}

func NewInMemory() *recordStore {
	return &recordStore{}
}

func (r *recordStore) AddRecords(ctx context.Context, login string, records []record.Record) error {
	// TODO: implement
	return nil
}
