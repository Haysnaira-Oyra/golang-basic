package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

// Data kosong
func main() {
	data := NewMap("Aryo")
	//cek isinya
	if data == nil {
		fmt.Println(data["name"])
	} else {
		fmt.Println(data["name"])
	}

}
