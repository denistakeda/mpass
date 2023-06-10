package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/denistakeda/mpass/internal/config"
	"github.com/denistakeda/mpass/internal/domain/record"
	"github.com/denistakeda/mpass/internal/logging"
	"github.com/denistakeda/mpass/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	loginPasswordRecord = &proto.Record{
		Id:             "1",
		LastUpdateDate: timestamppb.Now(),

		Record: &proto.Record_LoginPasswordRecord{
			LoginPasswordRecord: &proto.LoginPasswordRecord{
				Login:    "login",
				Password: "password",
			},
		},
	}

	textRecord = &proto.Record{
		Id:             "2",
		LastUpdateDate: timestamppb.Now(),

		Record: &proto.Record_TextRecord{
			TextRecord: &proto.TextRecord{Text: "just a text"},
		},
	}

	binaryRecord = &proto.Record{
		Id:             "3",
		LastUpdateDate: timestamppb.Now(),

		Record: &proto.Record_BinaryRecord{
			BinaryRecord: &proto.BinaryRecord{Binary: []byte("binary text")},
		},
	}

	bankCardRecord = &proto.Record{
		Id:             "4",
		LastUpdateDate: timestamppb.Now(),

		Record: &proto.Record_BankCardRecord{
			BankCardRecord: &proto.BankCardRecord{
				CardCode: "1234 1234 1234 1234",
				Month:    1,
				Day:      1,
				Code:     123,
			},
		},
	}
)

func Test_SignUp(t *testing.T) {
	serverTest(t, "successful sign up", func(t *testing.T, c proto.MpassServiceClient) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		req := proto.SignUpRequest{
			Login:    "login",
			Password: "password",
		}

		resp, err := c.SignUp(ctx, &req)

		assert.NoError(t, err, "should successfully create a user")
		assert.NotEmpty(t, resp.Token)
	})
}

func Test_SignIn(t *testing.T) {
	serverTest(t, "successfull sign in", func(t *testing.T, c proto.MpassServiceClient) {
		// Sign Up
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		signUpReq := proto.SignUpRequest{
			Login:    "login",
			Password: "password",
		}

		signUpResp, err := c.SignUp(ctx, &signUpReq)

		assert.NoError(t, err, "should successfully create a user")
		assert.NotEmpty(t, signUpResp.Token)

		// And then Sign In
		signInReq := proto.SignInRequest{
			Login:    "login",
			Password: "password",
		}

		signInResp, err := c.SignIn(ctx, &signInReq)
		assert.NoError(t, err, "should successfully sign in")
		assert.NotEmpty(t, signInResp.Token)
	})
}

func Test_AddRecords(t *testing.T) {
	serverTest(t, "should require authentication", func(t *testing.T, c proto.MpassServiceClient) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		_, err := c.AddRecords(ctx, &proto.AddRecordsRequest{
			Records: []*proto.Record{},
		})

		assert.Error(t, err)
		if err != nil {
			e, ok := status.FromError(err)
			assert.True(t, ok, "should return error with a status")
			if ok {
				assert.Equalf(t, codes.Unauthenticated, e.Code(), "should return Unauthenticated status code")
			}
		}
	})

	serverTest(t, "add single record", func(t *testing.T, c proto.MpassServiceClient) {
		ctx := authorisedContext(t, c, "login", "password")

		_, err := c.AddRecords(ctx, &proto.AddRecordsRequest{
			Records: []*proto.Record{
				loginPasswordRecord, textRecord, binaryRecord, bankCardRecord,
			},
		})

		assert.NoError(t, err)
	})
}

func Test_AllRecords(t *testing.T) {
	serverTest(t, "should require authentication", func(t *testing.T, c proto.MpassServiceClient) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		_, err := c.AllRecords(ctx, &empty.Empty{})

		assert.Error(t, err)
		if err != nil {
			e, ok := status.FromError(err)
			assert.True(t, ok, "should return error with a status")
			if ok {
				assert.Equalf(t, codes.Unauthenticated, e.Code(), "should return Unauthenticated status code")
			}
		}
	})

	serverTest(t, "get the list of records", func(t *testing.T, c proto.MpassServiceClient) {
		ctx := authorisedContext(t, c, "login", "password")

		records := []*proto.Record{
			loginPasswordRecord, textRecord, binaryRecord, bankCardRecord,
		}

		_, err := c.AddRecords(ctx, &proto.AddRecordsRequest{
			Records: records,
		})

		assert.NoError(t, err)

		resp, err := c.AllRecords(ctx, &empty.Empty{})

		assert.NoError(t, err)

		compareRecords(t, records, resp.Records, "store should include subset")
	})
}

// -- Test helpers --

// serverTest creates the environment for testing the server.
// It creates and runs the server before the test and then stops it afterwards.
// Also it provides a sat up client to use with the defined server.
func serverTest(t *testing.T, description string, f func(*testing.T, proto.MpassServiceClient)) {
	logService := logging.New()
	conf := config.Config{
		Host:   ":0",
		Secret: "secret",
	}

	s := buildServer(buildParams{
		conf:                conf,
		logService:          logService,
		useInMemoryStorages: true,
	})
	s.Start()
	defer s.Stop()

	conn, err := grpc.Dial(s.Host(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		assert.Fail(t, "failed to create a client")
	}
	defer conn.Close()

	c := proto.NewMpassServiceClient(conn)

	t.Run(description, func(t *testing.T) {
		f(t, c)
	})
}

func authorisedContext(t *testing.T, c proto.MpassServiceClient, login, password string) context.Context {
	req := proto.SignUpRequest{
		Login:    "login",
		Password: "password",
	}

	resp, err := c.SignUp(context.Background(), &req)
	require.NoError(t, err, "failed to sign up to the server")

	md := metadata.New(map[string]string{"authorization": fmt.Sprintf("Bearer %s", resp.Token)})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	return ctx
}

func compareRecords(t *testing.T, subset []*proto.Record, store []*proto.Record, label string) {
	recStore := make([]record.Record, len(store))
	for idx, item := range store {
		recStore[idx] = record.FromProto(item)
	}

	recSubset := make([]record.Record, len(subset))
	for idx, item := range store {
		recSubset[idx] = record.FromProto(item)
	}

	assert.Subset(t, recStore, recSubset, label)
}
