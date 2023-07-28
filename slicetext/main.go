package main

import "fmt"

func main() {
	var slice1 []string
	fmt.Println(len(slice1))
	fmt.Printf("%#v\n", slice1)
	slice1 = append(slice1, "")
	fmt.Println(len(slice1))

	fmt.Printf("%#v", slice1)
	var intSlice []int
	intSlice = append(intSlice, 27)

}
