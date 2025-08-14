package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func getStatusCode(endpoint string) {
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Oppps website is down")
	} else {
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)

	}

}

func main() {
	
	
	defer wg.Done()
	websiteList := []string{
		"https://jmi.ac.in",
		"https://google.com",
		"https://twitter.com",
		"https://apple.com",
		"https://jmicoe.in",
	}

	for _, web := range websiteList {
		getStatusCode(web)
		wg.Add(1)

	}
	

	wg.Wait()

}
