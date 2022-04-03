package services

import (
	"database/sql"

	"github.com/papaulito4ka/golangwebchat/cmd/db"
	"github.com/papaulito4ka/golangwebchat/cmd/dto"
)

type UserService struct {
	UserDb db.UserDB
}

func NewUserService(Db *sql.DB) UserService {
	return UserService{UserDb: db.NewUserDB(Db)}
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

func (us *UserService) FindById(id int64) (dto.UserDto, error) {
	user, err := us.UserDb.FindById(id)
	if err != nil {
		return dto.UserDto{}, err
	}

	userDto := dto.ToUserDto(user)

	return userDto, nil
}

func (us *UserService) Find(username, password string) (dto.UserDto, error) {
	user, err := us.UserDb.Find(username, password)
	if err != nil {
		return dto.UserDto{}, err
	}

	userDto := dto.ToUserDto(user)

	return userDto, nil
}

func (us *UserService) Create(username, password string) (int64, error) {
	userId, err := us.UserDb.Create(username, password)
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func (us *UserService) Friends(id int64) ([]dto.UserDto, error) {
	users, err := us.UserDb.Friends(id)
	if err != nil {
		return []dto.UserDto{}, nil
	}

	usersDto := []dto.UserDto{}
	for _, user := range users {
		usersDto = append(usersDto, dto.ToUserDto(user))
	}

	return usersDto, nil
}

func (us *UserService) AddFriend(userId, friendId int64) (int64, error) {
	chatId, err := us.UserDb.AddFriend(userId, friendId)
	if err != nil {
		return 0, err
	}
	return chatId, nil
}

func (us *UserService) Chats(id int64) ([]dto.ChatDto, error) {
	chats, err := us.UserDb.Chats(id)
	if err != nil {
		return []dto.ChatDto{}, nil
	}

	chatsDto := []dto.ChatDto{}
	for _, chat := range chats {
		chatsDto = append(chatsDto, dto.ToChatDto(chat))
	}

	return chatsDto, nil
}
