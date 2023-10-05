package user

import "github.com/karchx/realword-nx/model"

type Store interface {
	GetByEmail(string) (*model.User, error)
	Create(*model.User) error
}
