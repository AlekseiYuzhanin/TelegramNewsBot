package fetcher

import (
	"app/internal/storage/model"
	"app/internal/source"
	"context"
	"sync"
	"time"
)

type ArticleStorage interface{
	Store(ctx context.Context, article model.Article) error
}

type SourseProvider interface {
	Sources(ctx context.Context) ([]model.Source, error)
}

type Source interface{
	ID() int64
	Name() string
	Fetch(ctx context.Context) ([]model.Item, error)
}

type Fetcher struct {
	articles ArticleStorage
	sources SourseProvider

	fetchInterval time.Duration
	filteredKeywords []string
}

func New(
	articleStorage ArticleStorage,
	sourceProvider SourseProvider,
	fetchInterval time.Duration,
	filteredKeywords []string,
) *Fetcher{
	return &Fetcher{
		articles: articleStorage,
		sources: sourceProvider,
		fetchInterval: fetchInterval,
		filteredKeywords: filteredKeywords,
	}
}

func (f *Fetcher) Fetch(ctx context.Context) error {
	sources, err := f.sources.Sources(ctx)
	if err != nil{
		return err
	}
	var wg sync.WaitGroup

	for _, src := range sources{
		wg.Add(1)

		rssSource := source.NewRSSSourceFromModel(src)

		go func(source Source){
			
		}(src)
	}
}