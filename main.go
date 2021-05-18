package main

import (
	"log"
	"net/http"

	Handler "todo-app/service/api"
	DB "todo-app/util"

	"github.com/gorilla/mux"
)

func createServer() {
	port := ":9000"
	router := mux.NewRouter()

	// API
	router.HandleFunc("/tasks", Handler.GetTasks).Methods("GET")
	// router.HandleFunc("/tasks/{id}", taskHandler.GetTask).Methods("GET")
	// router.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
	// router.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("PUT")
	// router.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")

	log.Println("Server is listening on port", port)
	log.Fatal(http.ListenAndServe(port, router))

}

func main() {
	createServer()
	DB.ConnectCockRoachDB()
}
