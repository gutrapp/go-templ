package repositories

import "templ-go/internal/domain/entities"

type TodoRepository interface {
	Create(todo *entities.Todo) error
	Get(todoId string) (*entities.Todo, error)
}
