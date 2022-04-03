package routes

import "github.com/papaulito4ka/golangwebchat/cmd/global"

func InitAPI() {
	UserRouter := NewUserRouter(global.Router, global.DB)
	WebsocketRouter := NewWebsocketRouter(global.Router)

	UserRouter.Init()
	WebsocketRouter.Init()
}
