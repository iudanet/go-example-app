package main

import (
	"log"

	"github.com/iudanet/go-example-app/internal/config"
	"github.com/iudanet/go-example-app/internal/controller"
	"github.com/iudanet/go-example-app/internal/httpserver"
	"github.com/iudanet/go-example-app/internal/repository"
	"github.com/iudanet/go-example-app/internal/service"
	"github.com/iudanet/go-example-app/internal/worker"
)

func main() {
	cfg := config.NewConfig() // Создание конфигурации

	repo := repository.NewInMemoryMessageRepo()
	msgService := service.NewMessageService(repo)
	msgController := controller.NewMessageController(msgService)

	go worker.Start(repo) // Запуск воркера

	log.Printf("Server is running at %s:%d", cfg.ListenHost, cfg.ListenPort)
	log.Fatal(httpserver.NewHTTPServer(msgController, cfg).ListenAndServe())
}
