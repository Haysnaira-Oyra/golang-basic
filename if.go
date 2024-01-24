package main

import "fmt"

func main() {
	name := "ucup"

	if name == "Hayat" {
		fmt.Println("Hello Hayat")
	} else if name == "Ryo" {
		fmt.Println("Hello Ryo")
	} else if name == "Hayat" {
		fmt.Println("Hallo lagi Hayat")
	} else {
		fmt.Println("Hi, Boleh Kenalan?")
	}

	length := len(name)
	if length > 5 {
		fmt.Println("Nama Terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
