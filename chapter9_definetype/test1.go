package main

import (
	"fmt"
)

type Gallons float64
type Liters float64

func main() {
	fmt.Println(Gallons(10) + Gallons(34))
}
