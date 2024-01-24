package main

import (
	"fmt"
)

type Customer struct {
	Name    string
	Address string
	age     int
}
type Mahasiswa struct {
	nama, alamat string
	nim          int
}

// Struct Method
func (customer Customer) sayHallo(name string) {
	fmt.Println("Hello", name, "Nama Saya Adalah", customer.Name)
}

func main() {
	var aryo Customer
	//var akmal Customer
	//fmt.Println(aryo)
	//fmt.Println(akmal)

	aryo.Name = "Ariansyah Aryo Prasetio"
	aryo.Address = "Kunyit"
	aryo.age = 19

	fmt.Println(aryo)
	fmt.Println(aryo.Name)
	fmt.Println(aryo.Address)
	fmt.Println(aryo.age)

	var akmal Mahasiswa
	//fmt.Println(akmal)
	akmal.nama = "Akmal Rahim"
	akmal.alamat = "Atu atu"
	akmal.nim = 2201301025
	fmt.Println(akmal.nama)
	fmt.Println(akmal.alamat)
	fmt.Println(akmal.nim)

	//	Struct Literal
	Lucky := Customer{
		Name:    "Lucky",
		Address: "Jawa",
		age:     20,
	}
	fmt.Println(Lucky)

	Reza := Mahasiswa{
		nama:   "M. Maireza",
		alamat: "Tambang Ulang",
		nim:    22013011,
	}
	fmt.Println(Reza)

	Lucky.sayHallo("Adi")

}
