package session

import (
	"chat_app/api/models"
	"encoding/gob"
	"os"

	"github.com/gorilla/sessions"
)

var Store *sessions.CookieStore

const SessionName string = "session-token"

func init() {
	Store = sessions.NewCookieStore([]byte(os.Getenv("MY_SECRET_KEY")))
	Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 , //10 minutes
		HttpOnly: true,
	}

	gob.Register(models.User{})
}

func GetUserFromSession(newSession *sessions.Session) models.User {
	user, ok := newSession.Values["user"].(models.User)

	if !ok {
		return models.User{}
	}

	return user
}
