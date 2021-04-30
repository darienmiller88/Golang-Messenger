package controllers

import (
	"chat_app/api/models"
	"chat_app/api/session"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/gorilla/sessions"

	// "encoding/gob"

	renderv1 "gopkg.in/unrolled/render.v1"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type m map[string]interface{}

type UserController struct{
	Router *chi.Mux
	db     *gorm.DB
	newRender *renderv1.Render
}

func (u *UserController) Init(db *gorm.DB){
	u.Router = chi.NewRouter()
	u.db = db
	u.newRender = renderv1.New(renderv1.Options{
		Extensions: []string{".tmpl", ".html"},
		Directory: "../client/templates",
	})

	fmt.Println(session.Store.Options)

	u.Router.Use(render.SetContentType(render.ContentTypeJSON))
	u.Router.Post("/signup", u.signup)
	u.Router.Post("/signin", u.signin)
	u.Router.Post("/signout", u.signout)
	u.Router.Get("/home", u.home)
	u.Router.Get("/", u.test)
	u.Router.Get("/isauth", u.isAuth)
}

func (u *UserController) signin(res http.ResponseWriter, req *http.Request){
	newSession, _ := session.Store.Get(req, "cookie-name")
	user := models.User{}
	err  := render.DecodeJSON(req.Body, &user)
	password := user.Password

	
	if err != nil{
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	result := u.db.Where("username = ?", user.Username).First(&user).RowsAffected
	err     = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	//If the username or password is incorrect, return an error message.
	if result == 0 || err != nil{
		res.WriteHeader(http.StatusUnauthorized)
		render.JSON(res, req, m{"error_message": "Username or password incorrect. Please try again."})
		return
	}

	user.Authenticated = true
	newSession.Values["user"] = user
    newSession.Save(req, res)
	fmt.Println("values:", newSession.Values)
	render.JSON(res, req, m{})
}

func (u *UserController) signout(res http.ResponseWriter, req *http.Request){
	newSession, err := session.Store.Get(req, "cookie-name")

	if err != nil{
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

    // Revoke users authentication
    newSession.Values["user"] = models.User{}
	newSession.Options.MaxAge = -1
    newSession.Save(req, res)
	http.Redirect(res, req, "/signout", http.StatusFound)
}

func (u *UserController) signup(res http.ResponseWriter, req *http.Request){
	user := models.User{}
	err := render.DecodeJSON(req.Body, &user)	
	
	if err != nil{
		res.WriteHeader(http.StatusBadRequest)
	}

	response := u.validateSignUp(user, res, req)
	if len(response) != 0{
		render.JSON(res, req, response)
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(hashedPassword)
	u.db.Create(&user)
	http.Redirect(res, req, "/", http.StatusFound)
	//render.JSON(res, req, m{"success": true, "new_user": user})
}

//Function to validate user Sign up information, and return a response body potentially containing information
//to signal to the front end to display error messages to the user.
func (u *UserController) validateSignUp(user models.User, res http.ResponseWriter, req *http.Request) m{
	response := make(m)

	//Check to see if the username field is empty.
	if strings.ReplaceAll(user.Username, " ", "") == ""{
		response["no_username_err"] = "Enter your username!"
	}

	//Check to see if the password field is empty.
	if strings.ReplaceAll(user.Password, " ", "") == ""{
		response["no_password_err"] = "Enter your password!"
	}

	result := u.db.Where("username = ?", user.Username).First(&user)

	//If there was a user found in the database, and that user doesn't have a empty, return an error signaling that there  
	if result.RowsAffected > 0{	
		response["username_taken"] = fmt.Sprintf("Username %s is taken!", user.Username)
	}
	
	return response
}

func (u *UserController) test(res http.ResponseWriter, req *http.Request){
	var user models.User
	//users := []models.User{}

	user.Username = ""
	u.db.Where(m{"username": ""}).Find(&user)
	//u.db.Where(&models.User{Username: user.Username}).First(&user)
	render.JSON(res, req, user)
}

func (u *UserController) isAuth(res http.ResponseWriter, req *http.Request){
	newSession, _ := session.Store.Get(req, "cookie-name")
	user, ok := newSession.Values["user"].(models.User)

	if !ok{
		render.JSON(res, req,  models.User{Authenticated: true})
		return
	}

	fmt.Println("user:", user)
	render.JSON(res, req, user)
}


func (u *UserController) home(res http.ResponseWriter, req *http.Request) {
	newSession, _ := session.Store.Get(req, "cookie-name")	
	response, err := http.Get("http://localhost:7000/api/users/isauth")
	var user models.User

	if err != nil{
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	render.DecodeJSON(response.Body, &user)

	fmt.Println("user in home:", user)
	fmt.Println("values:", newSession.Values)

	render.JSON(res, req, user)
	//v.render.HTML(res, http.StatusOK, "signin", user)
}

func GetUser(s *sessions.Session) models.User {
	userPlaceholder := s.Values["user"]
	user, authenticated := userPlaceholder.(models.User)

	if !authenticated {
		return models.User{Authenticated: false}
	}

	return user
}
