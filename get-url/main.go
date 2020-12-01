package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/go-acme/lego/v3/log"
)

func getUrl(url string, result chan string) {
	resp, errGetUrl := http.Get(url)
	if errGetUrl != nil {
		log.Println(errGetUrl)
		return
	}
	
	urlContent, errReadBody := ioutil.ReadAll(resp.Body)
	if errReadBody != nil {
		log.Println(errReadBody)
		return
	}

	result <- string(urlContent)
	defer resp.Body.Close()
}

func main() {
	firstChanUrl := make(chan string)
	secondChanUrl := make(chan string)
	thirdChanUrl := make(chan string)

	f, err := os.OpenFile("./save.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Println(err)
		return
	}

	listChanUrl := []chan string{firstChanUrl, secondChanUrl, thirdChanUrl}
	listUrl := []string{"https://youtube.com", "https://google.com", "https://vnexpress.net"}

	for i := 0; i < len(listUrl); i++ {
		go getUrl(listUrl[i], listChanUrl[i])
	}

	for i := 0; i < len(listUrl); i++ {
		_, err := f.WriteString(<-listChanUrl[i])
		if err != nil {
			log.Println(err)
			return
		}
	}

	log.Println("main finished")
}
