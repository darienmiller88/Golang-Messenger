package controllers

import (
	//"fmt"

	"Golang-Messenger/api/middlewares"
	"Golang-Messenger/api/servestatic"
	"Golang-Messenger/api/session"
	"fmt"

	"net/http"

	// "os"
	// "path/filepath"

	"github.com/go-chi/chi"
	"gopkg.in/unrolled/render.v1"
)

type ViewsController struct {
	Router           *chi.Mux
	render           *render.Render
	staticFileServer servestatic.ServeStatic
}

func (v *ViewsController) Init() {
	v.staticFileServer = servestatic.ServeStatic{}
	v.Router = chi.NewRouter()
	v.render = render.New(render.Options{
		Extensions: []string{".tmpl", ".html"},
		Directory:  "./client/templates",
	})

	//Protect home page here instead
	v.Router.Use(middlewares.Authenticate)

	//Initialize the static file server, and give it the relative path to the client folder.
	v.staticFileServer.Init("./client", v.Router)

	//Afterwards, initialize all of the routes to be served by the following methods.
	v.Router.Get("/signup", v.signUp)
	v.Router.Get("/signin", v.signIn)
	v.Router.Get("/direct-messages", v.directMessages)
	v.Router.Get("/message-history", v.messageHistory)
	v.Router.Get("/", v.home)
	v.Router.Get("/*", v.staticFileServer.StaticFileHandler)
	v.Router.NotFound(v.notFoundPage404)
}

func (v *ViewsController) signUp(res http.ResponseWriter, req *http.Request) {
	v.render.HTML(res, http.StatusOK, "signup", nil)
}

func (v *ViewsController) signIn(res http.ResponseWriter, req *http.Request) {
	fmt.Println("status code:", req.Response)
	
	v.render.HTML(res, http.StatusOK, "signin", nil)
}

func (v *ViewsController) home(res http.ResponseWriter, req *http.Request) {
	newSession, _ := session.Store.Get(req, session.SessionName)
	user := session.GetUserFromSession(newSession)

	v.render.HTML(res, http.StatusOK, "home", user.Username)
}

func (v *ViewsController) messageHistory(res http.ResponseWriter, req *http.Request) {
	newSession, _ := session.Store.Get(req, session.SessionName)
	user := session.GetUserFromSession(newSession)

	v.render.HTML(res, http.StatusOK, "messageHistory", user.Username)
}

func (v *ViewsController) directMessages(res http.ResponseWriter, req *http.Request) {
	v.render.HTML(res, http.StatusOK, "dm", nil)
}

func (v *ViewsController) notFoundPage404(res http.ResponseWriter, req *http.Request){
	v.render.HTML(res, http.StatusNotFound, "404Error", nil)
}
