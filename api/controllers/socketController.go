package controllers

import (
	"fmt"
	"log"

	"github.com/go-chi/chi"
	"github.com/nkovacs/go-socket.io"
)

type SocketController struct{
	SocketServer *socketio.Server
	Router *chi.Mux
}

func (s *SocketController) Init(){
	var err error
	s.SocketServer, err = socketio.NewServer(nil)
	s.Router = chi.NewRouter()

	if err != nil{
		log.Fatal(err)
	}
	
	s.openSocketServer()
	s.Router.Get("/socket.io/",  s.SocketServer.ServeHTTP)
	s.Router.Post("/socket.io/", s.SocketServer.ServeHTTP)
}

func (s *SocketController) openSocketServer(){
	s.SocketServer.On("connection", func (so socketio.Socket)  {
		fmt.Println("client connected:", so.Id())
		so.Join("chat")
		fmt.Println(s.SocketServer.Count())
		so.On("message", func(message string){
			fmt.Println("client", so.Id(), ":", message)
			so.BroadcastTo("chat", "new_message", so.Id() + ":" + message)
		})

		so.On("disconnect", func(reason string) {
			fmt.Println("reason:", reason)
		})
	})
}

