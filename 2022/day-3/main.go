package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func common(strs ...string) int {
loop:
	for _, r := range strs[0] {
		for _, s := range strs[1:] {
			if !strings.ContainsRune(s, r) {
				continue loop
			}
		}
    return strings.IndexRune(" abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", r)
	}
	return 0
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Panicln(err)
	}

	split := strings.Fields(strings.TrimSpace(string(input)))
	part1, part2 := 0, 0

	for i, str := range split {
		part1 += common(str[:len(str)/2], str[len(str)/2:])

		if i%3 == 0 {
			part2 += common(split[i : i+3]...)
		}
	}

	fmt.Printf("part1: %v\n", part1)
  fmt.Printf("part2: %v\n", part2)
}
