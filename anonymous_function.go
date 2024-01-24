package main

import "fmt"

type Blacklist func(string) bool

func RegisterUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Kamu Diblokir", name)
	} else {
		fmt.Println("Selamat Datang", name)
	}
}

// Anonimous Function langsung didalam main
func main() {
	blacklist := func(name string) bool {
		return name == "Asu"
	}
	RegisterUser("Aryo", blacklist)

	RegisterUser("Asu", func(name string) bool {
		return name == "Asu"
	})
}
