package main

import (
	"fmt"
	"github.com/headfirstgo/calendar"
	"log"
)

func main() {
	event := calendar.Event{}
	err := event.SetDay(10)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println(event.Day())
}
