package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr " + man.Name
}

func main() {
	aryo := Man{"Aryo"}
	aryo.Married()

	fmt.Println(aryo.Name)
}
