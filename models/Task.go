package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Content string
	IsRead  bool
}

// Constructor function
func NewTask(content string, isRead bool) *Task {
	return &Task{Content: content, IsRead: isRead}
}
