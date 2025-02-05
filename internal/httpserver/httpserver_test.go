package httpserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/iudanet/go-example-app/internal/config"
	"github.com/iudanet/go-example-app/internal/controller"
	"github.com/iudanet/go-example-app/internal/repository"
)

// MockService для тестирования контроллера
type MockService struct {
	message string
}

func (m *MockService) GetMessage() string {
	return m.message
}

func (m *MockService) SetMessage(msg string) {
	m.message = msg
}

// TestNewHTTPServer проверяет, что HTTP-сервер работает корректно
func TestNewHTTPServer(t *testing.T) {
	// Создаем мок-репозиторий и сервис
	mockRepo := repository.NewInMemoryMessageRepo()
	mockRepo.SetMessage("Test Message")
	mockService := &MockService{message: "Test Message"}
	msgController := controller.NewMessageController(mockService)

	// Создаем новый HTTP-сервер
	srv := NewHTTPServer(msgController, &config.Config{
		ListenHost:   "localhost",
		ListenPort:   8080,
		ReadTimeout:  60,
		WriteTimeout: 60,
	})

	// Создаем тестовый HTTP-запрос
	req, err := http.NewRequest("GET", "/", http.NoBody)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	// Запускаем тестовый сервер
	w := httptest.NewRecorder()
	srv.Handler.ServeHTTP(w, req)

	// Проверяем статус-код
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Проверяем содержимое ответа
	expected := `<html><body><h1>Test Message</h1></body></html>`
	if w.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", w.Body.String(), expected)
	}
}
