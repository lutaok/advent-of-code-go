package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	textInput, err := os.ReadFile(currentDir + "/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1 result is: %v\n", part1(textInput))
	fmt.Printf("Part 2 result is: %v\n", part2(textInput))
}

func part1(textInput []byte) int {
	leftList, rightList := getTokensList(textInput)

	var distanceSum int
	if len(leftList) != len(rightList) {
		panic("cannot pair them up")
	}

	for index, value := range leftList {
		rightNumber := rightList[index]

		distance := value - rightNumber

		distance = max(distance, -distance)

		distanceSum += distance
	}

	return distanceSum
}

func part2(textInput []byte) int {
	leftList, rightList := getTokensList(textInput)

	occurrences := make(map[int]int)

	for _, valLeft := range leftList {
		var occ int
		for _, valRight := range rightList {
			if valRight > valLeft {
				break
			}
			if valRight == valLeft {
				occ += 1
			}
		}
		if occ == 0 {
			continue
		}
		occurrences[valLeft] = occ
	}

	var similarityScore int

	for key, value := range occurrences {
		similarityScore += key * value
	}

	return similarityScore
}

func getTokensList(textInput []byte) (leftList []int, rightList []int) {
	buffer := textInput
	for {
		advance, token, err := bufio.ScanLines(buffer, true)

		if err != nil {
			log.Fatal(err)
		}

		if advance == 0 {
			break
		}

		if advance <= len(buffer) {
			tokens := strings.Split(string(token), " ")

			firstToken := tokens[0]
			secondToken := tokens[len(tokens)-1]

			firstNumber, err := strconv.Atoi(firstToken)
			if err != nil {
				panic(err)
			}
			secondNumber, err := strconv.Atoi(secondToken)
			if err != nil {
				panic(err)
			}

			leftList = append(leftList, firstNumber)
			rightList = append(rightList, secondNumber)

			buffer = buffer[advance:]
		}
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	return leftList, rightList
}
