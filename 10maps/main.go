package main

import "fmt"

func main() {
	fmt.Println("Maps in Go Lang")

	laungages := make(map[string]string)

	laungages["JS"] = "JavaScript"
	laungages["C"] = "C Language"
	laungages["RB"] = "Ruby"
	laungages["PY"] = "Python"

	fmt.Println("List of all Languages : ", laungages)
	fmt.Println("PY shorts for : ", laungages["PY"])

	//loops are interesting in go lang

	for key, value := range laungages {
		fmt.Printf("For key %v value is %v \n", key, value)
	}

}
