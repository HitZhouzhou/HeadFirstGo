package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	URL  string
	Size int
}

func main() {
	size := make(chan Page)
	urls := []string{"https://example.com/", "https://github.com/HitZhouzhou/HeadFirstGo"}
	for _, url := range urls {
		go responseSize(url, size)
	}
	for i := 0; i < len(urls); i++ {
		page := <-size
		fmt.Printf("%s:%d\n", page.URL, page.Size)
	}
}
func responseSize(url string, channel chan Page) {
	fmt.Println("Getting", url)

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	channel <- Page{URL: url, Size: len(body)}
}
