package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Tanah Laut", "Kalimantan Selatan", "Indonesia"}
	var address2 *Address = &address1

	address2.City = "Pelaihari"
	fmt.Println(address1)
	fmt.Println(address1)

	//address2 = &Address{"Tevyat", "Mondo", "Elmanuk"}
	*address2 = Address{"Tevyat", "Mondo", "Elmanuk"}
	fmt.Println(address1)
	fmt.Println(address2)
}
