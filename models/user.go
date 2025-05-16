package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       int    `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	Id    int    `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}
