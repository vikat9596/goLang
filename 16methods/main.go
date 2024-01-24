package main

import "fmt"

func main() {
	fmt.Println("Structs in Go Lang")

	vikat := User{"Vikat", "vikat@go.dev", true, 27}
	fmt.Println(vikat)
	fmt.Printf("Vikat's details are:  %v\n", vikat)
	fmt.Printf("Vikat's details are:  %+v\n", vikat)
	fmt.Printf("Name is :  %v\n", vikat.Name)
	vikat.GetStatus()
	vikat.NewMail()

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is User Active: ", u.Status)
}

// passes on copy of obj
func (u User) NewMail() {
	u.Email = "test@gi.dev"
	fmt.Println("Email of  this user is :", u.Email)
}
