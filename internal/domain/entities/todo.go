package entities

import "github.com/google/uuid"

type Todo struct {
	Id          string
	Title       string
	Description string
	Done        bool
}

func NewTodo(description, title string) *Todo {
	return &Todo{
		Id:          uuid.New().String(),
		Title:       title,
		Description: description,
		Done:        false,
	}
}

func (t *Todo) ChangeDone() {
	t.Done = !t.Done
}
