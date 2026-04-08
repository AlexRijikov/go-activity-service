package scheduler

import (
	"context"
	"log"

	"github.com/AlexRijikov/go-activity-service/internal/repository"
	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	repo repository.EventRepository
}

func NewScheduler(r repository.EventRepository) *Scheduler {
	return &Scheduler{repo: r}
}

func (s *Scheduler) Start() {
	c := cron.New()

	c.AddFunc("@every 4h", func() {
		log.Println("Running aggregation job...")

		err := s.repo.AggregateEvents(context.Background())
		if err != nil {
			log.Println("Aggregation error:", err)
		}
	})

	c.Start()
}
