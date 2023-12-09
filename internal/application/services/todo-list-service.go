package services

import (
	"templ-go/internal/domain/entities"
	"templ-go/internal/domain/repositories"
	"templ-go/internal/infra/database"
)

type TodoListService struct {
	TodoListRepository repositories.TodoListRepository
}

func NewTodoListService() *TodoListService {
	return &TodoListService{
		TodoListRepository: database.NewTodoListRepositoryMemory(),
	}
}

func (ts *TodoListService) GetTotal(todoListId string) (int, error) {
	todoList, err := ts.TodoListRepository.Get(todoListId)
	if err != nil {
		return 0, nil
	}

	return len(todoList.Todos), nil
}

func (ts *TodoListService) GetPorcentage(todoListId string) (int, error) {
	todoList, err := ts.TodoListRepository.Get(todoListId)
	if err != nil {
		return 0, nil
	}

	var completeTodos int
	for _, todo := range todoList.Todos {
		if todo.Done {
			completeTodos++
		}
	}

	return (completeTodos / len(todoList.Todos)) * 100, nil
}

func (ts *TodoListService) CreateTodoList(todoList *entities.TodoList) error {
	err := ts.TodoListRepository.Create(todoList)
	if err != nil {
		return nil
	}
	return nil
}
