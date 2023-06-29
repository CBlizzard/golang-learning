package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// go greeter("Hello")
	// greeter("World")

	websiteList := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://gittdhub.com",
	}

	for _, website := range websiteList {
		go getStatusCode(website)
		wg.Add(1)
	}

	wg.Wait()
}

// func greeter(str string) {
// 	for i := 0; i < 3; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(str)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()

	result, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("kuch toh gadbad hai")
	} else {
		fmt.Printf("%d success code for %s \n", result.StatusCode, endpoint)
	}
}
