package main

import (
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
}

var c = make(chan User)

func main() {
	channelMultiple()
}

func channelMultiple() {
	cOne := make(chan string)
	cTwo := make(chan string)
	cThere := make(chan string)
	cFour := make(chan string)
	go Dosen(cOne, cTwo, cThere, cFour)
	go Mahasiswa(cOne, cTwo, cThere, cFour)
	time.Sleep(1 * time.Second)
}

func Dosen(cOne, cTwo, cThere, cFour chan string) {
	fmt.Println("DOSEN : Silahkan mau tanya apa ?")
	cOne <- "" // sent to channel
	<-cTwo     // blocking
	fmt.Printf("DOSEN : Apa tuh ?\n")
	cThere <- "" // sent to channel
	<-cFour      // blocking
	fmt.Printf("DOSEN : ciaaa ciaaa! kamu nanti keruangan saya !")
}

func Mahasiswa(cOne, cTwo, cThere, cFour chan string) {
	<-cOne // blocking
	pertanyaan := "Bedanya kamu sama modem apa ?"
	fmt.Printf("MAHASISWA : %v\n", pertanyaan)
	cTwo <- pertanyaan // sent to channel
	<-cThere           // blocking
	fmt.Printf("MAHASISWA : Modem terkoneksi ke internet, kalo kamu terkoneksi ke hatiku\n")
	cFour <- "" // sent to channel
}

func channelBasic() {
	fmt.Println("main() started !")
	go SayHi()
	fmt.Println("gorutine run but wating channel !")
	c <- User{
		FirstName: "fahmy",
		LastName:  "abida",
	}
	fmt.Println("main() stoped !")
}

func SayHi() {
	data := <-c
	fmt.Printf("Hallo my name is %v %v \n", data.FirstName, data.LastName)
}
