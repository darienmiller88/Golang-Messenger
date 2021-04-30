package controllers

import (
	"chat_app/api/models"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"gorm.io/gorm"
)

type MessageController struct{
	Router *chi.Mux
	db     *gorm.DB
}

func (m *MessageController) Init(db *gorm.DB){
	m.Router = chi.NewRouter()
	m.db = db

	m.Router.Use(render.SetContentType(render.ContentTypeJSON))
	m.Router.Post("/", m.sendChatMessages)
	m.Router.Get("/", m.getChatMessages)
}

func (m *MessageController) sendChatMessages(res http.ResponseWriter, req *http.Request){

}

func (m *MessageController) getChatMessages(res http.ResponseWriter, req *http.Request){
	messages := []models.Message{}

	m.db.Find(&messages)
	render.JSON(res, req, messages)
}