package views

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/papaulito4ka/golangwebchat/cmd/global"
	"github.com/papaulito4ka/golangwebchat/cmd/middleware"
)

func InitViews() {
	global.Router.HandleFunc("/", Home).Methods("GET")
	global.Router.HandleFunc("/users", Users).Methods("GET")
	global.Router.HandleFunc("/signin", SignIn).Methods("GET", "POST")
	global.Router.HandleFunc("/signup", SignUp).Methods("GET", "POST")
	global.Router.HandleFunc("/logout", Logout).Methods("GET")
	global.Router.HandleFunc("/chat/{chat_id:[0-9]+}", ChatRoom).Methods("GET")
	global.Router.Handle("/friends", middleware.AuthRequired(http.HandlerFunc(Friends))).Methods("GET")
	global.Router.HandleFunc("/", NotFound).Methods("GET")

	global.Handler = middleware.Logging(global.Router)
}

func Home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./assets/html/home.html",
		"./assets/html/base.html",
	}

	data := make(map[string]interface{})
	data["Session"] = GetSession(w, r)

	if err := RenderTemplate("home.html", w, files, data); err != nil {
		return
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	DeleteSession(w, r)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func ChatRoom(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./assets/html/chat_room.html",
		"./assets/html/base.html",
	}

	params := mux.Vars(r)
	data := make(map[string]interface{})
	data["RoomName"] = params["chat_id"]
	data["Session"] = GetSession(w, r)

	if err := RenderTemplate("chat_room.html", w, files, data); err != nil {
		return
	}
}

func Users(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./assets/html/users.html",
		"./assets/html/base.html",
	}

	data := make(map[string]interface{})
	users, err := global.UserService.FindAll()
	if err != nil {
		HandleInternalServerError(err, w)
		return
	}
	data["Users"] = users
	data["Chats"], err = global.UserService.Chats(GetSession(w, r).Values["id"].(int64))
	if err != nil {
		HandleInternalServerError(err, w)
		return
	}
	data["Session"] = GetSession(w, r)

	if err := RenderTemplate("users.html", w, files, data); err != nil {
		return
	}
}

func Friends(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./assets/html/friends.html",
		"./assets/html/base.html",
	}

	data := make(map[string]interface{})
	chats, err := global.UserService.Chats(GetSession(w, r).Values["id"].(int64))
	if err != nil {
		HandleInternalServerError(err, w)
		return
	}
	data["Chats"] = chats
	data["Session"] = GetSession(w, r)

	if err := RenderTemplate("friends.html", w, files, data); err != nil {
		return
	}
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./assets/html/signin.html",
		"./assets/html/base.html",
	}

	if r.Method == "GET" {
		data := make(map[string]interface{})
		data["Session"] = GetSession(w, r)

		if err := RenderTemplate("signin.html", w, files, data); err != nil {
			return
		}
	}
	if r.Method == "POST" {
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")

		user, err := global.UserService.Find(username, password)
		if err != nil {
			HandleInternalServerError(err, w)
			return
		}

		err = CreateSession(w, r, user)
		if err != nil {
			HandleInternalServerError(err, w)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./assets/html/signup.html",
		"./assets/html/base.html",
	}

	if r.Method == "GET" {
		data := make(map[string]interface{})
		data["Session"] = GetSession(w, r)

		if err := RenderTemplate("signup.html", w, files, data); err != nil {
			return
		}
	}
	if r.Method == "POST" {
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")

		userId, err := global.UserService.Create(username, password)
		if err != nil {
			HandleInternalServerError(err, w)
			return
		}

		user, err := global.UserService.FindById(userId)
		if err != nil {
			HandleInternalServerError(err, w)
			return
		}

		err = CreateSession(w, r, user)
		if err != nil {
			HandleInternalServerError(err, w)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./assets/html/not_found.html",
		"./assets/html/base.html",
	}

	if err := RenderTemplate("not_found.html", w, files, nil); err != nil {
		return
	}
}
