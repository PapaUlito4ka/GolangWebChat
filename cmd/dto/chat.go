package dto

import "github.com/papaulito4ka/golangwebchat/cmd/dao"

type ChatDto struct {
	Id     int64   `json:"chat_id"`
	User   UserDto `json:"user"`
	Friend UserDto `json:"friend"`
}

func ToChatDto(chatDao dao.ChatDao) ChatDto {
	return ChatDto{
		Id:     chatDao.Id,
		User:   ToUserDto(chatDao.User),
		Friend: ToUserDto(chatDao.Friend),
	}
}
