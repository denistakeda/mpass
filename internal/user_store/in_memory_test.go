package user_store

import (
	"context"
	"testing"
	"time"

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
