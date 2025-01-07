package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Content     string
	IsCompleted bool
}

// Constructor function
func NewTask(content string, IsCompleted bool) *Task {
	return &Task{Content: content, IsCompleted: IsCompleted}
}
