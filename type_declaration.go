package main

import "fmt"

func main() {
	type NoKTP string

	var ktpRyo NoKTP = "9999999999"
	var contoh string = "222222222"
	var contohKTP NoKTP = NoKTP(contoh)

	fmt.Println(ktpRyo)
	fmt.Println(contohKTP)

}
