package mock

import (
	"context"

	"github.com/karchx/realword-nx/conduit"
)

type UserService struct {
  CreateUserFn func(*conduit.User) error
}

func (m *UserService) CreateUser(_ context.Context, user *conduit.User) error {
  return m.CreateUserFn(user)
}
