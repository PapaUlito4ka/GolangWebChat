package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
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
	})
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, req)
		log.Printf("%s %s %s", req.Method, req.RequestURI, time.Since(start))
	})
}
