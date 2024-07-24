package fetcher

import (
	"app/internal/storage/model"
	"context"
	"time"
)

type ArticleStorage interface{
	Store(ctx context.Context, article model.Article) error
}

type SourseProvider interface {
	Store(ctx context.Context) ([]model.Source, error)
}

type Fetcher struct {
	articles ArticleStorage
	sources SourseProvider

	fetchInterval time.Duration
	filteredKeywords []string
}