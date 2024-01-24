package main

import "fmt"

func getNamaLengkap() (firstName string, middleName string, lastName string) {
	firstName = "Ariansyah"
	middleName = "Aryo"
	lastName = "Prasetio"

	return firstName, middleName, lastName
}

func main() {
	a, b, c := getNamaLengkap()
	fmt.Println(a, b, c)
}
