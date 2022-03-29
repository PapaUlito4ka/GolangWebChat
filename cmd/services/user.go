package services

import "github.com/papaulito4ka/golangwebchat/cmd/dto"

type UserService struct {
}

func (us *UserService) FindAll() dto.UsersDto {
	users := dto.UsersDto{}

	for i := 0; i < 5; i++ {
		users.Users = append(users.Users, dto.UserDto{
			Username: "Artem",
			Friends:  []uint64{1, 2, 3, 4, 5},
		})
	}

	return users
}
