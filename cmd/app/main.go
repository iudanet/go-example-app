package main

import (
	"log"
	"net/http"

	"github.com/iudanet/go-example-app/internal/controller"
	"github.com/iudanet/go-example-app/internal/repository"
	"github.com/iudanet/go-example-app/internal/service"
	"github.com/iudanet/go-example-app/internal/worker"
)

func main() {
	repo := repository.NewInMemoryMessageRepo()
	msgService := service.NewMessageService(repo)
	msgController := controller.NewMessageController(msgService)
	go worker.Start(repo) // Запуск воркера
	// HTTP роутинг
	http.HandleFunc("/", msgController.MessageHandler)
	log.Println("Server is running at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
