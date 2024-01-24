package main

import "fmt"

// Function kosong
// nilai yang dikembalikan bebas
// boleh int,bool atau string
func Ups() any {
	return "Aku Adalah Any"
}
func main() {
	var kosong any = Ups()
	fmt.Println(kosong)
}
