package main

import "fmt"

func main() {
	name := "Hayat"

	switch name {
	case "Hayat":
		fmt.Println("Hello Hayat")
	case "Aryo":
		fmt.Println("Hello Aryo")
	default:
		fmt.Println("Hi Boleh Kenalan?")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama Kamu Sudah Benar")
	}

	name = "Hayat Khizar"
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama Telalu Panjang")
	case length > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama Sudah Benar")

	}
}
