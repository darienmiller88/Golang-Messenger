package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"Golang-Messenger/api/session"
)

//Auth middleware to protect home page from unauthorized users, and the signin and signup from authorized users.
func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		//Check to see if the user visited either the signin or signup route
		nonRestrictedRoutesVisited := req.URL.String() == "/signin" || req.URL.String() == "/signup"
		staticPageHit := strings.HasPrefix(req.URL.String(), "/static")
		cookieName, _ := req.Cookie(session.SessionName)

		//newSession, _ := session.Store.Get(req, session.SessionName)

		//If the user is logged in (Cookie was found on client side), and they try to access either the sign up
		//or sign in page, redirect them back to the home page. On the contrary, if the user tries to
		//access any of the restricted pages while not logged in, redirect them to the sign in page
		if cookieName != nil && nonRestrictedRoutesVisited {
			http.Redirect(res, req, "/", http.StatusFound)
			return
		} else if cookieName == nil && !nonRestrictedRoutesVisited && !staticPageHit {
			fmt.Println("unauthorized route hit! Method:", req.Method, "url:", req.URL.String())
			http.Redirect(res, req, "/signin", http.StatusFound)
			return
		}

		next.ServeHTTP(res, req)
	})
}
