package controllers

import (
	//"fmt"
	"chat_app/api/models"
	"chat_app/api/servestatic"
	"chat_app/api/session"

	// "encoding/gob"
	"fmt"

	//"fmt"
	"net/http"

	// "os"
	// "path/filepath"

	"github.com/go-chi/chi"
	//"github.com/gorilla/sessions"
	chi_render "github.com/go-chi/render"
	"gopkg.in/unrolled/render.v1"
)

type ViewsController struct{
	Router *chi.Mux
	render *render.Render
	staticFileServer servestatic.ServeStatic
}

func (v *ViewsController) Init() {
	v.staticFileServer = servestatic.ServeStatic{}
	v.Router = chi.NewRouter()
	v.render = render.New(render.Options{
		Extensions: []string{".tmpl", ".html"},
		Directory: "../client/templates",
	})

	//Initialize the static file server, and give it the relative path to the client folder.
	v.staticFileServer.Init("../client")

	//Afterwards, initialize all of the routes to be served by the following methods.
	v.Router.Get("/signup", v.signUp)
	v.Router.Get("/signin", v.signIn)
	v.Router.Get("/direct-messages", v.directMessages)
	v.Router.Get("/message-history", v.messageHistory)	
	v.Router.Get("/", v.home)
	v.Router.Get("/isauth", v.isauth)
	v.Router.Get("/*", v.staticFileServer.StaticFileHandler)
}

func (v *ViewsController) signUp(res http.ResponseWriter, req *http.Request) {
	v.render.HTML(res, http.StatusOK, "signup", nil)
}

func (v *ViewsController) signIn(res http.ResponseWriter, req *http.Request) {
	v.render.HTML(res, http.StatusOK, "signin", nil)
}

func (v *ViewsController) home(res http.ResponseWriter, req *http.Request) {
	newSession, _ := session.Store.Get(req, "cookie-name")	
	response, err := http.Get("http://localhost:7000/api/users/isauth")
	var user models.User

	if err != nil{
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	chi_render.DecodeJSON(response.Body, &user)

	fmt.Println("user in home:", user)
	fmt.Println("values:", newSession.Values)

	// if !user.Authenticated {
	// 	newSession.Save(req, res)
	// 	http.Redirect(res, req, "/signin", http.StatusFound)
	// 	return
	// }

	v.render.JSON(res, http.StatusOK, user)
	//v.render.HTML(res, http.StatusOK, "signin", user)
}

func (v *ViewsController) directMessages(res http.ResponseWriter, req *http.Request) {
	v.render.HTML(res, http.StatusOK, "dm", nil)
}

func (v *ViewsController) messageHistory(res http.ResponseWriter, req *http.Request) {
	v.render.HTML(res, http.StatusOK, "messageHistory", nil)
}

func (v *ViewsController) isauth(res http.ResponseWriter, req *http.Request){
	newSession, _ := session.Store.Get(req, "cookie-name")
	val := newSession.Values["user"].(models.User)

	fmt.Println("existing session:", newSession.IsNew)
	v.render.JSON(res, http.StatusOK, val)
}

// func (v *ViewsController) getUser(s *sessions.Session) models.User {
// 	userPlaceholder := s.Values["authenticated"]
// 	user, authenticated := userPlaceholder.(models.User)

// 	if !authenticated {
// 		return models.User{Authenticated: false}
// 	}

// 	return user
// }
