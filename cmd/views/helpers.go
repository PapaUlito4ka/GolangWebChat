package views

import (
	"errors"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/papaulito4ka/golangwebchat/cmd/dto"
	"github.com/papaulito4ka/golangwebchat/cmd/middleware"
)

func RenderTemplate(templateName string, w http.ResponseWriter, files []string, data interface{}) error {

	ts, err := template.New(templateName).Funcs(template.FuncMap{
		"Contains": func(chats []dto.ChatDto, user dto.UserDto) int64 {
			for _, chat := range chats {
				if chat.Friend.Id == user.Id {
					return chat.Id
				}
			}
			return 0
		},
	}).ParseFiles(files...)
	if err != nil {
		HandleInternalServerError(err, w)
		return err
	}

	err = ts.Execute(w, data)
	if err != nil {
		HandleInternalServerError(err, w)
		return err
	}

	return nil
}

func CreateSession(w http.ResponseWriter, r *http.Request, user dto.UserDto) error {
	session, _ := middleware.Store.Get(r, middleware.SESSION_NAME)
	if !session.IsNew {
		return errors.New("session already exists")
	}

	session.Values["id"] = user.Id
	session.Values["username"] = user.Username
	err := session.Save(r, w)
	if err != nil {
		return err
	}

	return nil
}

func GetSession(w http.ResponseWriter, r *http.Request) *sessions.Session {
	session, _ := middleware.Store.Get(r, middleware.SESSION_NAME)
	return session
}

func DeleteSession(w http.ResponseWriter, r *http.Request) error {
	session, _ := middleware.Store.Get(r, middleware.SESSION_NAME)
	if session.IsNew {
		return errors.New("session is new")
	}
	session.Options.MaxAge = -1

	err := session.Save(r, w)
	if err != nil {
		return err
	}

	return nil
}

func HandleInternalServerError(err error, w http.ResponseWriter) {
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
