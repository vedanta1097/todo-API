package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vedanta1097/todo-list/model"
	"github.com/vedanta1097/todo-list/util"
)

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	todo := model.Todo{}
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		util.Respond(w, util.Message(false, "Invalid request"))
		return
	}
	resp := todo.Create()
	util.Respond(w, resp)
}

func ShowTodoList(w http.ResponseWriter, r *http.Request) {
	response := model.GetTodoList()
	util.Respond(w, response)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	// get id from route variables
	vars := mux.Vars(r)
	id := vars["id"]

	// get todo item according to retrieved id
	todo := model.GetTodo(id)

	// if there is no todo data of ID=id in database, then send error message
	if todo.ID <= 0 {
		util.Respond(w, util.Message(false, "Failed to update data. Invalid id."))
		return
	}

	// update todo according to data from client
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		fmt.Println("error decoding data.")
	}

	response := model.Save(&todo)
	util.Respond(w, response)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	todo := model.GetTodo(id)

	if todo.ID <= 0 {
		util.Respond(w, util.Message(false, "Failed to delete data. Invalid id."))
		return
	}

	response := model.Delete(&todo)
	util.Respond(w, response)
}
