package repository

import (
	"testing"
)

func TestInMemoryMessageRepo(t *testing.T) {
	repo := NewInMemoryMessageRepo()

	// Тест установки сообщения
	repo.SetMessage("Hello, World!")
	if got := repo.GetMessage(); got != "Hello, World!" {
		t.Errorf("GetMessage() = %v, want %v", got, "Hello, World!")
	}

	// Тест установки другого сообщения
	repo.SetMessage("Another Message")
	if got := repo.GetMessage(); got != "Another Message" {
		t.Errorf("GetMessage() = %v, want %v", got, "Another Message")
	}
}
