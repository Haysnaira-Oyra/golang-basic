package main

import "fmt"

func main() {
	counter := 1
	for counter <= 10 {
		fmt.Println("Perulangan ke - ", counter)
		counter++
	}
	fmt.Println("Selesai")

	for hitung := 1; hitung <= 5; hitung++ {
		fmt.Println("Perulangan ke - ", hitung)
	}
	fmt.Println("Perulangan Selesai")

	names := []string{"Ariansyah", "Aryo", "Prasetio"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
	//	For Range
	nama := []string{"Ariansyah", "Aryo", "Prasetio"}
	for index, nama := range nama {
		fmt.Println("index", index, "= ", nama)

	}

}
