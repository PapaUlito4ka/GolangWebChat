package global

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/papaulito4ka/golangwebchat/cmd/services"
)

var DB *sql.DB
var Router *mux.Router
var Handler http.Handler

var UserService services.UserService
var ChatService services.ChatService

func Init() error {
	var err error

	Router = mux.NewRouter()
	DB, err = sql.Open("pgx", "postgres://localhost:5432/golangwebchat?user=viktormartahin")
	if err != nil {
		return err
	}
	UserService = services.NewUserService(DB)
	ChatService = services.NewChatService(DB)

	return nil
}
