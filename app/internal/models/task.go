package models

import (
	"time"

	"gorm.io/gorm"
)

type Priority int

const (
	Low Priority = iota
	Medium
	High
)

type Task struct {
	gorm.Model
	CategoryID uint
	Category   Category `gorm:"foreignKey:CategoryID"`

	CompletedAt *time.Time
	Title       string
	Description *string
	Done        bool
	Priority    Priority
	Deadline    *time.Time
}
