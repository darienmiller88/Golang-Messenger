package models

import(
	"gorm.io/gorm"	
)

type Message struct{
	gorm.Model
	Username       string `json:"username"        gorm:"type:string"`
	MessageContent string `json:"message_content" gorm:"type:string"`
	MessageDate    string `json:"message_Date"    gorm:"type:string"`
}

//Amazon - $250.29
//Ebay   - $229.40