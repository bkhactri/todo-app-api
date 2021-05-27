package main

import (
	"log"
	"net/http"

	"github.com/tobias/todo-app/middlewares"
	"github.com/tobias/todo-app/routes"

	"github.com/gorilla/mux"
)

func createServer() {
	port := ":9000"
	router := mux.NewRouter()
	//User Service API
	router.HandleFunc("/auth/login", routes.SignIn).Methods("POST")
	router.HandleFunc("/auth/register", routes.SignUp).Methods("POST")

	// Task Service API
	router.HandleFunc("/tasks", middlewares.CheckAuthentication(routes.GetTasks)).Methods("GET")
	router.HandleFunc("/tasks/{id}", middlewares.CheckAuthentication(routes.GetTask)).Methods("GET")
	router.HandleFunc("/tasks", middlewares.CheckAuthentication(routes.CreateTask)).Methods("POST")
	router.HandleFunc("/tasks/{id}", middlewares.CheckAuthentication(routes.UpdateTask)).Methods("PUT")
	router.HandleFunc("/tasks/{id}", middlewares.CheckAuthentication(routes.DeleteTask)).Methods("DELETE")

	log.Println("Server is listening on port", port)
	log.Fatal(http.ListenAndServe(port, router))

}

func main() {
	createServer()
}
