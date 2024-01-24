package main

import "fmt"

func main() {
	var a = 12
	var b = 8
	var c = 8
	var d = 8
	var g = b + b - c*a/d
	fmt.Println(g)

	var i = 10
	i += 10
	fmt.Println(i)

	i += 5
	fmt.Println(i)

	//Unery Operator
	var j = 1
	j++
	fmt.Println(j)
	j++
	fmt.Println(j)

	j--
	fmt.Println(j)
}
