package routes

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/papaulito4ka/golangwebchat/cmd/controllers"
)

type UserRouter struct {
	Router         *mux.Router
	UserController controllers.UserController
}

func NewUserRouter(router *mux.Router, Db *sql.DB) UserRouter {
	return UserRouter{
		Router:         router,
		UserController: controllers.NewUserController(Db),
	}
}

func (UserRouter *UserRouter) Init() {
	UserRouter.Router.HandleFunc("/api/users/", UserRouter.UserController.FindAll)
	UserRouter.Router.HandleFunc("/api/users/{user_id:[0-9]+}/add_friend/{friend_id:[0-9]+}/", UserRouter.UserController.AddFriend)
}
