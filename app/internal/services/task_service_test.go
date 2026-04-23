package services

// import (
// 	"context"
// 	"errors"
// 	"testing"
// 	"time"

// 	"lab-1/internal/dto"
// 	"lab-1/internal/models"
// 	"lab-1/internal/repo"

// 	"gorm.io/gorm"
// )

// type mockRepo struct {
// 	createFn        func(context.Context, *models.Task) error
// 	getAllFn        func(context.Context) ([]models.Task, error)
// 	getByIDFn       func(context.Context, uint) (*models.Task, error)
// 	getByPriorityFn func(context.Context, int) ([]models.Task, error)
// 	updateFn        func(context.Context, *models.Task) error
// 	toggleFn        func(context.Context, uint) error
// 	deleteFn        func(context.Context, uint) error
// }

// var _ repo.TaskRepository = (*mockRepo)(nil)

// func (m *mockRepo) Create(ctx context.Context, task *models.Task) error {
// 	if m.createFn != nil {
// 		return m.createFn(ctx, task)
// 	}
// 	return nil
// }
// func (m *mockRepo) GetAll(ctx context.Context) ([]models.Task, error) {
// 	if m.getAllFn != nil {
// 		return m.getAllFn(ctx)
// 	}
// 	return nil, nil
// }
// func (m *mockRepo) GetByID(ctx context.Context, id uint) (*models.Task, error) {
// 	if m.getByIDFn != nil {
// 		return m.getByIDFn(ctx, id)
// 	}
// 	return nil, nil
// }
// func (m *mockRepo) GetByPriority(ctx context.Context, priority int) ([]models.Task, error) {
// 	if m.getByPriorityFn != nil {
// 		return m.getByPriorityFn(ctx, priority)
// 	}
// 	return nil, nil
// }
// func (m *mockRepo) Update(ctx context.Context, task *models.Task) error {
// 	if m.updateFn != nil {
// 		return m.updateFn(ctx, task)
// 	}
// 	return nil
// }
// func (m *mockRepo) ToggleDone(ctx context.Context, id uint) error {
// 	if m.toggleFn != nil {
// 		return m.toggleFn(ctx, id)
// 	}
// 	return nil
// }
// func (m *mockRepo) Delete(ctx context.Context, id uint) error {
// 	if m.deleteFn != nil {
// 		return m.deleteFn(ctx, id)
// 	}
// 	return nil
// }

// func TestCreateTask(t *testing.T) {
// 	deadline := time.Date(2026, 4, 18, 12, 0, 0, 0, time.UTC)

// 	var got *models.Task
// 	svc := NewTaskService(&mockRepo{
// 		createFn: func(_ context.Context, task *models.Task) error {
// 			got = task
// 			return nil
// 		},
// 	})

// 	task, err := svc.CreateTask(context.Background(), &dto.CreateTaskDTO{
// 		Title:       "task",
// 		Description: "desc",
// 		Priority:    99,
// 		Deadline:    deadline,
// 	})
// 	if err != nil {
// 		t.Fatalf("unexpected error: %v", err)
// 	}

// 	if task.Title != "task" || task.Description != "desc" || task.Priority != models.Low || task.Done {
// 		t.Fatalf("unexpected task: %+v", task)
// 	}
// 	if got == nil || got.Deadline == nil || !got.Deadline.Equal(deadline) {
// 		t.Fatalf("deadline was not passed correctly: %+v", got)
// 	}
// }

// func TestCreateTaskValidationError(t *testing.T) {
// 	svc := NewTaskService(&mockRepo{})
// 	_, err := svc.CreateTask(context.Background(), &dto.CreateTaskDTO{})
// 	if err == nil {
// 		t.Fatal("expected error, got nil")
// 	}
// }

// func TestUpdateTask(t *testing.T) {
// 	newTitle := "new title"
// 	newDesc := "new desc"
// 	newPriority := 2
// 	newDeadline := time.Date(2026, 4, 18, 13, 0, 0, 0, time.UTC)

// 	svc := NewTaskService(&mockRepo{
// 		getByIDFn: func(_ context.Context, id uint) (*models.Task, error) {
// 			return &models.Task{
// 				Model: gorm.Model{
// 					ID: id,
// 				},
// 				Title:       "old title",
// 				Description: "old desc",
// 				Priority:    models.Low,
// 			}, nil
// 		},
// 		updateFn: func(_ context.Context, task *models.Task) error {
// 			if task.Title != newTitle || task.Description != newDesc || task.Priority != models.High {
// 				return errors.New("task fields were not updated")
// 			}
// 			if task.Deadline == nil || !task.Deadline.Equal(newDeadline) {
// 				return errors.New("deadline was not updated")
// 			}
// 			return nil
// 		},
// 	})

// 	err := svc.UpdateTask(context.Background(), 1, &dto.UpdateTaskDTO{
// 		Title:       &newTitle,
// 		Description: &newDesc,
// 		Priority:    &newPriority,
// 		Deadline:    &newDeadline,
// 	})
// 	if err != nil {
// 		t.Fatalf("unexpected error: %v", err)
// 	}
// }

// func TestGetTaskByID(t *testing.T) {
// 	deadline := time.Date(2026, 4, 18, 12, 0, 0, 0, time.UTC)

// 	svc := NewTaskService(&mockRepo{
// 		getByIDFn: func(_ context.Context, id uint) (*models.Task, error) {
// 			return &models.Task{
// 				Model: gorm.Model{
// 					ID: id,
// 				},
// 				Title:       "task",
// 				Description: "desc",
// 				Priority:    models.Medium,
// 				Deadline:    &deadline,
// 				Done:        true,
// 			}, nil
// 		},
// 	})

// 	got, err := svc.GetTaskByID(context.Background(), 1)
// 	if err != nil {
// 		t.Fatalf("unexpected error: %v", err)
// 	}
// 	if got.ID != 1 || got.Title != "task" || got.Priority != int32(models.Medium) || !got.Done {
// 		t.Fatalf("unexpected response: %+v", got)
// 	}
// }

// func TestGetTaskByIDNotFound(t *testing.T) {
// 	svc := NewTaskService(&mockRepo{
// 		getByIDFn: func(_ context.Context, _ uint) (*models.Task, error) {
// 			return nil, nil
// 		},
// 	})

// 	got, err := svc.GetTaskByID(context.Background(), 1)
// 	if err != nil {
// 		t.Fatalf("unexpected error: %v", err)
// 	}
// 	if got.ID != 0 {
// 		t.Fatalf("expected empty dto, got %+v", got)
// 	}
// }

// func TestGetAllTasksAndByPriority(t *testing.T) {
// 	svc := NewTaskService(&mockRepo{
// 		getAllFn: func(_ context.Context) ([]models.Task, error) {
// 			return []models.Task{
// 				{Model: gorm.Model{
// 					ID: 1,
// 				}, Title: "a", Priority: models.Low},
// 				{Model: gorm.Model{
// 					ID: 2,
// 				}, Title: "b", Priority: models.High},
// 			}, nil
// 		},
// 		getByPriorityFn: func(_ context.Context, priority int) ([]models.Task, error) {
// 			return []models.Task{
// 				{Model: gorm.Model{
// 					ID: 2,
// 				}, Title: "b", Priority: models.Priority(priority)},
// 			}, nil
// 		},
// 	})

// 	all, err := svc.GetAllTasks(context.Background())
// 	if err != nil || len(all) != 2 {
// 		t.Fatalf("unexpected all tasks result: %+v, err=%v", all, err)
// 	}

// 	byPrio, err := svc.GetTasksByPriority(context.Background(), int(models.High))
// 	if err != nil || len(byPrio) != 1 || byPrio[0].Priority != int32(models.High) {
// 		t.Fatalf("unexpected priority result: %+v, err=%v", byPrio, err)
// 	}
// }
