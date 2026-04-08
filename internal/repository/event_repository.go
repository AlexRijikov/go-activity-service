package repository

import (
	"context"
	"database/sql"

	"github.com/AlexRijikov/go-activity-service/internal/model"
)

type eventRepository struct {
	db *sql.DB
}

func NewEventRepository(db *sql.DB) EventRepository {
	return &eventRepository{db: db}
}

func (r *eventRepository) CreateEvent(ctx context.Context, e *model.Event) error {
	query := `
        INSERT INTO events (user_id, action, metadata)
        VALUES ($1, $2, $3)
    `
	_, err := r.db.ExecContext(ctx, query, e.UserID, e.Action, e.Metadata)
	return err
}

func (r *eventRepository) GetEvents(ctx context.Context, userID int, from, to string) ([]model.Event, error) {
	query := `
        SELECT id, user_id, action, metadata, created_at
        FROM events
        WHERE user_id = $1 AND created_at BETWEEN $2 AND $3
        ORDER BY created_at DESC
    `

	rows, err := r.db.QueryContext(ctx, query, userID, from, to)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []model.Event

	for rows.Next() {
		var e model.Event
		if err := rows.Scan(&e.ID, &e.UserID, &e.Action, &e.Metadata, &e.CreatedAt); err != nil {
			return nil, err
		}
		events = append(events, e)
	}

	return events, nil
}

func (r *eventRepository) AggregateEvents(ctx context.Context) error {
	query := `
		INSERT INTO aggregated_stats (user_id, event_count, period_start, period_end)
		SELECT 
			user_id,
			COUNT(*),
			NOW() - INTERVAL '4 hours',
			NOW()
		FROM events
		WHERE created_at >= NOW() - INTERVAL '4 hours'
		GROUP BY user_id;
	`

	_, err := r.db.ExecContext(ctx, query)
	return err
}

func (r *eventRepository) GetStats(ctx context.Context) ([]model.Stat, error) {
	query := `
		SELECT user_id, event_count, period_start, period_end
		FROM aggregated_stats
		ORDER BY period_start DESC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stats []model.Stat

	for rows.Next() {
		var s model.Stat

		if err := rows.Scan(&s.UserID, &s.Count, &s.From, &s.To); err != nil {
			return nil, err
		}

		stats = append(stats, s)
	}

	return stats, nil
}
