package client_service

import (
	"context"
	"fmt"
	"time"

	"github.com/denistakeda/mpass/internal/domain/record"
	"github.com/denistakeda/mpass/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"
)

var (
	signUpTimeout = 5 * time.Second
	syncTimeout   = 10 * time.Second
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
		GetToken() (string, error)
		ItemsToSync() ([]record.Record, error)
		SyncRecords([]record.Record) error
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

func (c *clientService) Sync() error {
	token, err := c.clientStorage.GetToken()
	if err != nil {
		return errors.Wrap(err, "failed to get user token")
	}
	if token == "" {
		return errors.New("user is not signed in, use `mpass login` first")
	}

	client, err := c.grpcClient.GetClient()
	if err != nil {
		return errors.Wrapf(err, "failed to sync")
	}

	ctx, cancel := context.WithTimeout(context.Background(), syncTimeout)
	defer cancel()

	md := metadata.New(map[string]string{"authorization": fmt.Sprintf("Bearer %s", token)})
	ctx = metadata.NewOutgoingContext(context.Background(), md)

	toSync, err := c.clientStorage.ItemsToSync()
	if err != nil {
		return err
	}

	if len(toSync) > 0 {
		var addRecordsRequest proto.AddRecordsRequest
		for _, item := range toSync {
			addRecordsRequest.Records = append(addRecordsRequest.Records, item.ToProto())
		}
		_, err = client.AddRecords(ctx, &addRecordsRequest)
		if err != nil {
			return err
		}
	}

	resp, err := client.AllRecords(ctx, &empty.Empty{})
	if err != nil {
		return err
	}
	records := make([]record.Record, 0, len(resp.Records))
	for _, item := range resp.Records {
		records = append(records, record.FromProto(item))
	}

	err = c.clientStorage.SyncRecords(records)
	if err != nil {
		return err
	}

	return nil

}
