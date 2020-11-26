package main

import (
	"io/ioutil"
	"strings"

	"github.com/go-acme/lego/v3/log"
)

func countFirstFile(result chan int, filePath string, keyword string) {
	// Logic tính số lần xuất hiện của từ khóa trong file
	var numberOfOcc int
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println(err)
		result <- 0
		return
	}
	numberOfOcc = strings.Count( string(fileContent), keyword )
	
	result <- numberOfOcc
	defer close(result)
}

func countSecondFile(result chan int, filePath string, keyword string) {
	// Logic tính số lần xuất hiện của từ khóa trong file
	var numberOfOcc int
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println(err)
		result <- 0
		return
	}
	numberOfOcc = strings.Count( string(fileContent), keyword )
	
	result <- numberOfOcc
	defer close(result)
}

func main() {
	countFirstChan := make(chan int)
	countSecondChan := make(chan int)

	go countFirstFile(countFirstChan, "1.txt", "bạn")
	go countSecondFile(countSecondChan, "2.txt", "bạn")

	log.Println("Tổng số lần xuất hiện:", <-countFirstChan + <-countSecondChan)

	log.Println("main finished")
}