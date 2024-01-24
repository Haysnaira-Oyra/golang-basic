package main

import "fmt"

func main() {

	//Bisa ditulis gini dan juga
	//const firstName string = "Ryo"
	//const lastName = "Oyra"

	const (
		firstName string = "Ryo"
		lastName         = "Oyra"
	)

	//firstName = "Otong"
	//lastName = "ucup"
	//	Jadi Const gk bisa diubah
	fmt.Println(firstName)
	fmt.Println(lastName)
	//	Constant boleh bisa di declare tapi gk dipake
}
