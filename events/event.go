package events

import (
	"context"
	"germansanz93/go/cqrs/models"
)

type EventStore interface {
	PublishCreatedFeed(ctx context.Context, feed *models.Feed) error
	SubscribeCreatedFeed(ctx context.Context) (<-chan CreatedFeedMessage, error)
	OnCreateFeed(f func(CreatedFeedMessage)) error
	Close()
}

var eventStore EventStore

func SetEventStore(store EventStore) {
	eventStore = store
}

func PublishCreatedFeed(ctx context.Context, feed *models.Feed) error {
	return eventStore.PublishCreatedFeed(ctx, feed)
}

func SubscribeCreatedFeed(ctx context.Context) (<-chan CreatedFeedMessage, error) {
	return eventStore.SubscribeCreatedFeed(ctx)
}

func OnCreateFeed(ctx context.Context, f func(CreatedFeedMessage)) error {
	return eventStore.OnCreateFeed(f)
}

func Close() {
	eventStore.Close()
}
