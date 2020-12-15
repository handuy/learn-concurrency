package main

import (
	"time"

	"github.com/go-acme/lego/v3/log"
)

// googleSearch
func googleSearch(result chan string) {
	time.Sleep(1 * time.Second)
	result <- "found from Google"
}

// bingSearch
func bingSearch(result chan string) {
	time.Sleep(1 * time.Second)
	result <- "found from Bing"
}

func main() {
	chanGoogle := make(chan string)
	chanBing := make(chan string)
	go googleSearch(chanGoogle)
	go bingSearch(chanBing)

	select {
	case result := <-chanGoogle:
		log.Println(result)
	case result := <-chanBing:
		log.Println(result)
	}

	log.Println("finished main")
}