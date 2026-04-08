package usecase

import (
	"context"

	"github.com/AlexRijikov/go-activity-service/internal/model"
	"github.com/AlexRijikov/go-activity-service/internal/repository"
)

type eventUsecase struct {
	repo repository.EventRepository
}

func NewEventUsecase(r repository.EventRepository) EventUsecase {
	return &eventUsecase{repo: r}
}

func (u *eventUsecase) CreateEvent(ctx context.Context, event *model.Event) error {
	return u.repo.CreateEvent(ctx, event)
}

func (u *eventUsecase) GetEvents(ctx context.Context, userID int, from, to string) ([]model.Event, error) {
	return u.repo.GetEvents(ctx, userID, from, to)
}

func (u *eventUsecase) GetStats(ctx context.Context) ([]model.Stat, error) {
	return u.repo.GetStats(ctx)
}
