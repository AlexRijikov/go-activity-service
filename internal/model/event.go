package model

import (
	"encoding/json"
	"time"
)

type Event struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Action string `json:"action"`

	Metadata  json.RawMessage
	CreatedAt time.Time `json:"created_at"`
}
