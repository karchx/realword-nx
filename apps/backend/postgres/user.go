package postgres

import (
	"context"
	"database/sql"

	"github.com/karchx/realword-nx/conduit"
)

type UserService struct {
	db *DB
}

func NewUserService(db *DB) *UserService {
	return &UserService{db}
}

func (us *UserService) CreateUser(user conduit.User) error {
	result := us.db.Create(&user)
	if result.RowsAffected == 0 {
		return conduit.ErrNotCreated
	}

	return nil
}

func createUser(ctx context.Context, tx *sql.Tx, user *conduit.User) error {
	query := `
	INSERT INTO users (email, username, bio, image, password_hash)
	VALUES($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at
	`
	args := []interface{}{user.Email, user.Username, user.Bio, user.Image, user.PasswordHash}
	err := tx.QueryRowContext(ctx, query, args...).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return conduit.ErrDuplicateEmail
		case err.Error() == `pq: duplicate key value violates unique constraint "users_username_key"`:
			return conduit.ErrDuplicateUsername
		default:
			return err
		}
	}

	return nil
}
