package model

import "gorm.io/gorm"

type User struct {
	*gorm.Model

	Email    string `json:"email"`
	Password string `json:"password"`
}

// type ReqCreateUser struct {
// 	Surel    string `json:"surel"`
// 	Password string `json:"password"`
// }

// func coba() {
// 	user := User{}
// 	requestBody := ReqCreateUser{
// 		Surel:    "fahmy@mail.com",
// 		Password: "1234",
// 	}
// 	user.Email = requestBody.Surel
// 	user.Password = requestBody.Password
// 	//
// 	user2 := User{
// 		Email:    requestBody.Surel,
// 		Password: requestBody.Password,
// 	}
// }

// request body
// {
// 	"surel":"",
// 	"kunci":""
// }
