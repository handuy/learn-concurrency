package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var min = 100
var max = 200
var processTime int

func init() {
	rand.Seed(time.Now().UnixNano())
}

func readUrl(wg *sync.WaitGroup, goroutineId int, urlChan chan int, output chan string) {
	for url := range urlChan {
		// Giả lập thời gian xử lý crawl
		processTime = rand.Intn(max-min+1) + min
		time.Sleep(time.Duration(processTime) * time.Millisecond)
		// result := fmt.Sprintf("Goroutine số %d xử lý url thứ %d mất %d", goroutineId, url, processTime)
		// fmt.Println(result)

		output <- fmt.Sprintf("Goroutine số %d xử lý url thứ %d mất %d", goroutineId, url, processTime)
	}
	wg.Done()
}

func getResult(resultChan chan string) {
	for value := range resultChan {
		fmt.Println(value)
	}
}

func putJobsToChan(urlList chan int) {
	for i := 0; i < 1000; i++ {
		urlList <- i
	}
	close(urlList)
}

func main() {
	numberOfGoroutines := 5
	urlList := make(chan int, 50)
	result := make(chan string, 50)

	var wg sync.WaitGroup
	wg.Add(numberOfGoroutines)

	for i := 0; i < numberOfGoroutines; i++ {
		go readUrl(&wg, i, urlList, result)
	}

	go putJobsToChan(urlList)

	go getResult(result)

	wg.Wait()
	
	close(result)
}
