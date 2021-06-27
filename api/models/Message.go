package models

import(
	"gorm.io/gorm"	
)

type Message struct{
	gorm.Model            `json:"-"`

	//Username and ChatID will be the foreign keys for this model.
	Name           string `json:"user_name" gorm:"type:string; size:20; column:user_name"`
	ChatID         uint   `json:"chat_id"  gorm:"type:int"`
	
	//Omit the name of the group chat the message belongs too
	ChatName       string `json:"chatname,omitempty"  gorm:"-"`

	//These fields will represent the content of the actual message.
	MessageContent string `json:"message_content" gorm:"type:string"`
	MessageDate    string `json:"message_date"    gorm:"type:string"`

	//A "Message" has a "belongs to" relationship to both a "User" and to a "Chat". Have the "Username" and
	// "ChatID" fields of this model serve as the foreign keys, and have them reference the "Username" and
	// "ID" fields of the User and Chat models respectively.
	User 		   User   `json:"-" gorm:"-; foreignKey:Name;   references:Username constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Chat           Chat   `json:"-" gorm:"-; foreignKey:ChatID; references:ID"`
}