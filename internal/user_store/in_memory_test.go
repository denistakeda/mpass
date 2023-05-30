package user_store

import (
	"context"
	"testing"
	"time"

	"github.com/denistakeda/mpass/internal/domain"
	"github.com/stretchr/testify/assert"
)

func Test_inMemoryUserStore_AddNewUser(t *testing.T) {
	t.Run("add new user", func(t *testing.T) {
		s := NewInMemory()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		assert.NoError(t, s.AddNewUser(ctx, "login", "password"))
	})

	t.Run("add new user twice should fail", func(t *testing.T) {
		s := NewInMemory()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		assert.NoError(t, s.AddNewUser(ctx, "login", "password"))
		assert.Error(t, s.AddNewUser(ctx, "login", "password"))
	})
}

func Test_inMemoryUserStore_GetUser(t *testing.T) {
	t.Run("request non-existed user", func(t *testing.T) {
		s := NewInMemory()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		_, err := s.GetUser(ctx, "login")

		assert.Error(t, err)
	})

	t.Run("request existed user", func(t *testing.T) {
		s := NewInMemory()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		user := domain.User{Login: "login", PasswordHash: "password-hash"}

		assert.NoError(t, s.AddNewUser(ctx, user.Login, user.PasswordHash))

		got, err := s.GetUser(ctx, "login")

		assert.NoError(t, err)
		assert.Equal(t, user, got, "should return correct user")
	})
}
