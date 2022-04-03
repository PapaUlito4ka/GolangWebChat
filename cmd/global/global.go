package global

import (
	"database/sql"
	"net/http"
	"os"

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
	// if err := godotenv.Load(); err != nil {
	// 	fmt.Println(err.Error())
	// }

	// host, _ := os.LookupEnv("POSTGRES_HOST")
	// port, _ := os.LookupEnv("POSTGRES_PORT")
	// db, _ := os.LookupEnv("POSTGRES_DB")
	// user, _ := os.LookupEnv("POSTGRES_USER")
	// password, _ := os.LookupEnv("POSTGRES_PASSWORD")

	Router = mux.NewRouter()
	// DB, err = sql.Open("pgx", fmt.Sprintf("postgres://%s:%s/%s?user=%s&password=%s", host, port, db, user, password))
	DB, err = sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}
	UserService = services.NewUserService(DB)
	ChatService = services.NewChatService(DB)

	return nil
}
