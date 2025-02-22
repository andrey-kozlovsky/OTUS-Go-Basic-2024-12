package main

import (
	"fmt"

	"github.com/andrey-kozlovsky/OTUS-Go-Basic-2024-12/Project/internal/repository"
	"github.com/andrey-kozlovsky/OTUS-Go-Basic-2024-12/Project/internal/service"
)

func main() {
	// Генерируем данные
	service.GenerateData()

	// Выводим данные из репозитория
	fmt.Println("Tasks:")
	for _, t := range repository.Tasks {
		fmt.Printf("ID: %d, Title: %s, Description: %s, Completed: %v\n", t.GetID(), t.Title, t.Description, t.Completed)
	}

	fmt.Println("Users:")
	for _, u := range repository.Users {
		fmt.Printf("ID: %d, Username: %s, Email: %s\n", u.GetID(), u.Username, u.GetEmail())
	}

	fmt.Println("Notifications:")
	for _, n := range repository.Notifications {
		fmt.Printf("ID: %d, Message: %s, SentAt: %v, IsRead: %v\n", n.GetID(), n.Message, n.SentAt, n.IsRead())
	}
}
