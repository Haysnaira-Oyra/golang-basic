package helper

import "fmt"

var version = "1.0.0"      //Version ini tidak bisa diakses karena poke huruf kecil
var Application = "Golang" //ini bisa diakses asal Pake huruf besar nulis awalnya

func sayGoodBye(name string) string {
	return "Bye " + name
} //gk bisa diakses diluar helper

func SayHello(name string) string {
	return "Hello " + name
}

func Contoh() {
	sayGoodBye("Oyra")
	fmt.Println(version)
}
