package dto

import (
	"lab-1/internal/models"
	"time"
)

type CreateTaskDTO struct {
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	Priority    int        `json:"priority"`
	CategoryID  uint       `json:"category_id"`
	Deadline    *time.Time `json:"deadline"`
}

type TaskResponseDTO struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Priority    int32     `json:"priority"`
	CategoryID  uint      `json:"category_id"`
	Deadline    time.Time `json:"deadline"`
	Done        bool      `json:"done"`
}

func NewTaskResponseDTO(task *models.Task) TaskResponseDTO {
	var desc string
	if task.Description != nil {
		desc = *task.Description
	}

	var deadline time.Time
	if task.Deadline != nil {
		deadline = *task.Deadline
	}

	return TaskResponseDTO{
		ID:          task.ID,
		Title:       task.Title,
		Description: desc,
		Priority:    int32(task.Priority),
		CategoryID:  task.CategoryID,
		Deadline:    deadline,
		Done:        task.Done,
	}
}
