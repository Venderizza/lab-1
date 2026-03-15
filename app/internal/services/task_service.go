package services

import (
	"context"
	"errors"
	"lab-1/internal/dto"
	"lab-1/internal/models"
	"lab-1/internal/repo"
	"time"
)

type TaskService struct {
	taskRepo repo.TaskRepository
}

func NewTaskService(taskRepo repo.TaskRepository) *TaskService {
	return &TaskService{taskRepo: taskRepo}
}

func (s *TaskService) CreateTask(ctx context.Context, dto *dto.CreateTaskDTO) (*models.Task, error) {
	if dto.Title == "" {
		return nil, errors.New("title is required")
	}

	priority := dto.Priority
	if !IsValidPriority(priority) {
		priority = int(models.Low)
	}

	var deadline *time.Time
	if !dto.Deadline.IsZero() {
		deadline = &dto.Deadline
	}

	task := &models.Task{
		Title:       dto.Title,
		Description: dto.Description,
		Done:        false,
		Priority:    models.Priority(priority),
		Deadline:    deadline,
	}

	if err := s.taskRepo.Create(ctx, task); err != nil {
		return nil, err
	}
	return task, nil
}

func (s *TaskService) UpdateTask(ctx context.Context, id uint, dto *dto.UpdateTaskDTO) error {
	task, err := s.taskRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if task == nil {
		return errors.New("task not found")
	}

	if dto.Title != nil {
		task.Title = *dto.Title
	}
	if dto.Description != nil {
		task.Description = *dto.Description
	}
	if dto.Priority != nil && IsValidPriority(*dto.Priority) {
		task.Priority = models.Priority(*dto.Priority)
	}
	if dto.Deadline != nil {
		task.Deadline = dto.Deadline
	}

	return s.taskRepo.Update(ctx, task)
}

func (s *TaskService) GetTaskByID(ctx context.Context, id uint) (dto.TaskResponseDTO, error) {
	task, err := s.taskRepo.GetByID(ctx, id)
	if err != nil {
		return dto.TaskResponseDTO{}, err
	}
	if task == nil {
		return dto.TaskResponseDTO{}, nil
	}
	return dto.NewTaskResponseDTO(*task), nil
}

func (s *TaskService) GetAllTasks(ctx context.Context) ([]dto.TaskResponseDTO, error) {
	tasks, err := s.taskRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return createTaskResponseSlice(tasks), nil
}

func (s *TaskService) GetTasksByPriority(ctx context.Context, priority int) ([]dto.TaskResponseDTO, error) {
	tasks, err := s.taskRepo.GetByPriority(ctx, priority)
	if err != nil {
		return nil, err
	}
	return createTaskResponseSlice(tasks), nil
}

func (s *TaskService) ToggleTask(ctx context.Context, id uint) error {
	return s.taskRepo.ToggleDone(ctx, id)
}

func (s *TaskService) DeleteTask(ctx context.Context, id uint) error {
	return s.taskRepo.Delete(ctx, id)
}

func createTaskResponseSlice(tasks []models.Task) []dto.TaskResponseDTO {
	var resp []dto.TaskResponseDTO
	for _, t := range tasks {
		resp = append(resp, dto.NewTaskResponseDTO(t))
	}
	return resp
}

func IsValidPriority(p int) bool {
	return p >= int(models.Low) && p <= int(models.High)
}
