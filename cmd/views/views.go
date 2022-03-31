package views

import (
	"net/http"

	"github.com/papaulito4ka/golangwebchat/cmd/dao"
)

func Home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./assets/html/home.html",
		"./assets/html/base.html",
	}

	if err := RenderTemplate(w, files, nil); err != nil {
		return
	}
}

func Users(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./assets/html/users.html",
		"./assets/html/base.html",
	}

	data := make(map[string][]dao.UserDao)
	data["Users"] = []dao.UserDao{
		{Username: "Artem", Friends: []dao.UserDao{}},
	}

	if err := RenderTemplate(w, files, data); err != nil {
		return
	}
}

func Friends(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./assets/html/users.html",
		"./assets/html/base.html",
	}

	data := make(map[string][]dao.UserDao)
	data["Users"] = []dao.UserDao{
		{Username: "Artem", Friends: []dao.UserDao{}},
	}

	if err := RenderTemplate(w, files, data); err != nil {
		return
	}
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./assets/html/signin.html",
		"./assets/html/base.html",
	}

	if r.Method == "GET" {
		if err := RenderTemplate(w, files, nil); err != nil {
			return
		}
	}
	if r.Method == "POST" {

	}
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./assets/html/signup.html",
		"./assets/html/base.html",
	}

	if r.Method == "GET" {
		if err := RenderTemplate(w, files, nil); err != nil {
			return
		}
	}
	if r.Method == "POST" {

		username := r.Form.Get("username")
		password := r.Form.Get("password")

	}
}
