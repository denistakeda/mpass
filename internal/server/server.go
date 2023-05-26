package server

import (
	"context"
	"net"

	"github.com/denistakeda/mpass/internal/domain"
	"github.com/denistakeda/mpass/internal/domain/record"
	"github.com/denistakeda/mpass/internal/ports"
	pb "github.com/denistakeda/mpass/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type (
	server struct {
		pb.UnimplementedMpassServiceServer

		logger        zerolog.Logger
		authService   authService
		recordService recordService

		host     string
		usedHost string // provided host might differ from the actually used one
		s        *grpc.Server
	}

	authService interface {
		SignUp(ctx context.Context, login, password string) (string, error)
		SignIn(ctx context.Context, login, password string) (string, error)
		AuthenticateUser(ctx context.Context, token string) (domain.User, error)
	}

	recordService interface {
		AddRecords(ctx context.Context, login string, records []record.Record) error
	}
)

var userKey struct{}

type NewServerParams struct {
	Host          string
	LogService    ports.LogService
	AuthService   authService
	RecordService recordService
}

func New(params NewServerParams) *server {
	return &server{
		host:          params.Host,
		logger:        params.LogService.ComponentLogger("server"),
		authService:   params.AuthService,
		recordService: params.RecordService,
	}
}

func (s *server) Start() <-chan error {
	out := make(chan error, 2)

	listen, err := net.Listen("tcp", s.host)
	if err != nil {
		out <- errors.Wrapf(err, "can not start gRPC server on a host %q", s.host)
		return out
	}

	s.usedHost = listen.Addr().String()

	s.s = grpc.NewServer(grpc.UnaryInterceptor(auth.UnaryServerInterceptor(s.authFunc)))
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

func (s *server) Host() string {
	return s.usedHost
}

func (s *server) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	var resp pb.SignUpResponse

	token, err := s.authService.SignUp(ctx, req.Login, req.Password)
	if err != nil {
		s.logger.Error().Err(err).Str("login", req.Login).Msg("failed to sign up")
		return nil, status.Error(codes.Internal, "failed to sign up")
	}

	resp.Token = token

	return &resp, nil
}

func (s *server) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	var resp pb.SignInResponse

	token, err := s.authService.SignIn(ctx, req.Login, req.Password)
	if err != nil {
		s.logger.Error().Err(err).Str("login", req.Login).Msg("failed to sign in")
		return nil, status.Error(codes.Internal, "failed to sign in")
	}

	resp.Token = token

	return &resp, nil
}

func (s *server) AddRecords(ctx context.Context, req *pb.AddRecordsRequest) (*empty.Empty, error) {
	user, ok := ctx.Value(userKey).(domain.User)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "user is not authenticated")
	}

	s.recordService.AddRecords(ctx, user.Login, toDomainRecords(req.Records))

	return &empty.Empty{}, nil
}

// authFunc is used by a middleware to authenticate requests
func (s *server) authFunc(ctx context.Context) (context.Context, error) {
	token, err := auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return ctx, nil // not all endpoints require authorization
	}

	user, err := s.authService.AuthenticateUser(ctx, token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}

	return context.WithValue(ctx, userKey, user), nil
}

func toDomainRecords(recs []*pb.Record) []record.Record {
	res := make([]record.Record, 0, len(recs))

	for _, rec := range recs {
		res = append(res, record.FromProto(rec))
	}

	return res
}
