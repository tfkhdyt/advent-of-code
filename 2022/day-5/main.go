package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	rawInput, err := os.Open("input.txt")
	if err != nil {
		log.Panicln(err)
	}

	fileScanner := bufio.NewScanner(rawInput)

	stacks := [][]string{
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

	for fileScanner.Scan() {
		line := fileScanner.Text()
		var moveN, from, to uint8
		fmt.Sscanf(line, "move %d from %d to %d", &moveN, &from, &to)

		for i := 0; i < int(moveN); i++ {
			moveStack(&stacks, from, to)
		}
	}
	fmt.Printf("stacks: %v\n", stacks)

  var topStack = []string{}
  for _, v := range stacks {
    topStack = append(topStack, v[len(v)-1]) 
  }

  fmt.Printf("topStack: %v\n", strings.Join(topStack, ""))
}

func moveStack(stacks *[][]string, from uint8, to uint8) {
	// get popped crate
	crate := (*stacks)[from-1][len((*stacks)[from-1])-1]

	// remove last crate on "from" stack
	(*stacks)[from-1] = (*stacks)[from-1][:len((*stacks)[from-1])-1]

	// push crate to new stack
	(*stacks)[to-1] = append((*stacks)[to-1], crate)

	// fmt.Printf("crate: %v\n", string(crate))
}
