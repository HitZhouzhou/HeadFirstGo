// average calculates the average of several numbers
package main

import (
	"fmt"
	"github.com/headfirstgo/datafile"
	"log"
)

func main() {
	numbers, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	sum := 0.0
	for _, num := range numbers {
		sum += num

	}
	fmt.Println("The average is ", sum/float64(len(numbers)))

}
