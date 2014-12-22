package todo

import (
	"testing"
)

func Test_NewTask(t *testing.T) {
	task, err := NewTask("This is a test")

	if err != nil {
		t.Error(err)
	}

	if task.CreatedAt != task.UpdatedAt {
		t.Errorf("New tasks should have %v equal to %v", task.CreatedAt, task.UpdatedAt)
	}

	if task.IsDone {
		t.Error("New tasks should set IsDone to false")
	}
}

func Test_FailIfTextEmpty(t *testing.T) {
	task, err := NewTask("")
	if err != ErrNilText {
		t.Error("Tasks should fail if no text given")
	}

	if task != nil {
		t.Error("Task should not be created if no text given")
	}
}

func Test_TaskService_Get(t *testing.T) {
	service := TasksService{}
	a, b := service.Get(0)

	if b != nil {
		t.Error("Get should return nil")
	}

	if a == nil {
		t.Error("Get should return a task")
	}
}
