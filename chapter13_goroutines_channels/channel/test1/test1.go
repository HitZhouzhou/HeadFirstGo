package main

import (
	"fmt"
)

func greeting(mychannel chan string) {
	mychannel <- "hello"
}
func main() {
	mychannel := make(chan string)
	go greeting(mychannel)
	fmt.Println(<-mychannel)
}
