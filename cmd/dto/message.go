package dto

import "github.com/papaulito4ka/golangwebchat/cmd/dao"

type MessageDto struct {
	User    UserDto `json:"user"`
	Message string  `json:"message"`
	Time    string  `json:"time"`
}

func ToMessageDto(messageDao dao.MessageDao) MessageDto {
	return MessageDto{
		User:    ToUserDto(messageDao.User),
		Message: messageDao.Message,
		Time:    messageDao.Time.Format("Jan _2 15:04:05"),
	}
}
