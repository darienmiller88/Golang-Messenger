package session

import (
	"chat_app/api/models"
	"context"
	"encoding/gob"
	"fmt"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

var Store *sessions.CookieStore
var Ctx context.Context
const SessionName string = "session-token"

func init(){
	Ctx = context.Background()
	authKeyOne := securecookie.GenerateRandomKey(32)

	Store = sessions.NewCookieStore([]byte("secret-key"))
	Store.Options = &sessions.Options{
		Path: "/",
		MaxAge:   0,
		HttpOnly: true,
	}

	fmt.Println("initialized, key:", authKeyOne)
	gob.Register(models.User{})
}

func GetUserFromSession(newSession *sessions.Session) models.User{
	user, ok := newSession.Values["user"].(models.User)

	if !ok {
		return models.User{}
	}

	return user
}

//"maxAge" will determine the length of the lifespan of the cookie.
//MaxAge=0 means no Max-Age attribute specified and the cookie will be deleted after the browser session ends.
//MaxAge<0 means delete cookie immediately. 
//MaxAge>0 means Max-Age attribute present and given in seconds.
// func (s *SessionManager) InitSession(maxAge int) {
// 	authKeyOne       := securecookie.GenerateRandomKey(64)
// 	encryptionKeyOne := securecookie.GenerateRandomKey(32)

// 	s.Store = sessions.NewCookieStore(authKeyOne, encryptionKeyOne)
// 	s.Store.Options = &sessions.Options{
// 		MaxAge:   maxAge,
// 		HttpOnly: true,
// 	}

// 	gob.Register(models.User{})
// }

