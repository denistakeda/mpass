package user_store

import (
	"context"
	"sync"

	"github.com/pkg/errors"
)

type (
	inMemoryUserStore struct {
		users sync.Map
	}

	user struct {
		login        string
		passwordHash string
	}
)

func NewInMemory() *inMemoryUserStore {
	return &inMemoryUserStore{}
}

func (s *inMemoryUserStore) AddNewUser(ctx context.Context, login, passwordHash string) error {
	_, loaded := s.users.LoadOrStore(login, user{login: login, passwordHash: passwordHash})
	if loaded {
		return errors.Errorf("user with login %q already exists", login)
	}

	return nil
}
