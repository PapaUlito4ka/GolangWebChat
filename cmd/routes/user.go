package routes

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/papaulito4ka/golangwebchat/cmd/controllers"
	"github.com/papaulito4ka/golangwebchat/cmd/db"
	"github.com/papaulito4ka/golangwebchat/cmd/services"
)

type UserRouter struct {
	Router         *mux.Router
	UserController controllers.UserController
}

func NewUserRouter(router *mux.Router, Db *sql.DB) UserRouter {
	return UserRouter{
		Router: router,
		UserController: controllers.UserController{
			UserService: services.UserService{
				UserDb: db.UserDB{
					Db: Db,
				},
			},
		},
	}
}

func (UserRouter *UserRouter) Init() {
	UserRouter.Router.HandleFunc("/api/users/", UserRouter.UserController.FindAll)
}
