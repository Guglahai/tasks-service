package task

import "time"

type Task struct {
	ID        int `gorm:"primaryKey"`
	Task      string
	Is_done   bool
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
