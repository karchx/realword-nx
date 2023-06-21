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

type UserFilter struct {
	ID       *int
	Email    *string
	Username *string

	Limit  int
	Offset int
}

type Profile struct {
	Username  string `json:"username"`
	Bio       string `json:"bio"`
	Image     string `json:"image"`
	Following bool   `json:"following"`
}

func (user *User) SetPassword(password string) error {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.PasswordHash = string(hashBytes)

	return nil
}

func (user *User) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))

	return err == nil
}

func (user *User) Profile() *Profile {
	return &Profile{
		Username: user.Username,
		Bio:      user.Bio,
		Image:    user.Image,
	}
}

type UserService interface {
	Authenticate(email, password string) (*User, error)
	CreateUser(User) error
	UserByEmail(string) (*User, error)
	UserByUsername(string) (*User, error)
}
