package client_service

import (
	"github.com/denistakeda/mpass/internal/domain/record"
	"github.com/pkg/errors"
)

type (
	clientService struct {
		clientStorage clientStorage
	}

	clientStorage interface {
		SetRecord(record.Record) error
		GetRecord(string) (record.Record, error)
	}
)

func New(clientStorage clientStorage) *clientService {
	return &clientService{clientStorage: clientStorage}
}

func (c *clientService) SetRecord(r record.Record) error {
	if err := c.clientStorage.SetRecord(r); err != nil {
		return errors.Wrap(err, "failed to store record")
	}

	return nil
}

func (c *clientService) GetRecord(key string) (record.Record, error) {
	rec, err := c.clientStorage.GetRecord(key)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get record %q", key)
	}

	return rec, nil
}
