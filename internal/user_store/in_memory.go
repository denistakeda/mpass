package user_store

import (
	"context"
	"sync"

	"github.com/denistakeda/mpass/internal/domain"
	"github.com/pkg/errors"
)

type inMemoryUserStore struct {
	users sync.Map
}

func NewInMemory() *inMemoryUserStore {
	return &inMemoryUserStore{}
}

func (s *inMemoryUserStore) AddNewUser(ctx context.Context, login, passwordHash string) error {
	_, loaded := s.users.LoadOrStore(login, domain.User{Login: login, PasswordHash: passwordHash})
	if loaded {
		return errors.Errorf("user with login %q already exists", login)
	}

	return nil
}

func (s *inMemoryUserStore) GetUser(ctx context.Context, login string) (domain.User, error) {
	user, ok := s.users.Load(login)
	if !ok {
		return domain.User{}, errors.Errorf("no such user %q", login)
	}

	return user.(domain.User), nil
}
