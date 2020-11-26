package main

import (
	"github.com/go-acme/lego/v3/log"
)

func printNumber(numChan chan int) {
	var result int
	for i := 0; i <= 100; i++ {
		result += i
	}
	numChan <- result
}

func printChar(strChan chan string) {
	var result string
	for i := 'A'; i < 'A'+26; i++ {
		result += string(i)
	}
	strChan <- result
}

func main() {
	chanPrintNumber := make(chan int)
	chanPrintChar := make(chan string)

	go printNumber(chanPrintNumber)
	go printChar(chanPrintChar)

	log.Println("Kết quả từ printNumber:", <-chanPrintNumber)
	log.Println("Kết quả từ printChar:", <-chanPrintChar)

	log.Println("main finished")
}
