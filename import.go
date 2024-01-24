package main

import (
	"Golang-Dasar/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("Aryo")
	fmt.Println(result)
	fmt.Println(helper.Application)
	//fmt.Println(helper.version) tidak bisa diakses karena gk huruf besar
	//fmt.Println(helper.sayGoodbye) error juga karena tidak huruf besar
}
