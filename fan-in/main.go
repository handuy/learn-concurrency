package main

import (
	"fmt"

	"github.com/go-acme/lego/v3/log"
)

func printSth(msg string) chan string {
	// nhân viên order tạo 1 cái thẻ cho khách hàng
	result := make(chan string)
	
	// giao cho 1 nhân viên khác pha chế, chế biến đồ
	go func(){
		// kích hoạt cho thẻ rung lên
		for i := 0; i <= 5; i++ {
			result <- fmt.Sprintf("%s %d", msg, i)
		}
	}()

	// đưa cho khách hàng thẻ
	return result
}

func fanIn(chan1, chan2 chan string) chan string {
	c := make(chan string)

	go func(){
		for {
			select {
			case <-chan1:
				c <- <-chan1
			case <-chan2:
				c <- <-chan2
			}
		}
	}()

	return c
}

func main() {
	// order đồ ăn và nhận được thẻ chờ
	coffee := printSth("hello")
	bread := printSth("bread order")
	serve := fanIn(coffee, bread)

	// khách hàng nhận đc thông báo và ra quầy lấy đồ
	for i := 0; i <= 5; i++ {
		log.Println("receive from", <-serve)
	}

	log.Println("main finished")
}