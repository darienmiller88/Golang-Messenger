package models

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct{
	gorm.Model
	Username string `json:"username" gorm:"type:string; size:20; unique; not null"`
	Password string `json:"password" gorm:"type:string"`
	Authenticated bool `gorm:"-"`//Exclude column from users table
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	
	fmt.Println("called before save")
	fmt.Println("password: ", u.Password)
	return nil
}