package controllers

import (
	"chat_app/api/models"
	"chat_app/api/session"
	"net/http"

	"gorm.io/gorm"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type ChatController struct{
	Router *chi.Mux
	db *gorm.DB
}

type username struct{
	Username string `json:"username"`
}

func (c *ChatController) Init(db *gorm.DB){
	c.Router = chi.NewRouter()
	c.db = db

	c.Router.Post("/addnewchat", c.addNewChat)
}

func (c *ChatController) addNewChat(res http.ResponseWriter, req *http.Request){
	newChat := models.Chat{}
	err := render.DecodeJSON(req.Body, &newChat)

	if err != nil{
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	newSession, err := session.Store.Get(req, session.SessionName)

	if err != nil{
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	user := session.GetUserFromSession(newSession)

	if user.Username == ""{
		res.WriteHeader(http.StatusUnauthorized)
		return
	}	

	result := m{}

	//Before allowing a user to insert a new group chat, check to make sure they aren't adding a duplicate chat.
	//This is achieved via an inner join to potentially return a row with a user who is associated with a chat
	//that the user sent to this route.
	c.db.Model(&models.UsersChat{}).
	Select("username").
	Joins("JOIN chats ON users_chats.chatID = chats.id").
	Where("username = ? AND chat_name = ?", user.Username, newChat.ChatName).First(&result)

	if len(result) > 0{
		res.WriteHeader(http.StatusBadRequest)
		render.JSON(res, req, m{"error":  newChat.ChatName + " already exists!"})
		return
	}

	c.db.Create(&newChat)	

	newUserChat := models.UsersChat{}

	//After inserting a new chat into the "chats" tables, insert the username of the user, as well as the ID of 
	//the newly created chat into the "users_chat" table.
	newUserChat.Name = user.Username
	newUserChat.ChatID   = newChat.ID

	c.db.Create(&newUserChat)
	render.JSON(res, req, newChat)
}

//Make a post request to the following url to retrieve the username for the user currently logged.
	// url := fmt.Sprintf("http://localhost:%s/api/users/getusername", os.Getenv("PORT"))
	// postResponse, err := http.Post(url, "application/json", nil)

	// render.DecodeJSON(postResponse.Body, &result)
	// if err != nil{
	// 	res.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }