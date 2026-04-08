package usecase

import (
	"context"

	"github.com/AlexRijikov/go-activity-service/internal/model"
	"github.com/AlexRijikov/go-activity-service/internal/repository"
)

// Додати интерфейс
type EventUsecase struct {
	repo repository.EventRepository
}

func NewEventUsecase(r repository.EventRepository) *EventUsecase {
	return &EventUsecase{repo: r}
}

func (u *EventUsecase) CreateEvent(ctx context.Context, event *model.Event) error {
	return u.repo.CreateEvent(ctx, event)
}

func (u *EventUsecase) GetEvents(ctx context.Context, userID int, from, to string) ([]model.Event, error) {
	return u.repo.GetEvents(ctx, userID, from, to)
}

func (u *EventUsecase) GetStats(ctx context.Context) ([]map[string]interface{}, error) {
	return u.repo.GetStats(ctx)
}
