package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices")

	var fruitList = []string{"Apple", "Tomato", "peach"}
	fmt.Printf("Type of data: %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Guava")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 345
	highScores[2] = 683
	highScores[3] = 284
	//highScores[4] = 777

	highScores = append(highScores, 555, 666, 999)

	fmt.Println(highScores)
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	//how to remove value from slices based on index

	var courses = []string{"reactjs", "Solidity", "Ruby", "Python"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
