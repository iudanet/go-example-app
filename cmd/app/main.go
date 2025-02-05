package main

import (
	"log"
	"net/http"
	"time"

	"github.com/iudanet/go-example-app/internal/controller"
	"github.com/iudanet/go-example-app/internal/repository"
	"github.com/iudanet/go-example-app/internal/service"
	"github.com/iudanet/go-example-app/internal/worker"
)

// setupRoutes создает и возвращает настроенный ServeMux.
func setupRoutes(msgController controller.MessageController) *http.ServeMux {
	mux := http.NewServeMux()
	// Регистрация обработчика для маршрута "/"
	mux.HandleFunc("/", msgController.MessageHandler)
	return mux
}

func main() {
	repo := repository.NewInMemoryMessageRepo()
	msgService := service.NewMessageService(repo)
	msgController := controller.NewMessageController(msgService)
	go worker.Start(repo) // Запуск воркера

	log.Println("Server is running at :8080")
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      setupRoutes(msgController),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
