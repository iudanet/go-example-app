package httpserver

import (
	"fmt"
	"net/http"
	"time"

	"github.com/iudanet/go-example-app/internal/config"
	"github.com/iudanet/go-example-app/internal/controller"
)

// setupRoutes создает и возвращает настроенный ServeMux.
func setupRoutes(msgController controller.MessageController) *http.ServeMux {
	mux := http.NewServeMux()
	// Регистрация обработчика для маршрута "/"
	mux.HandleFunc("/", msgController.MessageHandler)
	return mux
}

// NewHTTPServer создает новый HttpServer и возвращает его объект.
func NewHTTPServer(msgController controller.MessageController, cfg *config.Config) *http.Server {
	mux := setupRoutes(msgController)
	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.ListenHost, cfg.ListenPort),
		Handler:      mux,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  120 * time.Second,
	}
	return srv
}
