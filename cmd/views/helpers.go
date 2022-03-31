package views

import (
	"errors"
	"html/template"
	"log"
	"net/http"

	"github.com/papaulito4ka/golangwebchat/cmd/dto"
	"github.com/papaulito4ka/golangwebchat/cmd/middleware"
)

func RenderTemplate(w http.ResponseWriter, files []string, data interface{}) error {

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return err
	}

	err = ts.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return err
	}

	return nil
}

func CreateSession(w http.ResponseWriter, r *http.Request, user dto.UserDto) error {
	session, _ := middleware.Store.Get(r, middleware.SESSION_NAME)
	if !session.IsNew {
		return errors.New("Session already exists")
	}

	session.Values["id"] = user.Id
	session.Values["username"] = user.Username
	err := session.Save(r, w)
	if err != nil {
		return err
	}

	return nil
}
