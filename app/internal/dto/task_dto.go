package dto

import (
	"errors"
	"lab-1/internal/models"
	"time"
)

type CreateTaskDTO struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Priority    int       `json:"priority"`
	Deadline    time.Time `json:"deadline"`
}

type UpdateTaskDTO struct {
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	Priority    *int       `json:"priority"`
	Deadline    *time.Time `json:"deadline"`
}

type TaskResponseDTO struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Priority    int32      `json:"priority"`
	Deadline    *time.Time `json:"deadline"`
	Done        bool       `json:"done"`
}

func NewTaskResponseDTO(task models.Task) TaskResponseDTO {
	desc := ""
	if task.Description != "" {
		desc = task.Description
	}

	return TaskResponseDTO{
		ID:          task.ID,
		Title:       task.Title,
		Description: desc,
		Priority:    int32(task.Priority),
		Deadline:    task.Deadline,
		Done:        task.Done,
	}
}

func (dto *CreateTaskDTO) Validate() error {
	if dto.Title == "" {
		return errors.New("title is required")
	}
	return nil
}
