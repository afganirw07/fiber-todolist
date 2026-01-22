package service

import (
	"todolist-backend/internal/model"
	"todolist-backend/internal/repository"
)

type TaskService struct {
	Repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{Repo: repo}
}

func (s *TaskService) Create(task *model.Task) error {
	return s.Repo.Create(task)
}

func (s *TaskService) GetByUser(userID uint, tasks *[]model.Task) error {
	return s.Repo.FindByUser(userID, tasks)
}
