package main

import (

	"github.com/go-acme/lego/v3/log"
)

func firstInput(nums ...int) chan int {
	result := make(chan int)
	
	go func(){
		for i := 0; i < len(nums); i++ {
			result <- nums[i]
		}
		close(result)
	}()

	return result
}

func secondInput(fromFirst chan int) chan int {
	result := make(chan int)

	go func(){
		for item := range fromFirst {
			result <- item * item
		}
		close(result)
	}()

	return result
}

func main() {
	firstChan := firstInput(1,2,3,4,5,6)
	secondChan := secondInput(firstChan)

	for item := range secondChan {
		log.Println("receive", item)
	}

	log.Println("main finished")
}