package aoc2023day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func CalibrationSum() int {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	textInput, err := os.ReadFile(currentDir + "/2023/day1/day1-input.txt")

	if err != nil {
		log.Fatal(err)
	}

	buffer := textInput
	var lines []string

	for {
		advance, token, err := bufio.ScanLines(buffer, true)
		if err != nil {
			log.Fatal(err)
		}
		if advance == 0 {
			break
		}
		if advance <= len(buffer) {
			replacedToken := strings.ReplaceAll(string(token), "one", "one1one")
			replacedToken = strings.ReplaceAll(replacedToken, "two", "two2two")
			replacedToken = strings.ReplaceAll(replacedToken, "three", "three3three")
			replacedToken = strings.ReplaceAll(replacedToken, "four", "four4four")
			replacedToken = strings.ReplaceAll(replacedToken, "five", "five5five")
			replacedToken = strings.ReplaceAll(replacedToken, "six", "six6six")
			replacedToken = strings.ReplaceAll(replacedToken, "seven", "seven7seven")
			replacedToken = strings.ReplaceAll(replacedToken, "eight", "eight8eight")
			replacedToken = strings.ReplaceAll(replacedToken, "nine", "nine9nine")
			lines = append(lines, replacedToken)
			buffer = buffer[advance:]
		}

	}

	var digitsSum int
	for _, value := range lines {
		var firstDigit rune
		var lastDigit rune

		for _, char := range value {
			if unicode.IsDigit(char) {
				if firstDigit <= 0 {
					firstDigit = char
				}
				lastDigit = char
			}
		}

		trimmedValue := strings.Trim(fmt.Sprintf("%c%c", firstDigit, lastDigit), "\x00")

		calibrationValue, err := strconv.Atoi(trimmedValue)
		if err != nil {
			log.Fatal(err)
		}

		digitsSum += calibrationValue
	}

	return digitsSum
}
