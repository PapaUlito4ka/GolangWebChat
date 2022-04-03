package controllers

import (
	"database/sql"
	"net/http"

	"github.com/papaulito4ka/golangwebchat/cmd/services"
)

type ChatController struct {
	ChatService services.ChatService
}

func NewChatController(Db *sql.DB) ChatController {
	return ChatController{
		ChatService: services.NewChatService(Db),
	}
}

func (chatController *ChatController) Find(w http.ResponseWriter, r *http.Request) {

}
