package controllers

import (
	"chat_app/api/models"
	"chat_app/api/session"
	"fmt"
	"net/http"
	"strconv"
	"unicode"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type m map[string]interface{}

type UserController struct {
	Router *chi.Mux
	db     *gorm.DB
}

func (u *UserController) Init(db *gorm.DB) {
	u.Router = chi.NewRouter()
	u.db = db

	//Todo: Add rate limiter to this route to protect against DDOS attacks.
	fmt.Println(session.Store.Options)
	u.Router.Use(render.SetContentType(render.ContentTypeJSON))
	u.Router.Post("/signup", u.signup)
	u.Router.Post("/signin", u.signin)
	u.Router.Post("/signout", u.signout)
	u.Router.Post("/getusername", u.getUserName)
}

func (u *UserController) getUserName(res http.ResponseWriter, req *http.Request) {
	newSession, err := session.Store.Get(req, session.SessionName)

	if err != nil{
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	user := session.GetUserFromSession(newSession)

	render.JSON(res, req, m{"username": user.Username})
}

func (u *UserController) signin(res http.ResponseWriter, req *http.Request) {
	user := models.User{}
	err := render.DecodeJSON(req.Body, &user)
	password := user.Password

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	result := u.db.Where("username = ?", user.Username).First(&user).RowsAffected
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	//If the username or password is incorrect, return an error message.
	if result == 0 || err != nil {
		res.WriteHeader(http.StatusUnauthorized)
		render.JSON(res, req, m{"error_message": "Username or password incorrect. Please try again."})
		return
	}

	newSession, err := session.Store.Get(req, "session-token")

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	user.Authenticated = true
	newSession.Values["user"] = user
	err = newSession.Save(req, res)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(res, req, m{})
}

func (u *UserController) signout(res http.ResponseWriter, req *http.Request) {
	newSession, err := session.Store.Get(req, session.SessionName)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Revoke users authentication
	newSession.Values["user"] = models.User{}
	newSession.Options.MaxAge = -1
	newSession.Save(req, res)

	render.JSON(res, req, m{})
}

func (u *UserController) signup(res http.ResponseWriter, req *http.Request) {
	user := models.User{}
	err := render.DecodeJSON(req.Body, &user)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	response := u.validateSignUp(user)
	if len(response) != 0 {
		render.JSON(res, req, response)
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(hashedPassword)
	u.db.Create(&user)

	//After creating the user, sign them in automatically, and redirect them to the home page
	newSession, err := session.Store.Get(req, session.SessionName)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	user.Authenticated = true
	newSession.Values["user"] = user
	err = newSession.Save(req, res)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(res, req, m{"success": true})
}

//Function to validate user Sign up information, and return a response body potentially containing information
//to signal to the front end to display error messages to the user.
func (u *UserController) validateSignUp(user models.User) m {
	response := make(m)
	usernameLen := len(user.Username)
	minimumUsernameLen := 5
	maximumUsernameLen := 20

	//Check to see if the username field is is at least n characters.
	if usernameLen < minimumUsernameLen || usernameLen > maximumUsernameLen {
		response["weak_username_err"] = fmt.Sprintf("At least %s characters, no more than %s.", 
			strconv.Itoa(minimumUsernameLen), strconv.Itoa(maximumUsernameLen)) 
	}

	//Check to see if the password has less than 8 characters or if it does not contain at least one number.
	passwordLen := 8
	passwordLongEnough := len(user.Password) < passwordLen
	passwordContainsNumber := !stringContainsNumber([]rune(user.Password))

	if passwordLongEnough || passwordContainsNumber {
		passwordErrors := []m{
			{"password_error": "At least " + strconv.Itoa(passwordLen) + " characters", "is_password_weak": passwordLongEnough},
			{"password_error": "At least one number", "is_password_weak": passwordContainsNumber},
		}

		response["password_errors"] = passwordErrors
	}

	result := u.db.Where("username = ?", user.Username).First(&user)

	//If there was a user found in the database, and that user doesn't have a empty, return an error signaling that there
	if result.RowsAffected > 0 {
		response["username_taken"] = fmt.Sprintf("Username %s is taken!", user.Username)
	}

	return response
}

func stringContainsNumber(s []rune) bool {
	for _, char := range s {
		if unicode.IsNumber(char) {
			return true
		}
	}

	return false
}
