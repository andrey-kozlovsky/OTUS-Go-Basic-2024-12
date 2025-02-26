package task

import (
	"time"
)

// Task представляет сущность задачи.
type Task struct {
	id          int
	Title       string
	Description string
	Completed   bool
	createdAt   time.Time
}

// NewTask создает новую задачу.
func NewTask(id int, title, description string) *Task {
	return &Task{
		id:          id,
		Title:       title,
		Description: description,
		Completed:   false,
		createdAt:   time.Now(),
	}
}

// GetID возвращает ID задачи.
func (t *Task) GetID() int {
	return t.id
}

// GetCreatedAt возвращает время создания задачи.
func (t *Task) GetCreatedAt() time.Time {
	return t.createdAt
}

// Complete отмечает задачу как выполненную.
func (t *Task) Complete() {
	t.Completed = true
}
