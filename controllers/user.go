package controllers

import (
	"time"
)

type User struct {
	Id          int       `form : "id" json : "id"`
	FullName    string    `form : "fullName" json : "fullName"`
	Email       string    `form : "email" json : "email"`
	Password    string    `form : "password" json : "password"`
	Dateofbirth time.Time `form : "dateofbirth" json : "dateofbirth"`
	Country     string    `form : "country" json : "country"`
	Gender      string    `form : "gender" json : "gender"`
	User_type   string    `form : "user_type" json : "user_type"`
}

type UserSearch struct {
	Id          int       `form : "id" json : "id"`
	FullName    string    `form : "fullName" json : "fullName"`
	Email       string    `form : "email" json : "email"`
	Password    string    `form : "password" json : "password"`
	Dateofbirth time.Time `form : "dateofbirth" json : "dateofbirth"`
	Country     string    `form : "country" json : "country"`
	Gender      string    `form : "gender" json : "gender"`
}

type UserResponse struct {
	Status  int    `form : "status" json : "status"`
	Message string `form : "message" json : "message"`
	Data    []User `form : "data" json : "data"`
}

type UserSearchResponse struct {
	Status  int          `form : "status" json : "status"`
	Message string       `form : "message" json : "message"`
	Data    []UserSearch `form : "data" json : "data"`
}
