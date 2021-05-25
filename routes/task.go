package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	db "github.com/tobias/todo-app/database"
	"github.com/tobias/todo-app/models"
	res "github.com/tobias/todo-app/utils"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	collection := db.ConnectTask()
	w.Header().Set("Content-Type", "application/json")
	var tasks []models.Task
	if err := collection.Find(&tasks); err != nil {
		res.JSON(w, 500, err)
	} else {
		res.JSON(w, 200, tasks)
	}
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	collection := db.ConnectTask()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var task models.Task
	if err := collection.First(&task, params["id"]); err != nil {
		res.JSON(w, 500, "Task not found")
	} else {
		res.JSON(w, 200, task)
	}
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	collection := db.ConnectTask()
	w.Header().Set("Content-Type", "application/json")
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	if err := collection.Create(&task); err != nil {
		res.JSON(w, 500, "Can not create Task")
	} else {
		res.JSON(w, 200, task)
	}
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	collection := db.ConnectTask()
	w.Header().Set("Content-Type", "application/")
	params := mux.Vars(r)
	var task models.Task
	collection.First(&task, params["id"])
	json.NewDecoder(r.Body).Decode(&task)
	if err := collection.Save(&task); err != nil {
		res.JSON(w, 500, "Can not update Task")
	} else {
		res.JSON(w, 200, task)
	}
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	collection := db.ConnectTask()
	w.Header().Set("Content-Type", "application/")
	params := mux.Vars(r)
	var task models.Task
	// permanently delete
	if err := collection.Unscoped().Delete(&task, params["id"]); err != nil {
		res.JSON(w, 500, "Can not delete Task")
	} else {
		res.JSON(w, 200, task)
	}
}
