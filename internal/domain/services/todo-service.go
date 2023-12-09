package services

import "templ-go/internal/domain/entities"

type TodoService interface {
	CreateTodo(todo *entities.Todo) (int, error)
}
