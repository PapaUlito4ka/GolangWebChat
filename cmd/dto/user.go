package dto

import "github.com/papaulito4ka/golangwebchat/cmd/dao"

type UserDto struct {
	Id       int64     `json:"user_id"`
	Username string    `json:"username"`
	Friends  []UserDto `json:"friends"`
}

func ToUserDto(userDao dao.UserDao) UserDto {
	userDto := UserDto{}

	userDto.Id = userDao.Id
	userDto.Username = userDao.Username
	for _, friend := range userDao.Friends {
		userDto.Friends = append(userDto.Friends, UserDto{
			Id:       friend.Id,
			Username: friend.Username,
			Friends:  []UserDto{},
		})
	}

	return userDto
}
