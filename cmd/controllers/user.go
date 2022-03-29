package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/papaulito4ka/golangwebchat/cmd/services"
)

type UserController struct {
	UserService services.UserService
}

func (UserController *UserController) FindAll(w http.ResponseWriter, r *http.Request) {
	users, err := json.Marshal(UserController.UserService.FindAll())
	if err != nil {
		fmt.Print(w, "Json decode error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(users)
}
