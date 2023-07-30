package main

import (
	"fmt"
)

func reportNap(name string, delay int) {
	for i := 0; i < delay; i++ {
		fmt.Println(name, "sleeping")
	}
	fmt.Println(name, "wakeup")
}
func send(mychannel chan string) {
	reportNap("sending goroutine", 2)
	fmt.Println("***sending value***")
	mychannel <- "a"
	fmt.Println("***sending value")
	mychannel <- "b"
}
func main() {
	mychannel := make(chan string)
	go send(mychannel)
	reportNap("receiving goroutine", 5)
	fmt.Println(<-mychannel)
	fmt.Println(<-mychannel)
}
