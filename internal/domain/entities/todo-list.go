package entities

import "github.com/google/uuid"

type TodoList struct {
	Id    string
	Name  string
	Todos []Todo
}

func NewTodoList() *TodoList {
	return &TodoList{
		Id:    uuid.New().String(),
		Todos: []Todo{},
	}
}

func (td *TodoList) AddTodo(todo Todo) {
	td.Todos = append(td.Todos, todo)
}

func (td *TodoList) RemoveTodo(index uint) {
	td.Todos = append(td.Todos[:index], td.Todos[index+1:]...)
}
