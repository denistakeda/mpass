package grpc_client

import (
	"github.com/denistakeda/mpass/proto"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type grpcClient struct {
	host string
	conn *grpc.ClientConn
	c    proto.MpassServiceClient
}

func New(host string) *grpcClient {
	return &grpcClient{host: host}
}

func (gc *grpcClient) GetClient() (proto.MpassServiceClient, error) {
	conn, err := grpc.Dial(gc.host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to instantiate a connection to %q", gc.host)
	}
	gc.conn = conn
	gc.c = proto.NewMpassServiceClient(conn)

	return gc.c, nil
}

func (gc *grpcClient) Close() {
	if gc.conn != nil {
		gc.conn.Close()
	}
}
