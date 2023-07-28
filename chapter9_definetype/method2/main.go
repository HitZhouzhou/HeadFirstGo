package main

import "fmt"

type Liters float64

type Gallons float64
type Milliliters float64

func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}
func main() {
	gas := Liters(10)
	fmt.Printf("%0.3f liters is equal to %0.3f gallons \n", gas, gas.ToGallons())

}
