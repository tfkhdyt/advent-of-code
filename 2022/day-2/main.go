package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type H map[string]uint

func main() {
	resultMap := map[string]map[string]uint{
		"X": H{
			"A": 3,
			"B": 0,
			"C": 6,
		},
		"Y": H{
			"A": 6,
			"B": 3,
			"C": 0,
		},
		"Z": H{
			"A": 0,
			"B": 6,
			"C": 3,
		},
	}

	splittedByLine := strings.Split(input, "\n")
	var totalScore uint = 0

	for _, line := range splittedByLine {
		if line == "" {
			continue
		}

		choice := strings.Split(line, " ")
		opponent := choice[0]
		user := choice[1]

		result := resultMap[user][opponent]
		switch user {
		case "X":
			result += 1
		case "Y":
			result += 2
		case "Z":
			result += 3
		}

		totalScore += result
	}

	fmt.Printf("totalScore: %v\n", totalScore)
}
