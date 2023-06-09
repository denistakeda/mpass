package record_store

import (
	"context"
	"sync"

	"github.com/denistakeda/mpass/internal/domain/record"
)

type inMemory struct {
	stores sync.Map
}

func NewInMemory() *inMemory {
	return &inMemory{}
}

func (r *inMemory) AddRecords(ctx context.Context, login string, records []record.Record) error {
	return r.getStore(login).addRecords(records)
}

func (r *inMemory) AllRecords(ctx context.Context, login string) ([]record.Record, error) {
	return r.getStore(login).allRecords(), nil
}

func (r *inMemory) getStore(login string) *store {
	s, _ := r.stores.LoadOrStore(login, newStore())
	return s.(*store)
}

// -- Store --

type store struct {
	mx      sync.Mutex
	records map[string]record.Record
}

func newStore() *store {
	return &store{records: make(map[string]record.Record)}
}

func (s *store) addRecords(records []record.Record) error {
	s.mx.Lock()
	defer s.mx.Unlock()

	for _, rec := range records {
		oldRec, ok := s.records[rec.GetId()]
		if !ok || rec.GetLastUpdateDate().After((oldRec).GetLastUpdateDate()) {
			s.records[rec.GetId()] = rec
		}
	}

	return nil
}

func (s *store) allRecords() []record.Record {
	s.mx.Lock()
	defer s.mx.Unlock()

	res := make([]record.Record, 0, len(s.records))
	for _, rec := range s.records {
		res = append(res, rec)
	}

	return res
}
