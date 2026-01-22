package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title     string `gorm:"not null" json:"title"`
	Completed bool   `gorm:"default:false" json:"completed"`
	UserID    uint   `gorm:"not null" json:"user_id"`
}
