package main

import (
	"fmt"
	"net/http"
	"os"

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

	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, global.Handler)
}
