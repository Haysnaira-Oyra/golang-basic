package main

import (
	"Golang-Dasar/database"
	_ "Golang-Dasar/internal" //dikasih _ underscore aja solusinya
	//Cara mengatasi Golang Komplain Golang-Dasar/database biar gk ilang
	"fmt"
)

func main() {
	fmt.Println(database.GetDataBase())
}
