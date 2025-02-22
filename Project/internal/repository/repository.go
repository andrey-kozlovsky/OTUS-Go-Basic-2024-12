package repository

import (
	"fmt"

	"github.com/andrey-kozlovsky/OTUS-Go-Basic-2024-12/Project/internal/model/notification"
	"github.com/andrey-kozlovsky/OTUS-Go-Basic-2024-12/Project/internal/model/task"
	"github.com/andrey-kozlovsky/OTUS-Go-Basic-2024-12/Project/internal/model/user"
)

// Storable интерфейс для всех сущностей, которые можно хранить.
type Storable interface {
	GetID() int
}

// Репозиторий для хранения данных
var (
	Tasks         []*task.Task
	Users         []*user.User
	Notifications []*notification.Notification
)

// AddItem добавляет элемент в соответствующий слайс.
func AddItem(item Storable) error {
	switch v := item.(type) {
	case *task.Task:
		Tasks = append(Tasks, v)
	case *user.User:
		Users = append(Users, v)
	case *notification.Notification:
		Notifications = append(Notifications, v)
	default:
		return fmt.Errorf("unsupported type: %T", v)
	}
	return nil
}
