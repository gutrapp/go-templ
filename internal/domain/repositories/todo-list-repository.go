package repositories

import "templ-go/internal/domain/entities"

type TodoListRepository interface {
	Create(todoList *entities.TodoList) error
	Get(todoListId string) (*entities.TodoList, error)
}
