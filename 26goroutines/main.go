package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup // usually these are pointers
var mut sync.Mutex    //pointer

func main() {
	// 	go greeter("Hello")
	// 	greeter("World")
	// }

	// func greeter(s string) {
	// 	for i := 0; i < 6; i++ {
	// 		time.Sleep(3 * time.Millisecond)
	// 		fmt.Println(s)
	// 	}
	// }
	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://virasinfotech.com",
		"https://google.com",
		"https://fb.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)

}

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Oops in Endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d Status code for %s\n", res.StatusCode, endpoint)
	}

}
