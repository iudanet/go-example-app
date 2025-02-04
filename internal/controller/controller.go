package controller

import (
	"net/http"
	"text/template"

	"github.com/iudanet/go-example-app/internal/service"
)

type MessageController struct {
	service *service.MessageService
}

func NewMessageController(s *service.MessageService) *MessageController {
	return &MessageController{service: s}
}

func (c *MessageController) MessageHandler(w http.ResponseWriter, r *http.Request) {
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
