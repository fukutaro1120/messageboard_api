package usecase

import (
	"context"
	"message-board-api/model"
	"message-board-api/repository"
)

// ArticleUseCase ...
type ArticleUseCase struct {
	articleRepo repository.ArticleRepository
}

// NewArticleUseCase ...
func NewArticleUseCase(articleRepo repository.ArticleRepository) *ArticleUseCase {
	return &ArticleUseCase{
		articleRepo: articleRepo,
	}
}

// List ...
func (uc *ArticleUseCase) List(ctx context.Context) ([]*model.Article, error) {
	return uc.articleRepo.List(ctx)
}

// Create ...
func (uc *ArticleUseCase) Create(ctx context.Context, sa *model.Article) error {
	return uc.articleRepo.Create(ctx, sa)
}

// Delete ...
func (uc *ArticleUseCase) Delete(ctx context.Context, sa *model.Article) error {
	return uc.articleRepo.Delete(ctx, sa)
}

// Update ...
func (uc *ArticleUseCase) Update(ctx context.Context, sa *model.Article) error {
	return uc.articleRepo.Update(ctx, sa)
}
