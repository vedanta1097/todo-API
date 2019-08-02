package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/vedanta1097/todo-list/controller"
)

func main() {
	r := mux.NewRouter()
	todo := r.PathPrefix("/api/v1/todo").Subrouter()
	todo.HandleFunc("/create", controller.CreateTodo).Methods("POST")
	todo.HandleFunc("/show", controller.ShowTodoList).Methods("GET")
	todo.HandleFunc("/{id}/update", controller.UpdateTodo).Methods("POST")
	todo.HandleFunc("/{id}/delete", controller.DeleteTodo).Methods("POST")
	http.Handle("/", r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	fmt.Println("API run in port " + port)

	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		fmt.Println(err)
	}

}
