package session

import (
	"chat_app/api/models"
	"encoding/gob"
	"os"

	"github.com/gorilla/sessions"
)

var Store *sessions.CookieStore

const SessionName string = "session-token"
const SessionLength int  = 60 * 20 //20 minutes

func init() {
	Store = sessions.NewCookieStore([]byte(os.Getenv("MY_SECRET_KEY")))
	Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   SessionLength, 
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
