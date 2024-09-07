package tasks

import (
	"github.com/dev-oleksandrv/db"
	"github.com/dev-oleksandrv/repository"
)

type TaskService struct {
	taskRepo repository.GORMTaskRepository
}

func NewTaskService(taskRepo repository.GORMTaskRepository) *TaskService {
	return &TaskService{ taskRepo }
}

func (s *TaskService) CreateTask(task db.Task) (db.Task, error) {
	return s.taskRepo.Create(task)
}

func (s *TaskService) GetAllTasks() ([]db.Task, error) {
	return s.taskRepo.GetAll()
}

func (s *TaskService) DeleteTask(id uint) error {
	return s.taskRepo.Delete(id)
}