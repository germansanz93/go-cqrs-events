package repository

import (
	"context"
	"germansanz93/go/cqrs/models"
)

type Repository interface {
	Close()
	InserFeed(ctx context.Context, feed *models.Feed) error
	ListFeeds(ctx context.Context) ([]*models.Feed, error)
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
