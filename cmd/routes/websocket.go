package routes

import (
	"github.com/gorilla/mux"
	"github.com/papaulito4ka/golangwebchat/cmd/controllers"
)

type WebsocketRouter struct {
	Router              *mux.Router
	WebsocketController controllers.WebsocketController
}

func NewWebsocketRouter(router *mux.Router) WebsocketRouter {
	return WebsocketRouter{
		Router:              router,
		WebsocketController: controllers.NewWebsocketController(),
	}
}

func (websocketRouter *WebsocketRouter) Init() {
	websocketRouter.Router.HandleFunc("/ws/chat/{chat_id:[0-9]+}/", websocketRouter.WebsocketController.HandleChatRoom)
}
