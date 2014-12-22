package todo

import (
	"errors"
	"time"
)

type Task struct {
	ID        int
	Text      string
	IsDone    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

var ErrNilText = errors.New("no text given")

func NewTask(text string) (*Task, error) {
	if len(text) == 0 {
		return nil, ErrNilText
	}
	now := time.Now()
	return &Task{Text: text, IsDone: false, CreatedAt: now, UpdatedAt: now}, nil
}

type TasksService struct{}

func (s *TasksService) Get(id int) (*Task, error) {
	return NewTask("This is a new task")
}

func (s *TasksService) GetAll() ([]*Task, error) {
	firstTask, _ := NewTask("First")
	tasks := []*Task{firstTask}

	return tasks, nil
}
