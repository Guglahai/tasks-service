package task

import "time"

type Task struct {
	ID        uint `gorm:"primaryKey"`
	Task      string
	Is_done   bool
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
