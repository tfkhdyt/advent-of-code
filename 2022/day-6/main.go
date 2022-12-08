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

func getFirstMarkerIndex(input *string, size int) int {
	for i := size; i < len(*input); i++ {
		a := strings.Split((*input)[i-size:i], "")

		if !isContainsDuplicate(a) {
			return i
		}
	}

	return 0
}

func main() {
	rawInput, err := os.ReadFile("input.txt")
	if err != nil {
		log.Panicln(err)
	}

	input := strings.TrimSpace(string(rawInput))

	part1 := getFirstMarkerIndex(&input, 4)
	part2 := getFirstMarkerIndex(&input, 14)

	fmt.Printf("part1: %v\n", part1)
	fmt.Printf("part2: %v\n", part2)
}
