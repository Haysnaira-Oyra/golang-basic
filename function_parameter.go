package main

import "fmt"

func SayHelloTo(firstName string, lastName string) {
	fmt.Println("Halo", firstName, lastName)
}
func main() {
	SayHelloTo("Aryo", "Prasetio")
	SayHelloTo("Hayat", "Khizar")
}
