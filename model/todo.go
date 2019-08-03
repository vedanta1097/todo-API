package model

import (
	"github.com/jinzhu/gorm"
	"github.com/vedanta1097/todo-list/util"
)

type Todo struct {
	gorm.Model
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Cost   uint   `json:"cost"`
}

func (todo *Todo) Create() map[string]interface{} {
	// input todo data to database
	getDB().Create(todo)

	if todo.ID <= 0 {
		return util.Message(false, "Failed to create todo, connection error.")
	}

	response := util.Message(true, "Todo has been created.")
	response["data"] = todo
	return response
}

func GetTodoList() map[string]interface{} {
	// create slice of type Todo
	todoList := make([]Todo, 0)

	// SELECT * FROM Todo
	err := getDB().Find(&todoList).Error

	if err != nil {
		return util.Message(false, "Failed to get data, connection error.")
	}

	response := util.Message(true, "Success.")
	response["data"] = todoList

	return response
}

func GetTodo(id string) Todo {
	todo := Todo{}
	// select todo where ID = id
	getDB().First(&todo, id)
	return todo
}

func Save(todo *Todo) map[string]interface{} {
	// save todo data
	err := getDB().Save(todo).Error

	if err != nil {
		return util.Message(false, "failed to save to DB")
	}
	return util.Message(true, "Data successfully updated.")
}

func Delete(todo *Todo) map[string]interface{} {
	// delete todo data
	err := getDB().Delete(todo).Error

	if err != nil {
		return util.Message(false, "failed to delete data.")
	}
	return util.Message(true, "Data successfully deleted.")
}
