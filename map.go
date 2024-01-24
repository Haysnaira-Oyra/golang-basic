package main

import "fmt"

func main() {
	//var person map[string]string = map[string]string{}
	//person["name"] = "aryo"
	//person["address"] = "kunyit"
	//fmt.Println(person["name"])

	person := map[string]string{
		"name":    "Aryo,",
		"address": "kunyit",
	}
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Aryo"
	book["yeayy"] = "benar"

	fmt.Println(book)
	delete(book, "yeayy")
	fmt.Println(book)
}
