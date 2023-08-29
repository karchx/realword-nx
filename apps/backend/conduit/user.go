package conduit

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

var AnonymousUser User

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Email        string    `gorm:"uniqueIndex" json:"email,omitempty"`
	Username     string    `gorm:"uniqueIndex" json:"username,omitempty"`
	Bio          string    `json:"bio,omitempty"`
	Image        string    `json:"image,omitempty"`
	Token        string    `json:"token,omitempty"`
	PasswordHash string    `json:"-" db:"password_has"`
	CreatedAt    time.Time `json:"-" db:"created_at"`
	UpdatedAt    time.Time `json:"-" db:"updated_at"`
}

type Following struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	FollowingID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"followingId" db:"following_id"`
	Following   *User     `gorm:"foreignKey:FollowingID" json:"following"`
	FollowerID  uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"followerId" db:"follower_id"`
	Follower    *User     `gorm:"foreignKey:FollowerID" json:"follower"`
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
	ProfileWithFollow(*User, *User) *Profile
}
