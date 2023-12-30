package article

import (
	"github.com/karchx/realword-nx/model"
	uuid "github.com/satori/go.uuid"
)

type Store interface {
	GetBySlug(string) (*model.Article, error)
	GetUserArticleBySlug(userID uuid.UUID, slug string) (*model.Article, error)
	CreateArticle(*model.Article) error
	UpdateArticle(*model.Article) error
}
