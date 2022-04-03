package main

import (
	"fmt"
	"net/http"

	"github.com/papaulito4ka/golangwebchat/cmd/global"
	"github.com/papaulito4ka/golangwebchat/cmd/routes"
	"github.com/papaulito4ka/golangwebchat/cmd/views"
)

func main() {
	if err := global.Init(); err != nil {
		fmt.Println(err.Error())
		return
	}

	routes.InitAPI()

	views.InitViews()

	http.ListenAndServe(":8080", global.Handler)
}
