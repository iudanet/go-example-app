package controller

import (
	"net/http"
	"text/template"

	"github.com/iudanet/go-example-app/internal/service"
)

type MessageController interface {
	MessageHandler(w http.ResponseWriter, r *http.Request)
}

func NewMessageController(s service.MessageService) MessageController {
	return &messageController{service: s}
}

type messageController struct {
	service service.MessageService
}

func (c *messageController) MessageHandler(w http.ResponseWriter, r *http.Request) {
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
	tmpl.Execute(w, data)
}
