package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}
var wg sync.WaitGroup
var mut sync.Mutex

func main() {

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

func getStatusCode(endpoint string) {
	defer wg.Done()

	result, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("kuch toh gadbad hai")
	} else {

		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d success code for %s \n", result.StatusCode, endpoint)
	}
}
