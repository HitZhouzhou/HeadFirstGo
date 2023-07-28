package main

import "fmt"

type MyType string

func (m MyType) sayHi() {
	fmt.Println("Hi from", m)
}
func main() {
	value := MyType("monkey and little pig")
	value.sayHi()

}
