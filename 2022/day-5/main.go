package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func moveStack1(stacks *[][]string, from uint8, to uint8) {
	// get popped crate
	crate := (*stacks)[from-1][len((*stacks)[from-1])-1]

	// remove last crate on "from" stack
	(*stacks)[from-1] = (*stacks)[from-1][:len((*stacks)[from-1])-1]

	// push crate to new stack
	(*stacks)[to-1] = append((*stacks)[to-1], crate)
}

func moveStack2(stacks *[][]string, moveN, from, to uint8) {
	// get popped n crates
	crate := (*stacks)[from-1][len((*stacks)[from-1])-int(moveN):]

	// remove n last crates on "from" stack
	(*stacks)[from-1] = (*stacks)[from-1][:len((*stacks)[from-1])-int(moveN)]

	// push crates to new stack
	(*stacks)[to-1] = append((*stacks)[to-1], crate...)
}

func getTopStack(stacks *[][]string) string {
	topStack := []string{}
	for _, v := range *stacks {
		topStack = append(topStack, v[len(v)-1])
	}

	return strings.Join(topStack, "")
}

func main() {
	// read raw input from text file
	rawInput, err := os.Open("input.txt")
	if err != nil {
		log.Panicln(err)
	}

	// create file scanner
	fileScanner := bufio.NewScanner(rawInput)

	// create initial stacks
	stacks1 := [][]string{
		{"B", "P", "N", "Q", "H", "D", "R", "T"},
		{"W", "G", "B", "J", "T", "V"},
		{"N", "R", "H", "D", "S", "V", "M", "Q"},
		{"P", "Z", "N", "M", "C"},
		{"D", "Z", "B"},
		{"V", "C", "W", "Z"},
		{"G", "Z", "N", "C", "V", "Q", "L", "S"},
		{"L", "G", "J", "M", "D", "N", "V"},
		{"T", "P", "M", "F", "Z", "C", "G"},
	}
	stacks2 := make([][]string, len(stacks1))
	copy(stacks2, stacks1)

	// loop over input by line
	for fileScanner.Scan() {
		// parse line string to variables
		line := fileScanner.Text()
		var moveN, from, to uint8
		fmt.Sscanf(line, "move %d from %d to %d", &moveN, &from, &to)

		for i := 0; i < int(moveN); i++ {
			moveStack1(&stacks1, from, to)
		}

		moveStack2(&stacks2, moveN, from, to)
	}

	// part 1
	topStack1 := getTopStack(&stacks1)
	fmt.Printf("topStack 1: %v\n", topStack1)

	// part 2
	topStack2 := getTopStack(&stacks2)
	fmt.Printf("topStack 2: %v\n", topStack2)
}

