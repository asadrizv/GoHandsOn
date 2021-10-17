package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	links := []string{
		"http://facebook.com",
		"http://google.com",
		"http://netflix.com",
		"http://stackoverflow.com",
	}
	var wg sync.WaitGroup

	urlChan := make(chan string, 100)

	checkUrl := func(link string) {
		_, err := http.Get(link)
		if err == nil {
			urlChan <- link + " Is working!!"
		} else {
			urlChan <- link + " Is down"
		}
	}

	renderResults := func() {
		for result := range urlChan {
			fmt.Println(result)
			wg.Done()
		}
	}

	for _, link := range links {
		wg.Add(1)
		go checkUrl(link)
	}

	go renderResults()
	wg.Wait()
}
