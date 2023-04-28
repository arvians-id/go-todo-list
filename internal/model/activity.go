package model

import "time"

type Activity struct {
	ActivityID int       `json:"id"`
	Title      string    `json:"title"`
	Email      string    `json:"email"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
