package controllers

import (
	"Golang-Messenger/api/models"
	"fmt"
	"time"

	"github.com/go-chi/chi"
	socketio "github.com/googollee/go-socket.io"
	"gorm.io/gorm"
)

type SocketController struct {
	SocketServer *socketio.Server
	Router       *chi.Mux
	db           *gorm.DB
	maxMessages  float64
	windowLength float64
	startTime    time.Time
	currentTime  time.Time
	allowance    float64
	users        map[string]int
}

func (s *SocketController) Init(db *gorm.DB) {
	s.SocketServer = socketio.NewServer(nil)
	s.Router = chi.NewRouter()
	s.db = db
	s.startTime = time.Now()
	s.maxMessages = 1.0
	s.windowLength = 1.0

	s.openSocketServer()
	//s.Router.Use(httprate.LimitByIP(5, 10*time.Second))
	s.Router.Handle("/", s.SocketServer)
}

func (s *SocketController) openSocketServer() {
	s.SocketServer.OnConnect("/", func(c socketio.Conn) error {
		fmt.Println(c.ID() + ": connected")
		c.Join("public_chat")
		return nil
	})

	s.SocketServer.OnEvent("/", "from_client", func(c socketio.Conn, message models.Message) {
		s.currentTime = time.Now()
		elapsed := s.currentTime.Sub(s.startTime).Seconds()
		fmt.Println(elapsed, "seconds have passed since this function was hit")
		s.startTime = s.currentTime
		s.allowance += elapsed * (float64(s.maxMessages) / float64(s.windowLength))
		newChat := models.Chat{ChatName: message.ChatName}

		if s.allowance > float64(s.maxMessages) {
			s.allowance = float64(s.maxMessages) // throttle
		}

		if s.allowance >= 1.0 {
			fmt.Println(c.ID()+":", message.MessageContent)
			s.allowance -= 1.0
			s.db.Create(&newChat)
			message.ChatID = newChat.ID
			s.db.Create(&message)
		} else {
			fmt.Println("message not added")
		}
	})

	s.SocketServer.OnDisconnect("/", func(c socketio.Conn, s string) {
		fmt.Println(c.ID() + ": disconnected")
		c.Leave("public_chat")
	})
}
