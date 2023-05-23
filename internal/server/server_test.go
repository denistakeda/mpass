package server

import (
	"context"
	"testing"
	"time"

	"github.com/denistakeda/mpass/internal/logging"
	"github.com/denistakeda/mpass/mocks/server"
	pb "github.com/denistakeda/mpass/proto"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func Test_server_StartStop(t *testing.T) {
	ctrl := gomock.NewController(t)

	s := New(NewServerParams{
		Host:        ":0",
		LogService:  logging.New(),
		AuthService: server_mock.NewMockauthService(ctrl),
	})

	s.Start()
	defer s.Stop()
}

func Test_server_SignUp(t *testing.T) {
	tests := []struct {
		name                    string
		req                     *pb.SignUpRequest
		authServiceExpectations func(*server_mock.MockauthService)
		wantToken               bool
		wantErr                 bool
	}{
		{
			name: "auth service returns an error",
			req:  &pb.SignUpRequest{Login: "login", Password: "password"},
			authServiceExpectations: func(as *server_mock.MockauthService) {
				as.EXPECT().
					SignUp(gomock.Any(), "login", "password").
					Return("", errors.New("mock error")).
					Times(1)
			},
			wantToken: false,
			wantErr:   true,
		},
		{
			name: "should successfully return a token",
			req:  &pb.SignUpRequest{Login: "login", Password: "password"},
			authServiceExpectations: func(as *server_mock.MockauthService) {
				as.EXPECT().
					SignUp(gomock.Any(), "login", "password").
					Return("token", nil).
					Times(1)
			},
			wantToken: true,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			authService := server_mock.NewMockauthService(ctrl)

			s := New(NewServerParams{
				Host:        ":0",
				LogService:  logging.New(),
				AuthService: authService,
			})

			if tt.authServiceExpectations != nil {
				tt.authServiceExpectations(authService)
			}

			s.Start()
			defer s.Stop()

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			got, err := s.SignUp(ctx, tt.req)
			if tt.wantToken {
				assert.NotNil(t, got)
				if got != nil {
					assert.NotEmpty(t, got.Token, "should have token")
				}
			}
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
