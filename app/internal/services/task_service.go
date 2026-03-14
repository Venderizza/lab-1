package services

import (
	"lab-1/internal/dto"
	"lab-1/internal/models"
	"lab-1/internal/repo"
)

type TaskService struct {
	taskRepo repo.TaskRepository
}

func NewTaskService(taskRepo repo.TaskRepository) *TaskService {
	return &TaskService{taskRepo: taskRepo}
}

func (s *TaskService) CreateTask(dto *dto.CreateTaskDTO) (*models.Task, error) {
	task := &models.Task{
		Title:       dto.Title,
		Description: dto.Description,
		Done:        false,
		Priority:    models.Priority(dto.Priority),
		Deadline:    dto.Deadline,
	}

	if err := s.taskRepo.Create(task); err != nil {
		return nil, err
	}
	return task, nil
}
