package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockService struct {
    message string
}

func (m *MockService) GetMessage() string {
    return m.message
}

func (m *MockService) SetMessage(msg string) {
    m.message = msg
}

func TestMessageHandler(t *testing.T) {
    // Используем мок-сервис
    mockService := &MockService{message: "Hello, Test!"}
    msgController := NewMessageController(mockService)

    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    // Запустим тестовый сервер
    w := httptest.NewRecorder()
    handler := http.HandlerFunc(msgController.MessageHandler)

    handler.ServeHTTP(w, req)

    // Проверяем статус-код
    if status := w.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    // Проверяем содержимое ответа
    expected := `<html><body><h1>Hello, Test!</h1></body></html>`
    if w.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v", w.Body.String(), expected)
    }
}
