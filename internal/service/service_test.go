package service

import (
	"testing"
)

// MockMessageRepo для тестирования
type MockMessageRepo struct {
    message string
}

func (m *MockMessageRepo) GetMessage() string {
    return m.message
}

func (m *MockMessageRepo) SetMessage(msg string) {
    m.message = msg
}

func TestMessageService(t *testing.T) {
    // Подготовим мок-репозиторий
    mockRepo := &MockMessageRepo{}
    service := NewMessageService(mockRepo)

    // Тест установки и получения сообщения
    service.SetMessage("Test Message")
    if got := service.GetMessage(); got != "Test Message" {
        t.Errorf("GetMessage() = %v, want %v", got, "Test Message")
    }

    // Тест установки другого сообщения
    service.SetMessage("Another Test Message")
    if got := service.GetMessage(); got != "Another Test Message" {
        t.Errorf("GetMessage() = %v, want %v", got, "Another Test Message")
    }
}
