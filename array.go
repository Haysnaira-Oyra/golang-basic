package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Aryo"
	names[1] = "Prasetio"
	names[2] = "Ariansyah"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		80,
		70, //data terakhir wajib pake koma
	}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	//	Array yang gk ditentukan panjangnya
	var array = [...]int{
		10, 80, 89, 21, 87, 128, 12,
	}
	fmt.Println(array)
	fmt.Println(len(array)) //untuk menghitung panjang array

}
