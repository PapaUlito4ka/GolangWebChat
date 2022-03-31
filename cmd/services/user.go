package services

import (
	"github.com/papaulito4ka/golangwebchat/cmd/db"
	"github.com/papaulito4ka/golangwebchat/cmd/dto"
)

type UserService struct {
	UserDb db.UserDB
}

func (us *UserService) FindAll() ([]dto.UserDto, error) {
	users, err := us.UserDb.FindAll()
	if err != nil {
		return []dto.UserDto{}, err
	}

	usersDto := []dto.UserDto{}
	for _, user := range users {
		usersDto = append(usersDto, dto.ToUserDto(user))
	}

	return usersDto, nil
}

func (us *UserService) Create(username, password string) (int64, error) {
	userId, err := us.UserDb.Create(username, password)
	if err != nil {
		return 0, err
	}
	return userId, nil
}
