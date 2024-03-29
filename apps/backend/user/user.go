package user

import (
	"github.com/karchx/realword-nx/model"
	uuid "github.com/satori/go.uuid"
)

type Store interface {
	GetByID(uuid.UUID) (*model.User, error)
	GetByEmail(string) (*model.User, error)
	GetByUsername(string) (*model.User, error)
	Create(*model.User) error
	Update(*model.User) error
	AddFollower(user *model.User, followerID uuid.UUID) error
	RemoveFollower(user *model.User, followerID uuid.UUID) error
	IsFollower(userID, followerID uuid.UUID) (bool, error)
}
