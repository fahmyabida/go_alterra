package main

import (
	"fmt"
	"sync"
	"time"
)

type User struct {
	Name string
	mx   sync.Mutex
}

func (u *User) SetName(s string) {
	u.mx.Lock()
	fmt.Println(s)
	u.Name = s
	u.mx.Unlock()
}

func main() {
	userOne := User{
		Name: "Fahmy",
		mx:   sync.Mutex{},
	}
	go userOne.SetName("Abida")
	go userOne.SetName("Asa")
	go userOne.SetName("Firdausi")
	time.Sleep(time.Second)
	fmt.Println(userOne.Name)
}
