package database

import (
	"errors"
	"templ-go/internal/domain/entities"
)

type TodoListRepositoryMemory struct {
	TodoLists []entities.TodoList
}

func NewTodoListRepositoryMemory() *TodoListRepositoryMemory {
	return &TodoListRepositoryMemory{
		TodoLists: []entities.TodoList{},
	}
}

func (tr *TodoListRepositoryMemory) Create(todo *entities.TodoList) error {
	tr.TodoLists = append(tr.TodoLists, *todo)
	return nil
}

func (tr *TodoListRepositoryMemory) Get(todoListId string) (*entities.TodoList, error) {
	for _, todo := range tr.TodoLists {
		if todo.Id == todoListId {
			return &todo, nil
		}
	}
	return nil, errors.New("ERROR: Couldn't find todo")
}
