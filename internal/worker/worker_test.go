package worker

import (
	"testing"
	"time"
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

func TestWorker(t *testing.T) {
	// Создаем мок-репозиторий
	mockRepo := &MockMessageRepo{}

	// Запускаем воркер в отдельной горутине
	go Start(mockRepo)

	// Ждем некоторое время, чтобы воркер мог обновить сообщение
	time.Sleep(6 * time.Second) // Ждем больше, чем 5 секунд, чтобы гарантировать, что воркер сработал

	// Проверяем, установлено ли сообщение
	expectedPrefix := "Current Time: "
	if got := mockRepo.GetMessage(); got[:len(expectedPrefix)] != expectedPrefix {
		t.Errorf("GetMessage() = %v, want it to start with %v", got, expectedPrefix)
	}
}
