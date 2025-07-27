package task

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	CreateTask(task *Task) (*Task, error)
	GetAllTasks() ([]*Task, error)
	GetTaskByID(id uint) (*Task, error)
	GetTasksByUserID(userID uint) ([]*Task, error)
	UpdateTask(task *Task) (Task, error)
	DeleteTask(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateTask(task *Task) (*Task, error) {
	if err := r.db.Create(&task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (r *repository) GetAllTasks() ([]*Task, error) {
	var tasks []*Task
	if err := r.db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *repository) GetTaskByID(id uint) (*Task, error) {
	var task *Task
	if err := r.db.First(&task, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (r *repository) GetTasksByUserID(userID uint) ([]*Task, error) {
	var tasks []*Task
	if err := r.db.Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *repository) UpdateTask(task *Task) (Task, error) {
	err := r.db.
		Model(&Task{}).
		Where("id = ?", task.ID).
		Updates(map[string]interface{}{
			"task":       task.Task,
			"is_done":    task.Is_done,
			"user_id":    task.UserID,
			"updated_at": time.Now(),
		}).Error
	if err != nil {
		log.Printf("Error updating task: %v", err)
		return Task{}, err
	}
	return *task, nil
}

func (r *repository) DeleteTask(id uint) error {
	return r.db.Delete(&Task{}, "id = ?", id).Error
}
