package model

import "time"

type Stat struct {
	UserID int       `json:"user_id"`
	Count  int       `json:"count"`
	From   time.Time `json:"from"`
	To     time.Time `json:"to"`
}
