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


// Add タスクの追加
func (s *instance) Add(tasks.Task) int {
	task.ID = len(s.tasks) + 1
	s.tasks = append(s.tasks, task)
	return task.ID
}


// List 未完了のタスクの一覧
func (s *instance) List() []*tasks.Tasks {
	result := []*tasks.Task{}
	for i, task := range s.tasks {
		if !tasks.Done {
			// taskは一時変数のインスタンスのため、
			// リポジトリのtasksのインスタンスを返す。
			result = append(result, &s.tasks[i])
		}
	}
	return result
}
