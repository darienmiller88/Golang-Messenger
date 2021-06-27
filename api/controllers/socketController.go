package controllers

import (
	"chat_app/api/models"
	"fmt"
	"log"
	"gorm.io/gorm"

	"github.com/go-chi/chi"
	//"github.com/mitchellh/mapstructure"
	// "github.com/graarh/golang-socketio"
	// "github.com/graarh/golang-socketio/transport"
	// "github.com/nkovacs/go-socket.io"
	"github.com/googollee/go-socket.io"
)

type SocketController struct{
	SocketServer *socketio.Server
	Router *chi.Mux
	db *gorm.DB
}

func (s *SocketController) Init(db *gorm.DB){
	var err error
	s.SocketServer = socketio.NewServer(nil)
	s.Router = chi.NewRouter()
	s.db = db

	if err != nil{
		log.Fatal(err)
	}
	
	s.openSocketServer()	
	s.Router.Handle("/", s.SocketServer)
}

func (s *SocketController) openSocketServer(){
	s.SocketServer.OnConnect("/", func(c socketio.Conn) error {
		fmt.Println(c.ID() + ": connected")
		c.Join("public_chat")
		return nil
	}) 

	s.SocketServer.OnEvent("/", "from_client", func(c socketio.Conn, message  models.Message){
		newChat := models.Chat{ChatName: message.ChatName}

		s.db.Create(&newChat)
		message.ChatID = newChat.ID
		s.db.Create(&message)		
		
		fmt.Println(c.ID() + ":", message)
	})

	s.SocketServer.OnDisconnect("/", func(c socketio.Conn, s string) {
		fmt.Println(c.ID() + ": disconnected")
		c.Leave("public_chat")
	})
}
