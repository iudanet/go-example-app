package repository

type MessageRepository interface {
	GetMessage() string
	SetMessage(msg string)
}

type InMemoryMessageRepo struct {
	message string
}

func (repo *InMemoryMessageRepo) GetMessage() string {
	return repo.message
}

func (repo *InMemoryMessageRepo) SetMessage(msg string) {
	repo.message = msg
}
