package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/papaulito4ka/golangwebchat/cmd/routes"
)

func main() {
	r := mux.NewRouter()

	ur := routes.NewUserRouter(r)
	ur.Init()

	http.ListenAndServe(":8080", r)
}
