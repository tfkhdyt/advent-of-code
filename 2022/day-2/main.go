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
	resultMapPartOne := map[string]map[string]uint{
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

	resultMapPartTwo := map[string]map[string]uint{
		"X": H{
			"A": 0,
			"B": 0,
			"C": 0,
		},
		"Y": H{
			"A": 3,
			"B": 3,
			"C": 3,
		},
		"Z": H{
			"A": 6,
			"B": 6,
			"C": 6,
		},
	}

	splittedByLine := strings.Split(input, "\n")
	var totalScorePartOne uint = 0
	var totalScorePartTwo uint = 0

	for _, line := range splittedByLine {
		if line == "" {
			continue
		}

		choice := strings.Split(line, " ")
		firstColumn := choice[0]
		secondColumn := choice[1]

		resultPartOne := resultMapPartOne[secondColumn][firstColumn]
		resultPartTwo := resultMapPartTwo[secondColumn][firstColumn]

		switch secondColumn {
		case "X":
			resultPartOne += 1
			switch firstColumn {
			case "A":
				resultPartTwo += 3
			case "B":
				resultPartTwo += 1
			case "C":
				resultPartTwo += 2
			}
		case "Y":
			resultPartOne += 2
			switch firstColumn {
			case "A":
				resultPartTwo += 1
			case "B":
				resultPartTwo += 2
			case "C":
				resultPartTwo += 3
			}
		case "Z":
			resultPartOne += 3
			switch firstColumn {
			case "A":
				resultPartTwo += 2
			case "B":
				resultPartTwo += 3
			case "C":
				resultPartTwo += 1
			}
		}

		totalScorePartOne += resultPartOne
		totalScorePartTwo += resultPartTwo
	}

	fmt.Printf("Total score part one: %v\n", totalScorePartOne)
	fmt.Printf("Total score part two: %v\n", totalScorePartTwo)
}
