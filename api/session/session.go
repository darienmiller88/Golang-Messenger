package session

import (
	"chat_app/api/models"
	"encoding/gob"
	"fmt"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

var Store *sessions.CookieStore

func init(){
	authKeyOne       := securecookie.GenerateRandomKey(64)
	encryptionKeyOne := securecookie.GenerateRandomKey(32)

	Store = sessions.NewCookieStore(authKeyOne, encryptionKeyOne)
	Store.Options = &sessions.Options{
		MaxAge:   0,
		HttpOnly: true,
	}

	fmt.Println("initialized")
	gob.Register(models.User{})
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

