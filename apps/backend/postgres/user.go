package postgres

import (
	"github.com/karchx/realword-nx/conduit"
	uuid "github.com/satori/go.uuid"
)

type UserService struct {
	db *DB
}

func NewUserService(db *DB) *UserService {
	return &UserService{db}
}

func (us *UserService) UserByEmail(email string) (*conduit.User, error) {
	var user conduit.User
	if err := us.db.Where(&conduit.User{Email: email}).First(&user).Error; err != nil {
		return nil, conduit.ErrNotFound
	}

	return &user, nil
}

func (us *UserService) UserByUsername(username string) (*conduit.User, error) {
	var user conduit.User

	if err := us.db.Where(&conduit.User{Username: username}).First(&user).Error; err != nil {
		return nil, conduit.ErrNotFound
	}

	return &user, nil
}

func (us *UserService) CreateUser(user conduit.User) error {
	return createUser(user, us)
}

func (us *UserService) Authenticate(email, password string) (*conduit.User, error) {
	user, err := us.UserByEmail(email)

	if err != nil {
		return nil, err
	}

	if !user.VerifyPassword(password) {
		return nil, conduit.ErrUnAuthorized
	}

	return user, nil
}

func (us *UserService) AddFollower(user *conduit.User, followerID uuid.UUID) error {
	return us.db.Model(user).Association("Followers").Append(&conduit.Follow{FollowerID: followerID, FollowingID: user.ID})
}

func (us *UserService) RemoveFollower(user *conduit.User, followerID uuid.UUID) error {
	f := conduit.Follow{
		FollowerID:  followerID,
		FollowingID: user.ID,
	}

	if err := us.db.Model(user).Association("Followers").Find(&f); err != nil {
		return err
	}

	if err := us.db.Delete(f).Error; err != nil {
		return err
	}

	return nil
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
