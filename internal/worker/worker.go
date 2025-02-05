package worker

import (
	"fmt"
	"time"

	"github.com/iudanet/go-example-app/internal/repository"
)

// Start запускает фоновый воркер для обновления сообщений.
func Start(repo repository.MessageRepository) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		nodeMessage := fmt.Sprintf("Current Time: %s", time.Now().Format(time.RFC1123))
		repo.SetMessage(nodeMessage)
	}
}
