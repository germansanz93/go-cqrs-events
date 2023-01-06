package search

import (
	"context"
	"germansanz93/go/cqrs/models"
)

type SearchRepository interface {
	IndexFeed(ctx context.Context, feed *models.Feed) error
	SearchFeed(ctx context.Context, query string) ([]models.Feed, error)
	Close()
}

var repo SearchRepository

func SetSearchRepository(r SearchRepository) {
	repo = r
}

func IndexFeed(ctx context.Context, feed *models.Feed) error {
	return repo.IndexFeed(ctx, feed)
}

func SearchFeed(ctx context.Context, query string) ([]models.Feed, error) {
	return repo.SearchFeed(ctx, query)
}

func Close() {
	repo.Close()
}
