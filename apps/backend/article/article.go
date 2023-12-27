package article

import "github.com/karchx/realword-nx/model"

type Store interface {
	GetBySlug(string) (*model.Article, error)
	CreateArticle(*model.Article) error
}
