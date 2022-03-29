package controllers

import (
	"Golang-Messenger/api/models"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"gorm.io/gorm"
)

type MessageController struct {
	Router *chi.Mux
	db     *gorm.DB
}

type mm map[string]interface{}

func (m *MessageController) Init(db *gorm.DB) {
	m.db = db
	m.Router = chi.NewRouter()

	m.Router.Get("/public-messages", m.getPublicMessages)
	m.Router.Post("/get-message-history", m.getMessageHistory)
	m.Router.Delete("/delete-message", m.deleteMessage)
}

func (m *MessageController) getPublicMessages(res http.ResponseWriter, req *http.Request) {
	var messages []models.Message

	m.db.Find(&messages)

	// data, err := json.MarshalIndent(messages, "", "\t")

	// if err != nil{
	// 	res.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }

	render.JSON(res, req, messages)
}

func (m *MessageController) deleteMessage(res http.ResponseWriter, req *http.Request) {
	var message models.Message
	render.DecodeJSON(req.Body, &message)

	rowsAffected := m.db.Where(&models.Message{
		Name:           message.Name,
		MessageContent: message.MessageContent,
		MessageDate:    message.MessageDate,
	}).Delete(&models.Message{}).RowsAffected

	render.JSON(res, req, mm{"rows_affected": rowsAffected})
}

func (m *MessageController) getMessageHistory(res http.ResponseWriter, req *http.Request) {
	var message map[string]interface{}
	var messages []map[string]interface{}

	render.DecodeJSON(req.Body, &message)
	m.db.Model(&models.Message{}).Select("user_name", "message_content", "message_date").
		Joins("JOIN chats ON messages.chat_id = chats.id").
		Where("user_name = ? AND chats.chat_name = ? ", message["user_name"], message["chat_name"]).Find(&messages)
		
	render.JSON(res, req, messages)
}
