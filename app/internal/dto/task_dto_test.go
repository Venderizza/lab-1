package dto

import (
	"testing"
	"time"

	"lab-1/internal/models"

	"gorm.io/gorm"
)

func TestCreateTaskDTOValidate(t *testing.T) {
	t.Run("title required", func(t *testing.T) {
		dto := CreateTaskDTO{}
		if err := dto.Validate(); err == nil {
			t.Fatal("expected validation error, got nil")
		}
	})

	t.Run("valid dto", func(t *testing.T) {
		dto := CreateTaskDTO{Title: "task"}
		if err := dto.Validate(); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestNewTaskResponseDTO(t *testing.T) {
	deadline := time.Date(2026, 4, 18, 12, 0, 0, 0, time.UTC)

	got := NewTaskResponseDTO(models.Task{
		Model: gorm.Model{
			ID: 1,
		},
		Title:       "demo",
		Description: "",
		Done:        true,
		Priority:    models.High,
		Deadline:    &deadline,
	})

	if got.ID != 1 || got.Title != "demo" || got.Priority != int32(models.High) || !got.Done {
		t.Fatalf("unexpected dto: %+v", got)
	}
	if got.Description != "" {
		t.Fatalf("unexpected description: %q", got.Description)
	}
	if got.Deadline == nil || !got.Deadline.Equal(deadline) {
		t.Fatalf("unexpected deadline: %+v", got.Deadline)
	}
}
