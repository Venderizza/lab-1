package models

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  int32
	Hash string
}

type UserRepo interface {
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id int64) (*User, error)
	FilterByAge(ctx context.Context, id int64) (*[]User, error)
}
