package repo

import (
	"context"
	"errors"
	"lab-1/internal/models"

	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(task *models.Task) error
	GetByID(id int64) (*models.Task, error)
	GetAll() ([]models.Task, error)
	Update(id int64) error
	Delete(id int64) error
}

type TaskIMP struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) *TaskIMP {
	return &TaskIMP{db: db}
}

func (r *TaskIMP) Create(ctx context.Context, task *models.Task) error {
	return r.db.WithContext(ctx).Create(task).Error
}

func (r *TaskIMP) GetByID(ctx context.Context, id int64) (*models.Task, error) {
	var user models.Task
	err := r.db.WithContext(ctx).First(&user, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *TaskIMP) GetAll() ([]models.Task, error) {
	var tasks []models.Task

	err := r.db.WithContext(nil).Find(&tasks).Error
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
