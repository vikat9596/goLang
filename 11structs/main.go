package main

import "fmt"

func main() {
	fmt.Println("Structs in Go Lang")

	vikat := User{"Vikat", "vikat@go.dev", true, 27}
	fmt.Println(vikat)
	fmt.Printf("Vikat's details are:  %v\n", vikat)
	fmt.Printf("Vikat's details are:  %+v\n", vikat)
	fmt.Printf("Name is :  %v\n", vikat.Name)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
