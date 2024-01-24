package main

import "fmt"

func sayHelloWithFilter(name string, filter func(string2 string) string) {
	filteredname := filter(name)
	fmt.Println("Hello", filteredname)
}

func spamFilter(name string) string {
	if name == "Asu" {
		return "..."
	} else {
		return name
	}
}
func main() {
	sayHelloWithFilter("Aryo", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Asu", filter)
}
