package service

import (
	"github.com/andrey-kozlovsky/OTUS-Go-Basic-2024-12/Project/internal/model/notification"
	"github.com/andrey-kozlovsky/OTUS-Go-Basic-2024-12/Project/internal/model/task"
	"github.com/andrey-kozlovsky/OTUS-Go-Basic-2024-12/Project/internal/model/user"
	"github.com/andrey-kozlovsky/OTUS-Go-Basic-2024-12/Project/internal/repository"
)

// GenerateData создает данные и передает их в репозиторий.
func GenerateData() {
	// Создаем задачу
	t := task.NewTask(1, "Купить продукты", "Молоко, хлеб, яйца")
	repository.AddItem(t)

	// Создаем пользователя
	u := user.NewUser(1, "john_doe", "john@example.com")
	repository.AddItem(u)

	// Создаем уведомление
	n := notification.NewNotification(1, "У вас новое сообщение")
	repository.AddItem(n)
}
