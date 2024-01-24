package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch cases")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice Value is 1 and you can open")
	case 2:
		fmt.Println("Can move 2 spot")
	case 3:
		fmt.Println("Can move 3 spot")
		fallthrough
	case 4:
		fmt.Println("Can move 4 spot")
	case 5:
		fmt.Println("Can move 5 spot")
	case 6:
		fmt.Println("Can move 6 spot and roll dice again")
	default:
		fmt.Println("What was that?")
	}
}
