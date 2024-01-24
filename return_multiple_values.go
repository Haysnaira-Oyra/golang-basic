package main

import "fmt"

func getFullName() (string, string) {
	return "Aryo", "Prasetio"
}
func main() {
	//firstName, lastName := getFullName()
	//fmt.Println(firstName, lastName)
	//	buat ngeingnore salah satu di kasih _

	firstName, _ := getFullName()
	fmt.Println(firstName)

}
