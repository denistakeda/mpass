package server

import (
	"net"

	"github.com/denistakeda/mpass/internal/ports"
	pb "github.com/denistakeda/mpass/proto"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

type (
	server struct {
		pb.UnimplementedMpassServiceServer

		logger zerolog.Logger

		host    string
		storage storage
		s       *grpc.Server
	}

	storage interface {
	}
)

type NewServerParams struct {
	Host       string
	Storage    storage
	LogService ports.LogService
}

func New(params NewServerParams) *server {
	return &server{
		host:    params.Host,
		storage: params.Storage,
		logger:  params.LogService.ComponentLogger("server"),
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
