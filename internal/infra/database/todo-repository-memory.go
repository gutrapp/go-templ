package database

import (
	"errors"
	"templ-go/internal/domain/entities"
)

type TodoRepositoryMemory struct {
	Todos []entities.Todo
}

func NewTodoRepositoryMemory() *TodoRepositoryMemory {
	return &TodoRepositoryMemory{
		Todos: []entities.Todo{},
	}
}

func (tr *TodoRepositoryMemory) Create(todo *entities.Todo) error {
	tr.Todos = append(tr.Todos, *todo)
	return nil
}

func (tr *TodoRepositoryMemory) Get(todoId string) (*entities.Todo, error) {
	for _, todo := range tr.Todos {
		if todo.Id == todoId {
			return &todo, nil
		}
	}
	return nil, errors.New("ERROR: Couldn't find todo")
}
