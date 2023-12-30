package model

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Slug        string `gorm:"uniqueIndex;not null"`
	Title       string `gorm:"not null"`
	Description string
	Body        string
	Author      User
	AuthorId    uuid.UUID
	Tags        []Tag `gorm:"many2many:article_tags;"`
}

type Tag struct {
	gorm.Model
	Tag      string    `gorm:"uniqueIndex"`
	Articles []Article `gorm:"many2many:article_tags;"`
}
