package controllers

import (
	//"fmt"
	"net/http"
	"chat_app/api/servestatic"
	// "os"
	// "path/filepath"

	"github.com/go-chi/chi"
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

	//Initialize the static file server, and give it the relativ path to the client folder.
	v.staticFileServer.Init("../client")

	//Afterwards, initialize all of the routes to be served by the following methods.
	v.Router.Get("/signup", v.signUp)
	v.Router.Get("/signin", v.signIn)
	v.Router.Get("/direct-messages", v.directMessages)
	v.Router.Get("/message-history", v.messageHistory)	
	v.Router.Get("/", v.home)
	v.Router.Get("/*", v.staticFileServer.StaticFileHandler)
}

func (v *ViewsController) signUp(res http.ResponseWriter, req *http.Request) {
	v.render.HTML(res, http.StatusOK, "signup", nil)
}

func (v *ViewsController) signIn(res http.ResponseWriter, req *http.Request) {
	v.render.HTML(res, http.StatusOK, "signin", nil)
}

func (v *ViewsController) home(res http.ResponseWriter, req *http.Request) {
	v.render.HTML(res, http.StatusOK, "home", nil)
}

func (v *ViewsController) directMessages(res http.ResponseWriter, req *http.Request) {
	v.render.HTML(res, http.StatusOK, "dm", nil)
}

func (v *ViewsController) messageHistory(res http.ResponseWriter, req *http.Request) {
	v.render.HTML(res, http.StatusOK, "messageHistory", nil)
}