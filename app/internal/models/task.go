package models

import (
	"gorm.io/gorm"
)

const (
	SMALL = iota
	MEDIUM
	HIGH
)

type Task struct {
	gorm.Model
	Title       string
	Description int32
	Done        string
	Priority    int
}
