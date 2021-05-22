package core

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
}
