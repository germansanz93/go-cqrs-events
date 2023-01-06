package repository

import (
	"context"
	"germansanz93/go/cqrs/models"
)

type Repository interface {
	InserFeed(ctx context.Context, feed *models.Feed) error
	ListFeeds(ctx context.Context) ([]*models.Feed, error)
	Close()
}

var repository Repository

func SetRepository(r Repository) {
	repository = r
}

func Close() {
	repository.Close()
}

func InserFeed(ctx context.Context, feed *models.Feed) error {
	return repository.InserFeed(ctx, feed)
}

func ListFeeds(ctx context.Context) ([]*models.Feed, error) {
	return repository.ListFeeds(ctx)
}
