package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/papaulito4ka/golangwebchat/cmd/middleware"
	"github.com/papaulito4ka/golangwebchat/cmd/routes"
	"github.com/papaulito4ka/golangwebchat/cmd/views"
)

func main() {
	r := mux.NewRouter()
	db, err := sql.Open("pgx", "postgresql://localhost:5432/golangwebchat?user=viktormartahin")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	userRouter := routes.NewUserRouter(r, db)

	userRouter.Init()

	r.HandleFunc("/", views.Home).Methods("GET")
	r.HandleFunc("/users", views.Users).Methods("GET")
	r.HandleFunc("/signin", views.SignIn).Methods("GET", "POST")
	r.HandleFunc("/signup", views.SignUp).Methods("GET", "POST")
	r.Handle("/friends", middleware.AuthRequired(http.HandlerFunc(views.Friends))).Methods("GET")

	handler := middleware.Logging(r)

	http.ListenAndServe(":8080", handler)
}
