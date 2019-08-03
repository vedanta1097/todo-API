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
	data := model.GetTodoList()
	response := util.Message(true, "Success!")
	response["data"] = data
	util.Respond(w, response)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	// get id from route variables
	vars := mux.Vars(r)
	id := vars["id"]

	// get todo item according to retrieved id
	todo := model.GetTodo(id)

	// get updated data from client
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		fmt.Println("error decoding data.")
	}

	response := model.SaveData(&todo)
	util.Respond(w, response)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {

}
