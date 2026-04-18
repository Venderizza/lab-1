package repo

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"lab-1/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func newTestDB(t *testing.T) *gorm.DB {
	t.Helper()

	dbPath := filepath.Join(t.TempDir(), "test.db")
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	if err := db.AutoMigrate(&models.Task{}); err != nil {
		t.Fatalf("migrate: %v", err)
	}
	return db
}

func TestTaskRepositoryCRUD(t *testing.T) {
	db := newTestDB(t)
	r := NewTaskRepo(db)
	ctx := context.Background()

	deadline := time.Date(2026, 4, 18, 12, 0, 0, 0, time.UTC)
	task := &models.Task{
		Title:       "task",
		Description: "desc",
		Priority:    models.Medium,
		Deadline:    &deadline,
	}

	if err := r.Create(ctx, task); err != nil {
		t.Fatalf("create: %v", err)
	}
	if task.ID == 0 {
		t.Fatal("expected id after create")
	}

	all, err := r.GetAll(ctx)
	if err != nil || len(all) != 1 {
		t.Fatalf("get all: %v, %v", all, err)
	}

	byID, err := r.GetByID(ctx, task.ID)
	if err != nil || byID == nil || byID.Title != "task" {
		t.Fatalf("get by id: %+v, err=%v", byID, err)
	}

	byPriority, err := r.GetByPriority(ctx, int(models.Medium))
	if err != nil || len(byPriority) != 1 {
		t.Fatalf("get by priority: %+v, err=%v", byPriority, err)
	}

	task.Title = "updated"
	task.Description = "updated desc"
	if err := r.Update(ctx, task); err != nil {
		t.Fatalf("update: %v", err)
	}

	updated, err := r.GetByID(ctx, task.ID)
	if err != nil || updated == nil || updated.Title != "updated" {
		t.Fatalf("updated task: %+v, err=%v", updated, err)
	}

	if err := r.ToggleDone(ctx, task.ID); err != nil {
		t.Fatalf("toggle: %v", err)
	}

	toggled, err := r.GetByID(ctx, task.ID)
	if err != nil || toggled == nil || !toggled.Done {
		t.Fatalf("toggled task: %+v, err=%v", toggled, err)
	}

	if err := r.Delete(ctx, task.ID); err != nil {
		t.Fatalf("delete: %v", err)
	}

	deleted, err := r.GetByID(ctx, task.ID)
	if err != nil {
		t.Fatalf("get deleted: %v", err)
	}
	if deleted != nil {
		t.Fatalf("expected nil after delete, got %+v", deleted)
	}
}
