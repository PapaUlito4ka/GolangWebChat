package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/papaulito4ka/golangwebchat/cmd/dto"
	"github.com/papaulito4ka/golangwebchat/cmd/services"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(Db *sql.DB) UserController {
	return UserController{UserService: services.NewUserService(Db)}
}

func (UserController *UserController) FindAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	usersDto, err := UserController.UserService.FindAll()
	if err != nil {
		fmt.Print(w, "Users fetch fail")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	users, err := json.Marshal(usersDto)
	if err != nil {
		fmt.Print(w, "Json marshall error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(users)
}

func (UserController *UserController) AddFriend(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r)
	userId, _ := strconv.Atoi(params["user_id"])
	friendId, _ := strconv.Atoi(params["friend_id"])
	chatId, err := UserController.UserService.AddFriend(int64(userId), int64(friendId))
	if err != nil {
		fmt.Print(w, `{"message":`+err.Error()+`}`)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	chat, err := json.Marshal(dto.ChatDto{Id: chatId})
	if err != nil {
		fmt.Print(w, "Json marshall error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(chat)
}
