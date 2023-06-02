package client_service

import (
	"context"
	"time"

	"github.com/denistakeda/mpass/internal/domain/record"
	"github.com/denistakeda/mpass/proto"
	"github.com/pkg/errors"
)

var (
	signUpTimeout = 5 * time.Second
)

type (
	clientService struct {
		clientStorage clientStorage
		grpcClient    grpcClient
	}

	clientStorage interface {
		SetRecord(record.Record) error
		GetRecord(string) (record.Record, error)
		SetToken(string) error
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

func (c *clientService) RegisterUser(login, password string) error {
	client, err := c.grpcClient.GetClient()
	if err != nil {
		return errors.Wrapf(err, "failed to register user %q", login)
	}

	ctx, cancel := context.WithTimeout(context.Background(), signUpTimeout)
	defer cancel()

	resp, err := client.SignUp(ctx, &proto.SignUpRequest{Login: login, Password: password})
	if err != nil {
		return errors.Wrapf(err, "failed to request user registration for user %q", login)
	}

	return c.clientStorage.SetToken(resp.Token)
}

func (c *clientService) LoginUser(login, password string) error {
	client, err := c.grpcClient.GetClient()
	if err != nil {
		return errors.Wrapf(err, "failed to login user %q", login)
	}

	ctx, cancel := context.WithTimeout(context.Background(), signUpTimeout)
	defer cancel()

	resp, err := client.SignIn(ctx, &proto.SignInRequest{Login: login, Password: password})
	if err != nil {
		return errors.Wrapf(err, "failed to request user login for user %q", login)
	}

	return c.clientStorage.SetToken(resp.Token)
}
