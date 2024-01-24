package main

import "fmt"

func getHello(name string) string {
	Hello := "Hello " + name
	return Hello
}

func main() {
	result := getHello("Aryo")
	fmt.Println(result)

	fmt.Println(getHello("Hayat"))
}
