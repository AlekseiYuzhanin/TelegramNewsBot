package fetcher

import (
	"app/internal/source"
	"app/internal/storage/model"
	"context"
	"log"
	"sync"
	"go.tomakado.io/containers/set"
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
			defer wg.Done()

			items, err := source.Fetch(ctx)
			if err != nil{
				log.Printf("[ERROR] Fetching items from source %s: %v", source.Name(), err)
				return
			}

			if err := f.processItems(ctx, source, items); err != nil{
				log.Printf("[ERROR] Processing items from source %s: %v", source.Name(),err)
				return
			}
		}(rssSource)
	}

	wg.Wait()
	return nil
}

func (f *Fetcher) processItems(ctx context.Context, source Source, items []model.Item) {
	for _, item := range items{
		item.Date = item.Date.UTC()
	}
}

func (f *Fetcher) itemShouldBeSkipped(item model.Item) bool{
	categoriesSet := set.New(item.Categories...)

	for _, keyword := range f.filteredKeywords{
		if categoriesSet.Contains(keyword){
			return true
		}
	}
}