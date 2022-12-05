package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	inputFile, err := os.Open("input.txt")
	defer inputFile.Close()
	if err != nil {
		log.Panicln(err)
	}

	fileScanner := bufio.NewScanner(inputFile)

	var part1, part2 uint = 0, 0

	for fileScanner.Scan() {
		var s1, e1, s2, e2 uint
		fmt.Sscanf(fileScanner.Text(), "%d-%d,%d-%d", &s1, &e1, &s2, &e2)

		if (s1 <= s2 && e1 >= e2) || (s2 <= s1 && e2 >= e1) {
			part1++
		}

		if s1 <= e2 && e1 >= s2 {
			part2++
		}
	}

	fmt.Printf("Part 1: %v\n", part1)
	fmt.Printf("Part 2: %v\n", part2)
}
