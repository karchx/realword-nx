package store

import (
	"errors"

	"github.com/karchx/realword-nx/model"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type ArticleStore struct {
	db *gorm.DB
}

func NewArticleStore(db *gorm.DB) *ArticleStore {
	return &ArticleStore{
		db: db,
	}
}

func (as *ArticleStore) GetBySlug(s string) (*model.Article, error) {
	var m model.Article

	err := as.db.Where(&model.Article{Slug: s}).
		Preload("Tags", func(db *gorm.DB) *gorm.DB {
			return db.Order("tag asc")
		}).
		Preload("Author").First(&m).Error

	return &m, err
}

func (as *ArticleStore) GetUserArticleBySlug(userID uuid.UUID, slug string) (*model.Article, error) {
	var m model.Article
	err := as.db.Where(&model.Article{Slug: slug, AuthorId: userID}).First(&m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &m, err
}

func (as *ArticleStore) CreateArticle(a *model.Article) error {
	tags := a.Tags
	tx := as.db.Begin()
	if err := tx.Create(&a).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, t := range a.Tags {
		err := tx.Where(&model.Tag{Tag: t.Tag}).First(&t).Error
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return err
		}
		if err := tx.Model(&a).Association("Tags").Append(&t); err != nil {
			tx.Rollback()
			return err
		}
	}

	a.Tags = tags

	return tx.Commit().Error
}

func (as *ArticleStore) UpdateArticle(a *model.Article) error {
	tx := as.db.Begin()
	if err := tx.Updates(&a).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (as *ArticleStore) DeleteArticle(a *model.Article) error {
	return as.db.Delete(a).Error
}

func (as *ArticleStore) UndoDeleteArticle(slug string) error {
	return as.db.Model(&model.Article{}).Unscoped().Where(&model.Article{Slug: slug}).Update("deleted_at", nil).Error
}
