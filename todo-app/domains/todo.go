package domains

import "time"

type Todo struct {
	Id          int64
	Title       string
	Description string
	IsDone      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DueDate     *time.Time
}
