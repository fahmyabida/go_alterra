package main

import (
	"fmt"
)

type IUser interface {
	ShowAge()
	FullName() string
	IsStillAlive()
}

type User struct {
	FirstName string
	LastName  string
	Age       int
	IsAlive   bool
}

func (u User) ShowAge() {
	fmt.Println(u.Age)
}

func (u User) FullName() string {
	return u.FirstName + " " + u.LastName
}

func (u User) IsStillAlive() {
	if u.IsAlive {
		fmt.Println("Masih hidup")
	} else {
		fmt.Println("Mati")
	}
}

func main() {
	var user1 IUser
	user1 = User{
		Age:       17,
		LastName:  "Abida",
		IsAlive:   true,
		FirstName: "Fahmy",
	}
	user1.IsStillAlive()
}
