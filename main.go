package main

import (
	"log"
	"net/http"

	Handler "todo-app/service/api"
	DB "todo-app/service/util"

	"github.com/gorilla/mux"
)

func createServer() {
	router := mux.NewRouter()

	// API
	router.HandleFunc("/tasks", Handler.GetTasks).Methods("GET")
	// router.HandleFunc("/tasks/{id}", taskHandler.GetTask).Methods("GET")
	// router.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
	// router.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("PUT")
	// router.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", router))

}

func main() {
	createServer()
	DB.ConnectCockRoachDB()
}
