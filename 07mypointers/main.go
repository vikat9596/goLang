package main

import "fmt"

func main() {
	fmt.Println("Welocme to the Pointers class")

	//var ptr *int

	//fmt.Println("Value of Pointer is: ", ptr)

	myNumber := 23

	var ptr = &myNumber

	fmt.Println("Value of actual pointer is:  ", ptr)
	fmt.Println("Value in pointer is:  ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New Value is: ", myNumber)

}
