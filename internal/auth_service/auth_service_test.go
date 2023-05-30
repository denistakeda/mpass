package auth_service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/denistakeda/mpass/internal/domain"
	"github.com/denistakeda/mpass/internal/logging"
	auth_service_mock "github.com/denistakeda/mpass/mocks/auth_service"
	"github.com/golang/mock/gomock"
)

func Test_authService_SignUp(t *testing.T) {
	type args struct {
		login    string
		password string
	}
	tests := []struct {
		name                  string
		args                  args
		userStoreExpectations func(us *auth_service_mock.MockuserStore)
		wantNotEmpty          bool
		wantErr               bool
	}{
		{
			name: "empty login",
			args: args{
				login:    "",
				password: "password",
			},
			wantNotEmpty: false,
			wantErr:      true,
		},
		{
			name: "empty password",
			args: args{
				login:    "login",
				password: "",
			},
			wantNotEmpty: false,
			wantErr:      true,
		},
		{
			name: "fail to store",
			args: args{
				login:    "login",
				password: "password",
			},
			wantNotEmpty: false,
			wantErr:      true,
			userStoreExpectations: func(us *auth_service_mock.MockuserStore) {
				us.EXPECT().
					AddNewUser(gomock.Any(), "login", gomock.Any()).
					Return(errors.New("mock error")).
					Times(1)
			},
		},
		{
			name: "stored successfully",
			args: args{
				login:    "login",
				password: "password",
			},
			wantNotEmpty: true,
			wantErr:      false,
			userStoreExpectations: func(us *auth_service_mock.MockuserStore) {
				us.EXPECT().
					AddNewUser(gomock.Any(), "login", gomock.Any()).
					Return(nil).
					Times(1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			userStore := auth_service_mock.NewMockuserStore(ctrl)

			a := New(NewAuthServiceParams{
				Secret: "secret",

				LogService: logging.New(),
				UserStore:  userStore,
			})

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			if tt.userStoreExpectations != nil {
				tt.userStoreExpectations(userStore)
			}

			got, err := a.SignUp(ctx, tt.args.login, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("authService.SignUp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantNotEmpty && got == "" {
				t.Errorf("result is empty. expected not empty result")
			}
		})
	}
}

func Test_authService_SignIn(t *testing.T) {
	type args struct {
		login    string
		password string
	}
	tests := []struct {
		name                  string
		args                  args
		userStoreExpectations func(us *auth_service_mock.MockuserStore)
		wantNotEmpty          bool
		wantErr               bool
	}{
		{
			name: "empty login",
			args: args{
				login:    "",
				password: "password",
			},
			wantNotEmpty: false,
			wantErr:      true,
		},
		{
			name: "empty password",
			args: args{
				login:    "login",
				password: "",
			},
			wantNotEmpty: false,
			wantErr:      true,
		},
		{
			name: "fail to get user",
			args: args{
				login:    "login",
				password: "password",
			},
			wantNotEmpty: false,
			wantErr:      true,
			userStoreExpectations: func(us *auth_service_mock.MockuserStore) {
				us.EXPECT().
					GetUser(gomock.Any(), "login").
					Return(domain.User{}, errors.New("mock error")).
					Times(1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			userStore := auth_service_mock.NewMockuserStore(ctrl)

			a := New(NewAuthServiceParams{
				Secret: "secret",

				LogService: logging.New(),
				UserStore:  userStore,
			})

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			if tt.userStoreExpectations != nil {
				tt.userStoreExpectations(userStore)
			}

			got, err := a.SignIn(ctx, tt.args.login, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("authService.SignUp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantNotEmpty && got == "" {
				t.Errorf("result is empty. expected not empty result")
			}
		})
	}
}
