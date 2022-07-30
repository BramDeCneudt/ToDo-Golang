package db

import "example.com/todo-server/types"

type DbServiceInterface interface {
	Save(todo types.Todo) types.Todo
	Delete(id int)
	GetAll() []types.Todo
	Get(id int) types.Todo
}