package usecase

import (
	"context"

	"github.com/AlexRijikov/go-activity-service/internal/model"
)

type EventUsecase interface {
	CreateEvent(ctx context.Context, event *model.Event) error
	GetEvents(ctx context.Context, userID int, from, to string) ([]model.Event, error)
	GetStats(ctx context.Context) ([]model.Stat, error)
}
