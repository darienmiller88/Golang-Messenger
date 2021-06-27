package models

import(
	"gorm.io/gorm"
)

type Chat struct{
	gorm.Model
	ChatName string `json:"chatname" gorm:"type:string; size:50"`	
}