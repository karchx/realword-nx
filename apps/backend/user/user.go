package user

import "github.com/karchx/realword-nx/model"

type Store interface {
	Create(*model.User) error
}
