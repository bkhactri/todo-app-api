package core

import (
	"encoding/json"
	"net/http"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	task := &Task{Title: "Task 1", Description: "Make demo todo app", Author: "Tobias"}
	json.NewEncoder(w).Encode(task)
}

// func GetTask(w http.ResponseWriter, r *http.Request) {

// }

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var task Task
	json.NewDecoder(r.Body).Decode(&task)
	db.Create(&task)
	json.NewEncoder(w).Encode(task)
}

// func UpdateTask(w http.ResponseWriter, r *http.Request) {

// }

// func DeleteTask(w http.ResponseWriter, r *http.Request) {

// }
