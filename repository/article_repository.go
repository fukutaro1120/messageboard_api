package repository

import (
	"context"
	"message-board-api/model"
)

// ArticleRepository ...
type ArticleRepository interface {
	List(context.Context) ([]*model.Article, error)
	Create(context.Context, *model.Article) error
	Delete(context.Context, *model.Article) error
	Update(context.Context, *model.Article) error
}
