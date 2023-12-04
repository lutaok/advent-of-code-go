package aoc2023day4

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

type Card struct {
	id      string
	winning []string
	owned   []string
}

func CalculateWinningCardPoints() int {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	textInput, err := os.ReadFile(currentDir + "/2023/day4/day4-input.txt")

	if err != nil {
		log.Fatal(err)
	}

	buffer := textInput
	cardIdRegexp := regexp.MustCompile(`[0-9]+`)
	cardRegexp := regexp.MustCompile(`Card\s*[0-9]+:`)
	numbersRegexp := regexp.MustCompile(`\s`)

	var points int

	var originalCards []Card

	for {
		advance, token, err := bufio.ScanLines(buffer, true)

		if err != nil {
			log.Fatal(err)
		}

		if advance == 0 {
			break
		}

		if advance <= len(buffer) {
			cardId := cardIdRegexp.FindString(string(token))

			line := cardRegexp.ReplaceAllString(string(token), "")
			cardNumbers := strings.Split(line, " | ")

			if len(cardNumbers) != 2 {
				log.Fatal("no expected format found")
			}

			winning := numbersRegexp.Split(cardNumbers[0], -1)
			owned := numbersRegexp.Split(cardNumbers[1], -1)

			card := Card{
				id:      cardId,
				winning: winning,
				owned:   owned,
			}

			originalCards = append(originalCards, card)
			buffer = buffer[advance:]
		}
	}

	var wonCards []int
	var i int
	for i < len(originalCards) {
		wonCards = append(wonCards, 1)
		i++
	}

	for index, card := range originalCards {
		winningMap := make(map[string]int)

		for _, value := range card.winning {
			if value != "" {
				winningMap[value] = 1
			}
		}

		var linePoints int
		var winningCount int
		for _, value := range card.owned {
			_, ok := winningMap[value]

			if ok {
				winningCount += 1
				if linePoints > 0 {
					linePoints *= 2
				} else {
					linePoints += 1
				}
			}
		}
		points += linePoints

		j := index + 1
		currentInstances := wonCards[index]
		for j <= index+winningCount {
			wonCards[j] = currentInstances + wonCards[j]
			j++
		}
	}

	var totalInstances int
	for _, value := range wonCards {
		totalInstances += value
	}

	// Commented to solve Day 4 Part 2
	// return points
	return totalInstances
}
