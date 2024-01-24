package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"Validation Error"}
	}

	if id != "Aryo" {
		return &notFoundError{"Data tidak ditemukan"}
	}

	// OK
	return nil
}
func main() {
	//Bagaimana cara mengaksesnya
	err := SaveData("", nil)
	if err != nil {
		//terjadi error
		if validationErr, ok := err.(*validationError); ok {
			fmt.Println("Validation error", validationErr.Error())
		} else if notFoundError, ok := err.(*notFoundError); ok {
			fmt.Println("not found error:", notFoundError.Error())
		} else {
			fmt.Println("Unknown Error", err.Error())
		}
	} else {
		//sukses
		fmt.Println("Sukses")
	}
}
