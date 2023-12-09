package services

import "templ-go/internal/domain/entities"

type TodoListService interface {
	CreateTodoList(todoList *entities.TodoList) error
	GetTotal(todoListId string) (int, error)
	GetPorcentage(todoListId string) (int, error)
}
