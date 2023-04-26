package model

import "time"

type Todo struct {
	TodoID          int       `json:"id"`
	ActivityGroupID int       `json:"activity_group_id"`
	Title           string    `json:"title"`
	IsActive        bool      `json:"is_active"`
	Priority        string    `json:"priority"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}
