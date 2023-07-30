package main

import (
	"fmt"
)

func calmDown() {
	fmt.Println("Before recover")
	fmt.Println(recover())
	fmt.Println("After recover")
	fmt.Println("This is the second line after the recover")
}
func freakOut() {
	defer calmDown()
	fmt.Println("This code execute before calmDown")
}
func main() {
	freakOut()
	fmt.Println("Exiting normally")
}
