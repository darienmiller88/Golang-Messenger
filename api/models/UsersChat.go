package models

import(
	"gorm.io/gorm"
)

type UsersChat struct{
	gorm.Model
	Name     string `json:"username" gorm:"type:string; size:20; column:user_name"`
	ChatID   uint   `json:"chat_id"  gorm:"type:int"`	

	//A "UserChat" has a "belongs to" relationship to both a "User" and to a "Chat".
	User User `gorm:"foreignKey:Name;   references:Username; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Chat Chat `gorm:"foreignKey:ChatID; references:ID"`
}