package service

import "github.com/iudanet/go-example-app/internal/repository"

type MessageService struct {
	Repo repository.MessageRepository
}

func (s *MessageService) GetMessage() string {
	return s.Repo.GetMessage()
}

func (s *MessageService) SetMessage(msg string) {
	s.Repo.SetMessage(msg)
}
