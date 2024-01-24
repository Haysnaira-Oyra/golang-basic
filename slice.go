package main

import "fmt"

func main() {
	names := [...]string{"aryo", "oyra", "oyra", "ucup", "otong", "budi"}

	slice1 := names[4:6]
	fmt.Println(slice1)
	slice2 := names[2:5]
	fmt.Println(slice2)
	slice3 := names[1:]
	fmt.Println(slice3)
	var slice4 []string = names[:]
	fmt.Println(slice4)
	var slice5 []string = names[5:]
	fmt.Println(slice5)

}
