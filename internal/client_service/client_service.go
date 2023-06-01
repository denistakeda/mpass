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
