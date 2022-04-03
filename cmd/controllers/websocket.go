package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/papaulito4ka/golangwebchat/cmd/global"
)

type WebsocketController struct {
	Clients  map[*websocket.Conn]bool
	Upgrader websocket.Upgrader
}

type WebsocketData struct {
	ChatId  string `json:"chatId"`
	UserId  int64  `json:"userId"`
	User    string `json:"user"`
	Message string `json:"message"`
}

func NewWebsocketController() WebsocketController {
	return WebsocketController{
		Clients: make(map[*websocket.Conn]bool, 0),
		Upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}
}

func (websocketController *WebsocketController) HandleChatRoom(w http.ResponseWriter, r *http.Request) {
	conn, err := websocketController.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()

	websocketController.Clients[conn] = true
	defer delete(websocketController.Clients, conn)

	for {
		mt, message, err := conn.ReadMessage()

		if err != nil || mt == websocket.CloseMessage {
			break
		}

		go websocketController.writeMessage(message)
		go fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(message))
	}
}

func (websocketController *WebsocketController) writeMessage(message []byte) {
	websocketData := WebsocketData{}
	err := json.Unmarshal(message, &websocketData)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	chatId, _ := strconv.Atoi(websocketData.ChatId)

	_, err = global.ChatService.SendMessage(int64(chatId), websocketData.UserId, websocketData.Message)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for conn := range websocketController.Clients {
		conn.WriteMessage(websocket.TextMessage, message)
	}
}
