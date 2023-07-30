package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func OpenFile(filename string) (*os.File, error) {
	fmt.Println("Opening", filename)
	return os.Open(filename)
}

func CloseFile(file *os.File) {
	file.Close()
}

func GetFloat(fileName string) ([]float64, error) {
	var numbers []float64

	file, err := OpenFile(fileName)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil

}

func main() {
	numbers, err := GetFloat(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	sum := 0.0
	for _, number := range numbers {
		sum += number
	}
	fmt.Printf("Sum : %0.2f\n", sum)
}
