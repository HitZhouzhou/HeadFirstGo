package main

import (
	"fmt"
	"github.com/headfirstgo/calendar"
	"log"
)

func main() {
	data := calendar.Data{}
	err := data.SetYear(2023)
	if err != nil {
		log.Fatal(err)

	}
	err = data.SetMonth(7)
	if err != nil {
		log.Fatal(err)
	}
	err = data.SetDay(27)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data.Year())
	fmt.Println(data.Month())
	fmt.Println(data.Day())

}
