package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"lab-1/internal/dto"
	"lab-1/internal/models"
	"lab-1/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type handlerMockRepo struct {
	createFn   func(context.Context, *models.Task) error
	getByIDFn  func(context.Context, uint) (*models.Task, error)
	updateFn   func(context.Context, *models.Task) error
	deleteFn   func(context.Context, uint) error
	toggleFn   func(context.Context, uint) error
	getAllFn   func(context.Context) ([]models.Task, error)
	priorityFn func(context.Context, int) ([]models.Task, error)
}

func (m *handlerMockRepo) Create(ctx context.Context, task *models.Task) error {
	if m.createFn != nil {
		return m.createFn(ctx, task)
	}
	return nil
}
func (m *handlerMockRepo) GetAll(ctx context.Context) ([]models.Task, error) {
	if m.getAllFn != nil {
		return m.getAllFn(ctx)
	}
	return nil, nil
}
func (m *handlerMockRepo) GetByID(ctx context.Context, id uint) (*models.Task, error) {
	if m.getByIDFn != nil {
		return m.getByIDFn(ctx, id)
	}
	return nil, nil
}
func (m *handlerMockRepo) GetByPriority(ctx context.Context, priority int) ([]models.Task, error) {
	if m.priorityFn != nil {
		return m.priorityFn(ctx, priority)
	}
	return nil, nil
}
func (m *handlerMockRepo) Update(ctx context.Context, task *models.Task) error {
	if m.updateFn != nil {
		return m.updateFn(ctx, task)
	}
	return nil
}
func (m *handlerMockRepo) ToggleDone(ctx context.Context, id uint) error {
	if m.toggleFn != nil {
		return m.toggleFn(ctx, id)
	}
	return nil
}
func (m *handlerMockRepo) Delete(ctx context.Context, id uint) error {
	if m.deleteFn != nil {
		return m.deleteFn(ctx, id)
	}
	return nil
}

func TestCreateTaskHandler(t *testing.T) {
	h := NewTaskHandler(services.NewTaskService(&handlerMockRepo{
		createFn: func(_ context.Context, task *models.Task) error {
			if task.Title != "task" {
				t.Fatalf("unexpected title: %q", task.Title)
			}
			return nil
		},
	}))

	app := fiber.New()
	app.Post("/api/tasks", h.CreateTask)

	req := httptest.NewRequest(http.MethodPost, "/api/tasks", strings.NewReader(`{"title":"task","priority":1}`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("app test: %v", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected 201, got %d", resp.StatusCode)
	}

	var got dto.TaskResponseDTO
	if err := json.NewDecoder(resp.Body).Decode(&got); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if got.Title != "task" {
		t.Fatalf("unexpected response: %+v", got)
	}
}

func TestCreateTaskHandlerValidationError(t *testing.T) {
	h := NewTaskHandler(services.NewTaskService(&handlerMockRepo{}))

	app := fiber.New()
	app.Post("/api/tasks", h.CreateTask)

	req := httptest.NewRequest(http.MethodPost, "/api/tasks", strings.NewReader(`{"title":""}`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("app test: %v", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", resp.StatusCode)
	}
}

func TestGetTaskByIDNotFound(t *testing.T) {
	h := NewTaskHandler(services.NewTaskService(&handlerMockRepo{
		getByIDFn: func(_ context.Context, _ uint) (*models.Task, error) {
			return nil, nil
		},
	}))

	app := fiber.New()
	app.Get("/api/tasks/:id", h.GetTaskByID)

	req := httptest.NewRequest(http.MethodGet, "/api/tasks/10", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("app test: %v", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", resp.StatusCode)
	}
}

func TestUpdateTaskHandler(t *testing.T) {
	h := NewTaskHandler(services.NewTaskService(&handlerMockRepo{
		getByIDFn: func(_ context.Context, id uint) (*models.Task, error) {
			return &models.Task{Model: gorm.Model{
				ID: id,
			}, Title: "old"}, nil
		},
		updateFn: func(_ context.Context, task *models.Task) error {
			if task.Title != "new" {
				t.Fatalf("unexpected title: %q", task.Title)
			}
			return nil
		},
	}))

	app := fiber.New()
	app.Put("/api/tasks/:id", h.UpdateTask)

	req := httptest.NewRequest(http.MethodPut, "/api/tasks/1", strings.NewReader(`{"title":"new"}`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("app test: %v", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusNoContent {
		t.Fatalf("expected 204, got %d", resp.StatusCode)
	}
}
