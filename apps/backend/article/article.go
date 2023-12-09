package article

import "github.com/karchx/realword-nx/model"

type Store interface {
	CreateArticle(*model.Article) error
}
