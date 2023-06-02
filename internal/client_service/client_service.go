package client_service

import (
	"github.com/denistakeda/mpass/internal/domain/record"
	"github.com/denistakeda/mpass/proto"
	"github.com/pkg/errors"
)

type (
	clientService struct {
		clientStorage clientStorage
		grpcClient    grpcClient
	}

	clientStorage interface {
		SetRecord(record.Record) error
		GetRecord(string) (record.Record, error)
	}

	grpcClient interface {
		GetClient() (proto.MpassServiceClient, error)
	}
)

func New(clientStorage clientStorage, grpcClient grpcClient) *clientService {
	return &clientService{clientStorage: clientStorage, grpcClient: grpcClient}
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
