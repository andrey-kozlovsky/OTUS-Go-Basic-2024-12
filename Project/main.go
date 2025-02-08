package main

import (
	"fmt"

	"github.com/andrey-kozlovsky/OTUS-Go-Basic-2024-12/Project/internal/model/notification"
	"github.com/andrey-kozlovsky/OTUS-Go-Basic-2024-12/Project/internal/model/task"
	"github.com/andrey-kozlovsky/OTUS-Go-Basic-2024-12/Project/internal/model/user"
)

func main() {
	// Создаем задачу
	t := task.NewTask(1, "Купить продукты", "Молоко, хлеб, яйца")
	fmt.Printf("Задача: %+v\n", t)
	fmt.Printf("ID задачи: %d\n", t.GetID())
	fmt.Printf("Задача создана: %v\n", t.GetCreatedAt())

	// Отмечаем задачу как выполненную
	t.Complete()
	fmt.Printf("Задача выполнена: %v\n", t.Completed)

	// Создаем пользователя
	u := user.NewUser(1, "john_doe", "john@example.com")
	fmt.Printf("Пользователь: %+v\n", u)
	fmt.Printf("Email пользователя: %s\n", u.GetEmail())

	// Обновляем email пользователя
	u.SetEmail("john.doe@example.com")
	fmt.Printf("Новый email пользователя: %s\n", u.GetEmail())

	// Создаем уведомление
	n := notification.NewNotification(1, "У вас новое сообщение")
	fmt.Printf("Уведомление: %+v\n", n)
	fmt.Printf("Уведомление прочитано: %v\n", n.IsRead())

	// Отмечаем уведомление как прочитанное
	n.MarkAsRead()
	fmt.Printf("Уведомление прочитано: %v\n", n.IsRead())
}
