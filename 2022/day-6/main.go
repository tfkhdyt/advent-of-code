package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

func isContainsDuplicate[T comparable](slice []T) bool {
	tempSlice := []T{}

	for _, char := range slice {
		if slices.Contains(tempSlice, char) {
			return true
		} else {
			tempSlice = append(tempSlice, char)
		}
	}

	return false
}

func getFirstMarkerIndex(input *string, size int, ch chan int) {
	defer close(ch)
	for i := size; i < len(*input); i++ {
		a := strings.Split((*input)[i-size:i], "")

		if !isContainsDuplicate(a) {
			ch <- i
			return
		}
	}
}

func main() {
	rawInput, err := os.ReadFile("input.txt")
	if err != nil {
		log.Panicln(err)
	}

	input := strings.TrimSpace(string(rawInput))

	part1 := make(chan int)
	go getFirstMarkerIndex(&input, 4, part1)

	part2 := make(chan int)
	go getFirstMarkerIndex(&input, 14, part2)

	for {
		select {
		case packet, ok := <-part1:
			if ok {
				fmt.Printf("packet (part 1): %v\n", packet)
				part1 = nil
			}
		case message, ok := <-part2:
			if ok {
				fmt.Printf("message (part 2): %v\n", message)
				part2 = nil
			}
		}

		if part1 == nil && part2 == nil {
			break
		}
	}
}
