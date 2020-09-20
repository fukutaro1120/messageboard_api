package inmemory

import (
	"context"
	"message-board-api/model"
	"sync"
)

type InMemoryArticleRepository struct {
	mu   sync.RWMutex
	data []*model.Article
	idx  map[int64]int
	lid  int64
}

func NewArticleRepository() *InMemoryArticleRepository {
	return &InMemoryArticleRepository{
		data: make([]*model.Article, 0),
		idx:  make(map[int64]int),
		lid:  0,
	}
}

func (r *InMemoryArticleRepository) List(ctx context.Context) ([]*model.Article, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	as := make([]*model.Article, 0, len(r.data))
	for _, a := range r.data {
		as = append(as, a)
	}
	return as, nil
}

func (r *InMemoryArticleRepository) Create(ctx context.Context, a *model.Article) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// lid := r.lid + 1
	// a2 := *a
	// a2.ID = lid
	// r.data = append(r.data, &a2)
	// r.idx[lid] = len(r.data) - 1
	// r.lid = lid
	// a.ID = lid
	return nil
}
