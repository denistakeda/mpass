package server

import (
	"context"
	"net"

	"github.com/denistakeda/mpass/internal/ports"
	pb "github.com/denistakeda/mpass/proto"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type (
	server struct {
		pb.UnimplementedMpassServiceServer

		logger      zerolog.Logger
		authService authService

		host string
		s    *grpc.Server
	}

	authService interface {
		SignUp(ctx context.Context, login, password string) (string, error)
	}
)

type NewServerParams struct {
	Host        string
	LogService  ports.LogService
	AuthService authService
}

func New(params NewServerParams) *server {
	return &server{
		host:        params.Host,
		logger:      params.LogService.ComponentLogger("server"),
		authService: params.AuthService,
	}
}

func (s *server) Start() <-chan error {
	out := make(chan error, 2)

	listen, err := net.Listen("tcp", s.host)
	if err != nil {
		out <- errors.Wrapf(err, "can not start gRPC server on a host %q", s.host)
		return out
	}

	s.s = grpc.NewServer()
	pb.RegisterMpassServiceServer(s.s, s)

	s.logger.Info().Msg("gRPC server started")

	go func() {
		if err := s.s.Serve(listen); err != nil {
			out <- errors.Wrap(err, "gRPC server has failed")
		}
	}()

	return out
}

func (s *server) Stop() {
	s.s.Stop()
	s.logger.Info().Msg("gRPC server stopped")
}

func (s *server) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	var resp pb.SignUpResponse

	token, err := s.authService.SignUp(ctx, req.Login, req.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create a user")
	}

	resp.Token = token

	return &resp, nil
}
