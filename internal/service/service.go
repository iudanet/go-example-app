package service

import "github.com/iudanet/go-example-app/internal/repository"

// MessageService интерфейс для бизнес-логики
type MessageService interface {
	GetMessage() string
	SetMessage(msg string)
}

// Реализация MessageService
type messageService struct {
	Repo repository.MessageRepository
}

func (s *messageService) GetMessage() string {
	return s.Repo.GetMessage()
}

func (s *messageService) SetMessage(msg string) {
	s.Repo.SetMessage(msg)
}

// NewMessageService создает новый экземпляр MessageService
func NewMessageService(repo repository.MessageRepository) MessageService {
	return &messageService{Repo: repo}
}
