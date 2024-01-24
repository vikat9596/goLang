package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in go lang")

	content := "This needs to go in the file - virasinfotech@gmail.com"

	file, err := os.Create("./myfile.txt")

	if err != nil {
		panic(err)
	}

	lenght, err := io.WriteString(file, content)

	checkNilErr(err)
	fmt.Println("Length is: ", lenght)
	defer file.Close()
	readFile("./myfile.txt")

}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("Text data inside thge file is \n", string(databyte))

}

//check nil error

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
