package main

import "fmt"

func random() any {
	//var result any = random
	return 100
}

func main() {
	//var result any = random()
	//var resultstring string = result.(string)
	//fmt.Println(resultstring)
	//
	//var resultINT int = result.(int)
	//fmt.Println(resultINT)

	var result any = random()
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown", value)

	}
}
