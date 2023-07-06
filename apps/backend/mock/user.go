package mock

import (
	"github.com/karchx/realword-nx/conduit"
)

type UserService struct {
	CreateUserFn     func(conduit.User) error
	AuthenticateFn   func() *conduit.User
	UserByEmailFn    func(string) *conduit.User
	UserByUsernameFn func(string) (*conduit.User, error)
}

func (m *UserService) CreateUser(user conduit.User) error {
	return m.CreateUserFn(user)
}

func (m *UserService) Authenticate(email, password string) (*conduit.User, error) {
	return m.AuthenticateFn(), nil
}

func (m *UserService) UserByEmail(email string) (*conduit.User, error) {
	return m.UserByEmailFn(email), nil
}

func (m *UserService) UserByUsername(username string) (*conduit.User, error) {
  return m.UserByUsernameFn(username)
}
