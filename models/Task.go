package models

import "gorm.io/gorm"

type Task struct {
	//ID      uint `gorm:"primary key;autoIncrement" json:"id"`
	gorm.Model
	Content string
	IsRead  bool
}

// Constructor function
func NewTask(content string, isRead bool) *Task {
	return &Task{Content: content, IsRead: isRead}
}
