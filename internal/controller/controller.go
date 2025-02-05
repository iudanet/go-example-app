package controller

import (
	"net/http"
	"text/template"

	"github.com/iudanet/go-example-app/internal/service"
)

// MessageController handles HTTP requests related to messages.
type MessageController interface {
	MessageHandler(w http.ResponseWriter, r *http.Request)
}

// NewMessageController creates a new instance of MessageController.
func NewMessageController(s service.MessageService) MessageController {
	return &messageController{service: s}
}

type messageController struct {
	service service.MessageService
}

func (c *messageController) MessageHandler(w http.ResponseWriter, _ *http.Request) {
	data := struct {
		Message string
	}{
		Message: c.service.GetMessage(),
	}

	tmpl, err := template.New("message").Parse(`<html><body><h1>{{.Message}}</h1></body></html>`)
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
		return
	}
}
