package conduit

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `gorm:"primaryKey" json:"-"`
	Email    string `gorm:"uniqueIndex" json:"email,omitempty"`
	Username string `gorm:"uniqueIndex" json:"username,omitempty"`
	Bio      string `json:"bio,omitempty"`
	Image    string `json:"image,omitempty"`
	Token    string `json:"token,omitempty"`
	//Following    []*User   `json: "-"`
	//Followers    []*User   `json: "-"`
	PasswordHash string    `json:"-" db:"password_has"`
	CreatedAt    time.Time `json:"-" db:"created_at"`
	UpdatedAt    time.Time `json:"-" db:"updated_at"`
}

func (user *User) SetPassword(password string) error {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.PasswordHash = string(hashBytes)

	return nil
}

type UserService interface {
	CreateUser(User) error
}
