package main

import (
	"sync"

	"github.com/go-acme/lego/v3/log"
)

func printNumber(wg *sync.WaitGroup, numChan chan int) {
	var result int
	for i := 0; i <= 100; i++ {
		result += i
	}
	numChan <- result
	wg.Done()
}

func printChar(wg *sync.WaitGroup, strChan chan string) {
	var result string
	for i := 'A'; i < 'A'+26; i++ {
		result += string(i)
	}
	strChan <- result
	wg.Done()
}

func main() {
	chanPrintNumber := make(chan int)
	chanPrintChar := make(chan string)

	var wg sync.WaitGroup
	wg.Add(2)

	go printNumber(&wg, chanPrintNumber)
	go printChar(&wg, chanPrintChar)

	log.Println("Kết quả từ printNumber:", <-chanPrintNumber)
	log.Println("Kết quả từ printChar:", <-chanPrintChar)

	wg.Wait()

	log.Println("main finished")
}
