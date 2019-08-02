package controller

import (
	"encoding/json"
	"net/http"

	"github.com/vedanta1097/todo-list/model"
	"github.com/vedanta1097/todo-list/util"
)

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	todo := &model.Todo{}
	err := json.NewDecoder(r.Body).Decode(todo)
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

func DeleteTodo(w http.ResponseWriter, r *http.Request) {

}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {

}
