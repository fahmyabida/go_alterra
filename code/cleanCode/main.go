package main

import (
	"fmt"
	"time"
)

const ONE_DAY = 24 * time.Hour

func main() {
	fmt.Println(divide(8, 2))

}

func divide(devidend, devisor int) int {
	var result int
	if devisor > devidend {
		return -1
	} else {
		result = devidend / devisor
		return result
	}
	return result
}

type User struct {
	firstName   string
	lastName    string
	yearOfBirth int
}

func AddOrUpdateFirstName(isAdd bool, user User) User {
	if isAdd {
		userNew := user
		return userNew
	} else {
		userUpdate := user
		return userUpdate
	}
}

func Add(user User)    {}
func Update(user User) {}
