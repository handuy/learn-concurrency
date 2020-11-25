package main

import (
	"fmt"
	"time"
)

func printNumber() {
	for i := 0; i <= 100; i++ {
		fmt.Printf("%d ", i)
	}
}

func printChar() {
	for i := 'A'; i < 'A'+26; i++ {
		fmt.Printf("%c ", i)
	}
}

func main() {
	go printNumber()
	go printChar()
	time.Sleep(3 * time.Second)
}