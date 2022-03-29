package routes

import (
	"github.com/gorilla/mux"
	"github.com/papaulito4ka/golangwebchat/cmd/controllers"
	"github.com/papaulito4ka/golangwebchat/cmd/services"
)

type UserRouter struct {
	Router         *mux.Router
	UserController controllers.UserController
}

func NewUserRouter(router *mux.Router) UserRouter {
	return UserRouter{
		Router: router,
		UserController: controllers.UserController{
			UserService: services.UserService{},
		},
	}
}

func (UserRouter *UserRouter) Init() {
	UserRouter.Router.HandleFunc("/api/users/", UserRouter.UserController.FindAll)
}
