package main

import "fmt"

func main() {
	fmt.Println("Fuctions in Go Lang")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is : ", result)

	proRes := proAdder(2, 4, 5, 6, 6, 3)
	fmt.Println("Pro result is: ", proRes)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func greeter() {
	fmt.Println("Namastey from Go Lang")
}
