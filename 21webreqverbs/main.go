package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to Web verb video")

	//PerformGetRequest()
	//PerformPostJsonRequest()
	PerformPostFormRequest()

}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code : ", response.StatusCode)
	fmt.Println("Content Length : ", response.ContentLength)

	var responseString strings.Builder      //2nd method
	content, _ := io.ReadAll(response.Body) //1st

	byteCount, _ := responseString.Write(content) //2nd

	fmt.Println("Byte Count: ", byteCount) //2nd
	fmt.Println(responseString.String())   //2nd

	fmt.Println(string(content)) //1st

}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	//fake json payload/ data

	requestBody := strings.NewReader(`
		{
			"coursename": "Let's go with golang",
			"price": 0,
			"platform" : "learncodeonline.in"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))

}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	//form data
	data := url.Values{}
	data.Add("firstname", "Vikat")
	data.Add("lastname", "Rane")
	data.Add("Email", "vikat9596@gmail.com")

	response, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}
