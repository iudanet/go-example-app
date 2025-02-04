package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/iudanet/go-example-app/internal/controller"
	"github.com/iudanet/go-example-app/internal/repository"
	"github.com/iudanet/go-example-app/internal/service"
)

func main() {
	repo := &repository.InMemoryMessageRepo{}
	msgService := service.NewMessageService(repo)
	msgController := controller.NewMessageController(msgService)
	go startWorker(repo) // Запуск воркера
	// HTTP роутинг
	http.HandleFunc("/", msgController.MessageHandler)
	log.Println("Server is running at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Periodic worker function
func startWorker(repo repository.MessageRepository) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			nodeMessage := fmt.Sprintf("Current Time: %s", time.Now().Format(time.RFC1123))
			repo.SetMessage(nodeMessage)
		}
	}
}
