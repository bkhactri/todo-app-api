package api

import (
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	task := &Task{Title: "Task 1", Description: "Make demo todo app", Author: "Tobias"}
	json.NewEncoder(w).Encode(task)
}

// func GetTask(w http.ResponseWriter, r *http.Request) {

// }

// func CreateTask(w http.ResponseWriter, r *http.Request) {

// }

// func UpdateTask(w http.ResponseWriter, r *http.Request) {

// }

// func DeleteTask(w http.ResponseWriter, r *http.Request) {

// }
