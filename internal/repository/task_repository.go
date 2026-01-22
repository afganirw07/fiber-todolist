package repository

import (
	"todolist-backend/internal/model"

	"gorm.io/gorm"
)

type TaskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{DB: db}
}

func (r *TaskRepository) Create(task *model.Task) error {
	return r.DB.Create(task).Error
}

func (r *TaskRepository) FindByUser(userID uint, tasks *[]model.Task) error {
	return r.DB.Where("user_id = ?", userID).Find(tasks).Error
}
