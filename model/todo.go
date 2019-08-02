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

func GetTodoList() []*Todo {
	// create array of type Todo
	todoList := make([]*Todo, 0)

	// SELECT * FROM Todo
	err := getDB().Find(&todoList).Error

	if err != nil {
		return nil
	}
	return todoList
}
