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
	Title       string
	Description string
	Done        bool
	Priority    Priority
	Deadline    *time.Time `gorm:"default:NULL"`
}
