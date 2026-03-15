package repo

import (
	"context"
	"errors"
	"lab-1/internal/models"

	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(ctx context.Context, task *models.Task) error
	GetAll(ctx context.Context) ([]models.Task, error)
	GetByID(ctx context.Context, id uint) (*models.Task, error)
	GetByPriority(ctx context.Context, id int) ([]models.Task, error)
	ToggleDone(ctx context.Context, id uint) error
	Update(ctx context.Context, task *models.Task) error
	Delete(ctx context.Context, id uint) error
}

type TaskIMP struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) TaskRepository {
	return &TaskIMP{db: db}
}

func (r *TaskIMP) Create(ctx context.Context, task *models.Task) error {
	return r.db.WithContext(ctx).Create(task).Error
}

func (r *TaskIMP) GetAll(ctx context.Context) ([]models.Task, error) {
	var tasks []models.Task

	if err := r.db.WithContext(ctx).Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *TaskIMP) GetByID(ctx context.Context, id uint) (*models.Task, error) {
	var task models.Task
	err := r.db.WithContext(ctx).First(&task, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &task, nil
}

func (r *TaskIMP) GetByPriority(ctx context.Context, priority int) ([]models.Task, error) {
	var tasks []models.Task
	if err := r.db.WithContext(ctx).
		Where("priority = ?", priority).
		Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskIMP) Update(ctx context.Context, task *models.Task) error {
	return r.db.WithContext(ctx).Model(&models.Task{}).Where("id = ?", task.ID).Updates(task).Error
}

func (r *TaskIMP) ToggleDone(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).
		Model(&models.Task{}).
		Where("id = ?", id).
		Update("done", gorm.Expr("NOT done")).Error
}

func (r *TaskIMP) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Task{}, id).Error
}
