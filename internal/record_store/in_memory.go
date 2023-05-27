package record_store

import (
	"context"
	"sync"

	"github.com/denistakeda/mpass/internal/domain/record"
)

type recordStore struct {
	stores sync.Map
}

func NewInMemory() *recordStore {
	return &recordStore{}
}

func (r *recordStore) AddRecords(ctx context.Context, login string, records []record.Record) error {
	return r.getStore(login).addRecords(records)
}

func (r *recordStore) getStore(login string) *store {
	s, _ := r.stores.LoadOrStore(login, newStore())
	return s.(*store)
}

// -- Store --

type store struct {
	mx      sync.RWMutex
	records map[string]*record.Record
}

func newStore() *store {
	return &store{records: make(map[string]*record.Record)}
}

func (s *store) addRecords(records []record.Record) error {
	s.mx.Lock()
	defer s.mx.Unlock()

	for _, rec := range records {
		oldRec, ok := s.records[rec.GetId()]
		if !ok || rec.GetLastUpdateDate().After((*oldRec).GetLastUpdateDate()) {
			s.records[rec.GetId()] = &rec
		}
	}

	return nil
}
