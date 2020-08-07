package repository

import (
	"github.com/74th/vscode-book-golang/model/tasks"
	"testing"
)

func TestAdd(t *testing.T) {
	rep := New()
	rep.Add(tasks.Task{
		Text: "new task",
	})

	if len(rep.(*instance).tasks) != 3 {
		t.Errorf("タスクが追加されていること %d", len(rep.(*instance).tasks))
	} else {
		addedTask := rep.(*instance).tasks[2]
		if addedTask.Text != "new task" {
			t.Errorf("追加したタスクが末尾に追加されていること %s", addedTask.Text)
		}
		if addedTask.ID <= 2 {
			t.Errorf("タスクに新しいIDがふられること %d", addedTask.ID)
		}
		for i, task := range rep.(*instance).tasks {
			if i != 2 {
				if addedTask.ID == task.ID {
					t.Errorf("既存のタスクとは異なるIDが振られていること %d == %d", addedTask.ID, task.ID)
				}
				if addedTask.Text == task.Text {
					t.Errorf("既存のタスクと上書きしていないこと %s == %s", addedTask.Text, task.Text)
				}
			}
		}
	}
}




