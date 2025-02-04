package repository

type MessageRepository interface {
	GetMessage() string
	SetMessage(msg string)
}

type inMemoryMessageRepo struct {
	message string
}

// NewInMemoryMessageRepo создает новый экземпляр MessageRepository
func NewInMemoryMessageRepo() MessageRepository {
	return &inMemoryMessageRepo{}
}

func (repo *inMemoryMessageRepo) GetMessage() string {
	return repo.message
}

func (repo *inMemoryMessageRepo) SetMessage(msg string) {
	repo.message = msg
}
