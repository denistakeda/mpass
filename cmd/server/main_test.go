package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/denistakeda/mpass/internal/config"
	"github.com/denistakeda/mpass/internal/logging"
	"github.com/denistakeda/mpass/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
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

	s := buildServer(conf, logService)
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