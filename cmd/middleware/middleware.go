package middleware

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/papaulito4ka/golangwebchat/cmd/global"
)

const (
	SESSION_NAME = "user-session"
)

var Store = sessions.NewCookieStore([]byte("some super-secret-key"))

func AuthRequired(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := Store.Get(r, SESSION_NAME)

		if session.IsNew {
			http.Redirect(w, r, "/signin", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func AuthNotRequired(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := Store.Get(r, SESSION_NAME)

		if !session.IsNew {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func HasChatAccess(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := Store.Get(r, SESSION_NAME)

		if session.IsNew {
			http.Redirect(w, r, "/signin", http.StatusSeeOther)
			return
		}

		params := mux.Vars(r)
		chatId, _ := strconv.Atoi(params["chat_id"])
		userId := session.Values["id"].(int64)
		chats, _ := global.UserService.Chats(userId)
		key := false
		for _, chat := range chats {
			if chat.Id == int64(chatId) {
				key = true
				break
			}
		}

		if !key {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
	})
}
