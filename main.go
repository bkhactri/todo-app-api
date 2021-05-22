package main

import (
	"log"
	"net/http"

	Handler "todo-app/core"

	"github.com/gorilla/mux"
)

func createServer() {
	port := ":9000"
	router := mux.NewRouter()

	// API
	router.HandleFunc("/tasks", Handler.GetTasks).Methods("GET")
	// router.HandleFunc("/tasks/{id}", Handler.GetTask).Methods("GET")
	router.HandleFunc("/tasks", Handler.CreateTask).Methods("POST")
	// router.HandleFunc("/tasks/{id}", Handler.UpdateTask).Methods("PUT")
	// router.HandleFunc("/tasks/{id}", Handler.DeleteTask).Methods("DELETE")

	log.Println("Server is listening on port", port)
	log.Fatal(http.ListenAndServe(port, router))

}

func main() {
	Handler.ConnectCockRoachDB()
	createServer()
}
