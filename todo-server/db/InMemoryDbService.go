package db

import (
	"example.com/todo-server/types"
)

type InMemoryDbService struct {
	todos map[int]types.Todo
}

func NewInMemoryDbService() *InMemoryDbService {
	var service InMemoryDbService
	service.todos = make(map[int]types.Todo)
	return &service
}

func (service InMemoryDbService) Save(todo types.Todo) types.Todo {
	todo = createTodo(todo.Description)
	service.todos[todo.Id] = todo
	return todo
}

func (service InMemoryDbService) Delete(id int) {
	delete(service.todos, id)
}

func (service InMemoryDbService) GetAll() []types.Todo {
	return service.getTodos()
}

func (service InMemoryDbService) Get(id int) types.Todo {
	return service.todos[id]
}

func (service InMemoryDbService) getTodos() []types.Todo {
	var todoSlice = []types.Todo{}
	for _, element := range service.todos {
		todoSlice = append(todoSlice, element)
	}
	return todoSlice
}

var todoId = 1
func createTodo(description string) types.Todo {
	var todo types.Todo
	todo.Id = todoId
	todo.Description = description

	todoId++
	return todo
}