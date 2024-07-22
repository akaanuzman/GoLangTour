package model

import "time"

type TodoCreate struct {
	Title       string
	Description string
	IsDone      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DueDate     *time.Time
}
