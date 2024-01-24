package main

import "fmt"

// kode program factorial for loop
func factorialLoop(value int) int {
	result := 1

	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

func FactorialRecrusive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * FactorialRecrusive(value-1)
	}
}

func main() {
	//result := 5 * 4 * 3 * 2 * 1
	//fmt.Println(result)
	fmt.Println(factorialLoop(5))
	fmt.Println(FactorialRecrusive(5))
}
