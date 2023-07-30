package main

import (
	"fmt"
)

func abc(channel1 chan string) {
	channel1 <- "a"
	channel1 <- "b"
	channel1 <- "c"
}
func def(channel2 chan string) {
	channel2 <- "d"
	channel2 <- "e"
	channel2 <- "f"
}
func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go abc(channel1)
	go def(channel2)
	fmt.Println(<-channel1)
	fmt.Println(<-channel2)
	fmt.Println(<-channel1)
	fmt.Println(<-channel2)
	fmt.Println(<-channel1)
	fmt.Println(<-channel2)
}
