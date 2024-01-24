package main

import "fmt"

func endApp() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("Terjadi Panic", message)
}

func runAPP(error bool) {
	defer endApp()

	if error {
		panic("Ups Error")
	}

	//	Tambahin Recover
	//message := recover()
	//fmt.Println("Terjadi Panic", message)
	//	itu tadi cara yang salah
}

func main() {
	runAPP(true)
	fmt.Println("Aryoooooooooo")
}
