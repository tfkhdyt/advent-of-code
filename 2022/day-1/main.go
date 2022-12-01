package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func main() {
	const fileName = "input.txt"
	cwd, err := os.Getwd()
	if err != nil {
		log.Panicln(err)
	}

	filePath := filepath.Join(cwd, fileName)

	rawInput, err := os.ReadFile(filePath)
	if err != nil {
		log.Panicln(err)
	}

	elvesFoodsRawCalories := strings.Split(string(rawInput), "\n\n")

	elvesFoodsCalories := []uint64{}

	for _, elfFoodsRawCalories := range elvesFoodsRawCalories {
		foodsCalories := strings.Split(elfFoodsRawCalories, "\n")
		var totalOfElfFoodsCalories uint64 = 0

		for _, foodCalories := range foodsCalories {
			if foodCalories == "" {
				continue
			}
			foodCaloriesUint, err := strconv.ParseUint(foodCalories, 10, 64)
			if err != nil {
				log.Panicln(err)
			}

			totalOfElfFoodsCalories += foodCaloriesUint
		}

		elvesFoodsCalories = append(elvesFoodsCalories, totalOfElfFoodsCalories)
	}

	sort.Slice(elvesFoodsCalories, func(i, j int) bool {
		return elvesFoodsCalories[i] > elvesFoodsCalories[j]
	})

	fmt.Printf("Most calories: %v\n", elvesFoodsCalories[0])

	var totalCaloriesOf3Elves uint64 = 0
	for i := 0; i < 3; i++ {
		totalCaloriesOf3Elves += elvesFoodsCalories[i]
	}

	fmt.Printf("Total calories carried by top 3 elves: %v", totalCaloriesOf3Elves)
}
