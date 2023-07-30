package main

import (
	"fmt"
)

func calmDown() {
	fmt.Println("Before recover")
	fmt.Println(recover())
	fmt.Println("After recover")
	fmt.Println("The second line after the recover")
}
func freakOut() {
	defer calmDown()
	panic("oh no")
}
func main() {
	freakOut()
	fmt.Println("Exiting normally")
}
