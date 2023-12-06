package aoc2023day6

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func BoatRaces() int {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	textInput, err := os.ReadFile(currentDir + "/2023/day6/day6-input.txt")

	if err != nil {
		log.Fatal(err)
	}

	buffer := textInput

	timeRegexp := regexp.MustCompile(`Time: \s*`)
	distanceRegexp := regexp.MustCompile(`Distance: \s*`)

	var timeNumbers []int
	var distanceNumbers []int

	for {
		advance, token, err := bufio.ScanLines(buffer, true)

		if err != nil {
			log.Fatal(err)
		}

		if advance == 0 {
			break
		}

		if advance <= len(buffer) {
			input := string(token)
			if timeRegexp.MatchString(input) {
				time := strings.Join(strings.Split(timeRegexp.ReplaceAllString(input, ""), " "), "")

				number, _ := strconv.Atoi(time)

				timeNumbers = append(timeNumbers, number)
				// Commented for Part 2
				// times := strings.Split(timeRegexp.ReplaceAllString(input, ""), " ")
				// for _, value := range times {
				// 	if value != "" {
				// 		number, _ := strconv.Atoi(value)

				// 		timeNumbers = append(timeNumbers, number)
				// 	}
				// }
			}

			if distanceRegexp.MatchString(input) {
				distance := strings.Join(strings.Split(distanceRegexp.ReplaceAllString(input, ""), " "), "")

				number, _ := strconv.Atoi(distance)

				distanceNumbers = append(distanceNumbers, number)
				// Commented for Part 2
				// distances := strings.Split(distanceRegexp.ReplaceAllString(input, ""), " ")
				// for _, value := range distances {
				// 	if value != "" {
				// 		number, _ := strconv.Atoi(value)

				// 		distanceNumbers = append(distanceNumbers, number)
				// 	}
				// }
			}

			buffer = buffer[advance:]
		}
	}

	var winningWays []int = make([]int, len(timeNumbers))

	for i := 0; i < len(timeNumbers); i++ {
		time := timeNumbers[i]
		recordDistance := distanceNumbers[i]

		startWinningIndex := -1
		endWinningIndex := -1
		minBreakRecordIndex := 1
		maxBreakRecordIndex := time - 1

		for minBreakRecordIndex < maxBreakRecordIndex {
			if startWinningIndex == -1 {
				travelledDistance := minBreakRecordIndex * (time - minBreakRecordIndex)

				if travelledDistance > recordDistance {
					startWinningIndex = minBreakRecordIndex
				}

				minBreakRecordIndex++
			}

			if endWinningIndex == -1 {
				travelledDistance := maxBreakRecordIndex * (time - maxBreakRecordIndex)

				if travelledDistance > recordDistance {
					endWinningIndex = maxBreakRecordIndex
				}

				maxBreakRecordIndex--
			}

			if startWinningIndex > -1 && endWinningIndex > -1 {
				break
			}
		}

		winningWays[i] += endWinningIndex - startWinningIndex + 1
	}

	totalWinningWays := 1

	for _, val := range winningWays {
		totalWinningWays *= val
	}

	return totalWinningWays
}
