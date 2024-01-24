package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in go lang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Cider"

	fmt.Println("Fruit List is: ", fruitList)
	fmt.Println("Length of Fruit List is: ", len(fruitList))

	var vegList = [3]string{"potata", "beans", "mushroom"}
	fmt.Println("Veg List is : ", vegList)
}
