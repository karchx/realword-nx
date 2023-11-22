package model

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID         uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Username   string    `gorm:"uniqueIndex;not null"`
	Email      string    `gorm:"uniqueIndex;not null"`
	Password   string    `gorm:"not null"`
	Bio        *string
	Image      *string
	Followers  []Follow `gorm:"foreignKey:FollowingID"`
	Followings []Follow `gorm:"foreignKey:FollowerID"`
}

type Follow struct {
	Follower    User
	FollowerID  uuid.UUID `gorm:"primaryKey;type:uuid"`
	Following   User
	FollowingID uuid.UUID `gorm:"primaryKey;type:uuid"`
}

func (u *User) HashPassword(plain string) (string, error) {
	if len(plain) == 0 {
		return "", errors.New("password should not be empty")
	}

	h, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)

	return string(h), err
}

func (u *User) CheckPassword(plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plain))
	return err == nil
}
