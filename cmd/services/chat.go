package services

import (
	"database/sql"

	"github.com/papaulito4ka/golangwebchat/cmd/db"
	"github.com/papaulito4ka/golangwebchat/cmd/dto"
)

type ChatService struct {
	ChatDb db.ChatDB
}

func NewChatService(Db *sql.DB) ChatService {
	return ChatService{
		ChatDb: db.NewChatDb(Db),
	}
}

func (chatService *ChatService) Find(userId, friendId int64) (dto.ChatDto, error) {
	chat, err := chatService.ChatDb.Find(userId, friendId)
	if err != nil {
		return dto.ChatDto{}, err
	}

	return dto.ToChatDto(chat), nil
}

func (chatService *ChatService) SendMessage(chatId, userId int64, message string) (int64, error) {
	messageId, err := chatService.ChatDb.SendMessage(chatId, userId, message)
	if err != nil {
		return 0, err
	}

	return messageId, nil
}
