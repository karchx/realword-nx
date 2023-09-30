package conduit

import (
	"time"

	log "github.com/gothew/l-og"
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
	Followers    []Follow  `gorm:"foreignKey:FollowingID"`
	Followings   []Follow  `gorm:"foreignKey:FollowerID"`
	CreatedAt    time.Time `json:"-" db:"created_at"`
	UpdatedAt    time.Time `json:"-" db:"updated_at"`
}

type Follow struct {
	Following   User
	FollowingID uuid.UUID `gorm:"type:uuid;primaryKey" json:"followingId" db:"following_id"`
	Follower    User
	FollowerID  uuid.UUID `gorm:"type:uuid;" json:"followerId" db:"follower_id"`
	CreatedAt   time.Time `json:"-" db:"created_at"`
	UpdatedAt   time.Time `json:"-" db:"updated_at"`
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

func (user *User) ProfileWithFollow(u *User) *Profile {
	return &Profile{
		Username:  user.Username,
		Bio:       user.Bio,
		Image:     user.Image,
		Following: u.IsFollowing(user),
	}
}

func (me *User) IsFollowing(user *User) bool {
  log.Info(user)
  if user.Followers == nil {
    return false
  }

  for _, u := range user.Followers {
    if me.Username == u.Follower.Username {
      return true
    }
  }
  return false
}

/*func (me *User) IsFollowing(id uuid.UUID) bool {
	if me.Followers == nil {
		return false
	}

	for _, f := range me.Followers {
		if f.FollowerID == id {
			return true
		}
	}
	return false
}*/

type UserService interface {
	Authenticate(email, password string) (*User, error)
	CreateUser(User) error
	UserByEmail(string) (*User, error)
	UserByUsername(string) (*User, error)
	AddFollower(*User, uuid.UUID) error
	RemoveFollower(*User, uuid.UUID) error
}
