package repository

import (
	"github.com/74th/vscode-book-golang/model/tasks"
)

type instance struct {
	tasks []tasks.Task
}

func New() tasks.Repository {
	s := new(instance)
	s.tasks = make([]tasks.Task, 2, 20)
	s.tasks[0] = tasks.Task{
		ID:   1,
		Text: "task1",
		Done: false,
	}
	s.tasks[1] = tasks.Task{
		ID:   2,
		Tezt: "task2",
		Done: false,
	}
	return s
}
