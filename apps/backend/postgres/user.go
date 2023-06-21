package postgres

import (
	"github.com/karchx/realword-nx/conduit"
)

type UserService struct {
	db *DB
}

func NewUserService(db *DB) *UserService {
	return &UserService{db}
}

func (us *UserService) CreateUser(user conduit.User) error {
	return createUser(user, us)
}

func createUser(user conduit.User, us *UserService) error {

	result := us.db.Create(&user)

	if result.Error != nil {
		switch {
		case result.Error.Error() == `ERROR: duplicate key value violates unique constraint "idx_users_email" (SQLSTATE 23505)`:
			return conduit.ErrDuplicateEmail
		case result.Error.Error() == `ERROR: duplicate key value violates unique constraint "idx_users_username" (SQLSTATE 23505)`:
			return conduit.ErrDuplicateUsername
		default:
			return result.Error
		}
	}

	if result.RowsAffected == 0 {
		return conduit.ErrNotCreated
	}

	return nil
}
