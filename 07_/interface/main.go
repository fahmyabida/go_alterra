package main

import "fmt"

type User struct {
	FirstName string
}

func main() {
	var iface interface{}
	iface = true
	fmt.Println(iface)

	switch iface.(type) {
	case float64:
		fmt.Println(iface.(float64) + 20.5)
	case string:
		fmt.Println("tampilkan value stringnya " + iface.(string))
	default:
		fmt.Println("tipe data not found in switch case")
	}

}
