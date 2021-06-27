package controllers

import (
	"github.com/go-chi/chi"
	"gorm.io/gorm"
)

type APIControllers struct {
	db     *gorm.DB
	Router *chi.Mux
}

func (a *APIControllers) Init(db *gorm.DB) {
	a.db = db
	a.Router = chi.NewRouter()

	userController := UserController{}
	chatController := ChatController{}
	messageController := MessageController{}

	userController.Init(a.db)
	chatController.Init(a.db)
	messageController.Init(a.db)

	a.Router.Mount("/users", userController.Router)
	a.Router.Mount("/chats", chatController.Router)
	a.Router.Mount("/messages", messageController.Router)
}
