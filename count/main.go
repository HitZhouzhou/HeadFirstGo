package main

import (
	"fmt"
	"github.com/headfirstgo/datafile"
	"log"
)

func main() {
	lines, err := datafile.GetStrings("vote.txt")
	if err != nil {
		log.Fatal(err)
	}
	counters := make(map[string]int)

	for _, name := range lines {
		counters[name]++
	}

	for name, count := range counters {
		fmt.Printf("Votes for %s:%v\n", name, count)
	}
}
