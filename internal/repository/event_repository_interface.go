package repository

import (
	"context"

	"github.com/AlexRijikov/go-activity-service/internal/model"
)

type EventRepository interface {
	CreateEvent(ctx context.Context, e *model.Event) error
	GetEvents(ctx context.Context, userID int, from, to string) ([]model.Event, error)
	AggregateEvents(ctx context.Context) error
	GetStats(ctx context.Context) ([]model.Stat, error)
}
