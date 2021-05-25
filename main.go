package main

import (
	"log"
	"net/http"

	"github.com/tobias/todo-app/routes"

	"github.com/gorilla/mux"
)

func createServer() {
	port := ":9000"
	router := mux.NewRouter()

	// API
	router.HandleFunc("/tasks", routes.GetTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", routes.GetTask).Methods("GET")
	router.HandleFunc("/tasks", routes.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", routes.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", routes.DeleteTask).Methods("DELETE")

	log.Println("Server is listening on port", port)
	log.Fatal(http.ListenAndServe(port, router))

}

func main() {
	createServer()
}
