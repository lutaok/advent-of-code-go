package aoc2023day9

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func SandDetection() int {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	textInput, err := os.ReadFile(currentDir + "/2023/day9/day9-input.txt")

	if err != nil {
		log.Fatal(err)
	}

	buffer := textInput

	var histories [][]int

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

			var tokenNumbers []int

			for _, v := range tokens {
				valueNumber, _ := strconv.Atoi(v)

				tokenNumbers = append(tokenNumbers, valueNumber)
			}

			// For part 2 reverse input
			slices.Reverse(tokenNumbers)

			histories = append(histories, tokenNumbers)

			buffer = buffer[advance:]
		}
	}

	var sum int

	for i := range histories {
		predictionSum := 0
		calculateHistoryDifference(histories[i], &predictionSum, -1)
		m := len(histories[i])
		sum += predictionSum + histories[i][m-1]
	}

	return sum
}

func calculateHistoryDifference(history []int, predictionSum *int, difference int) {
	if checkIfAllZeros(history) {
		return
	}

	m := len(history)

	historyDifferences := make([]int, m-1)

	for j := m - 1; j > 0; j-- {
		localDifference := history[j] - history[j-1]

		historyDifferences[j-1] = localDifference

		if j == m-1 {
			difference = localDifference
		}
	}

	calculateHistoryDifference(historyDifferences, predictionSum, difference)
	*predictionSum += difference
}

func checkIfAllZeros(arr []int) bool {
	for _, v := range arr {
		if v != 0 {
			return false
		}
	}

	return true
}
