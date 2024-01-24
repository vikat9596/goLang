package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welocome := "Welcome to user input"
	fmt.Println(welocome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the Rating for our Pizza:")

	//comma ok //err ok

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating", input)
}
