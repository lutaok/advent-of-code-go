package aoc2023day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Gear struct {
	currentNumber string
	starRow       int
	starColumn    int
}

type StarGear struct {
	value string
	count int
}

func CalculateEngineSchematicSum() int {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	textInput, err := os.ReadFile(currentDir + "/2023/day3/day3-input.txt")

	if err != nil {
		log.Fatal(err)
	}

	buffer := textInput
	var engineMatrix [][]string
	var lineCount int
	for {
		advance, token, err := bufio.ScanLines(buffer, true)

		if err != nil {
			log.Fatal(err)
		}

		if advance == 0 {
			break
		}

		if advance <= len(buffer) {
			engineMatrix = append(engineMatrix, []string{})
			lineChars := strings.Split(string(token), "")

			engineMatrix[lineCount] = append(engineMatrix[lineCount], lineChars...)

			lineCount++

			buffer = buffer[advance:]
		}
	}

	n := len(engineMatrix)

	var partNumberSum int
	var i int

	var gears []Gear
	for i < n {
		row := engineMatrix[i]
		m := len(row)

		var j int
		for j < m {
			column := row[j]

			char, _ := utf8.DecodeRune([]byte(column))

			if char == '.' || unicode.IsSymbol(char) || unicode.IsPunct(char) {
				j++
				continue
			}

			// when number search for symbol neighbours
			if unicode.IsDigit(char) {
				gear := Gear{
					currentNumber: column,
					starRow:       -1,
					starColumn:    -1,
				}

				span := j
				for {
					if canCheckRight(i, span, n, m) {
						nextCol, _ := utf8.DecodeRune([]byte(row[span+1]))

						if nextCol == '.' || unicode.IsSymbol(nextCol) || unicode.IsPunct(nextCol) {
							break
						}

						gear.currentNumber = fmt.Sprintf("%s%c", gear.currentNumber, nextCol)
						span++
					} else {
						break
					}
				}

				var isPart bool
				var thatChar rune
				k := j
				for k <= span {
					if canCheckTop(i, k, n, m) && canCheckLeft(i, k, n, m) {
						isPart, thatChar = checkIfNeighbourIsSymbol(k-1, engineMatrix[i-1])
					}

					if isPart {
						number, _ := strconv.Atoi(gear.currentNumber)

						if thatChar == '*' {
							gear.starColumn = k - 1
							gear.starRow = i - 1
						}

						partNumberSum += number
						break
					}

					if canCheckTop(i, k, n, m) {
						isPart, thatChar = checkIfNeighbourIsSymbol(k, engineMatrix[i-1])
					}

					if isPart {
						number, _ := strconv.Atoi(gear.currentNumber)
						if thatChar == '*' {
							gear.starColumn = k
							gear.starRow = i - 1
						}
						partNumberSum += number
						break
					}

					if canCheckTop(i, k, n, m) && canCheckRight(i, k, n, m) {
						isPart, thatChar = checkIfNeighbourIsSymbol(k+1, engineMatrix[i-1])
					}

					if isPart {
						number, _ := strconv.Atoi(gear.currentNumber)

						if thatChar == '*' {
							gear.starColumn = k + 1
							gear.starRow = i - 1
						}

						partNumberSum += number
						break
					}

					if canCheckRight(i, k, n, m) {
						isPart, thatChar = checkIfNeighbourIsSymbol(k+1, engineMatrix[i])
					}

					if isPart {
						number, _ := strconv.Atoi(gear.currentNumber)

						if thatChar == '*' {
							gear.starColumn = k + 1
							gear.starRow = i
						}

						partNumberSum += number
						break
					}

					if canCheckBottom(i, k, n, m) && canCheckRight(i, k, n, m) {
						isPart, thatChar = checkIfNeighbourIsSymbol(k+1, engineMatrix[i+1])
					}

					if isPart {
						number, _ := strconv.Atoi(gear.currentNumber)

						if thatChar == '*' {
							gear.starColumn = k + 1
							gear.starRow = i + 1
						}

						partNumberSum += number
						break
					}

					if canCheckBottom(i, k, n, m) {
						isPart, thatChar = checkIfNeighbourIsSymbol(k, engineMatrix[i+1])
					}

					if isPart {
						number, _ := strconv.Atoi(gear.currentNumber)

						if thatChar == '*' {
							gear.starColumn = k
							gear.starRow = i + 1
						}

						partNumberSum += number
						break
					}

					if canCheckBottom(i, k, n, m) && canCheckLeft(i, k, n, m) {
						isPart, thatChar = checkIfNeighbourIsSymbol(k-1, engineMatrix[i+1])
					}

					if isPart {
						number, _ := strconv.Atoi(gear.currentNumber)

						if thatChar == '*' {
							gear.starColumn = k - 1
							gear.starRow = i + 1
						}

						partNumberSum += number
						break
					}

					if canCheckLeft(i, k, n, m) {
						isPart, thatChar = checkIfNeighbourIsSymbol(k-1, engineMatrix[i])
					}

					if isPart {
						number, _ := strconv.Atoi(gear.currentNumber)

						if thatChar == '*' {
							gear.starColumn = k - 1
							gear.starRow = i
						}

						partNumberSum += number
						break
					}

					k++
				}

				j += span - j
				if isPart && thatChar == '*' {
					gears = append(gears, gear)
				}
			}
			j++
		}
		i++
	}

	starPositionGears := make(map[string]StarGear)

	for _, gear := range gears {

		starPositionKey := fmt.Sprintf("%d,%d", gear.starColumn, gear.starRow)

		data, ok := starPositionGears[starPositionKey]

		if ok {
			gearValue, _ := strconv.Atoi(gear.currentNumber)
			previousValue, _ := strconv.Atoi(data.value)
			starPositionGears[starPositionKey] = StarGear{
				value: fmt.Sprintf("%d", gearValue*previousValue),
				count: data.count + 1,
			}
		} else {
			starPositionGears[starPositionKey] = StarGear{
				value: gear.currentNumber,
				count: 1,
			}
		}

	}

	var gearsSum int

	for _, starGear := range starPositionGears {
		if starGear.count == 2 {
			starGearNumber, _ := strconv.Atoi(starGear.value)
			gearsSum += starGearNumber
		}
	}

	// Commented to solve Part 2 of Day 3
	// return partNumberSum
	return gearsSum
}

func checkIfNeighbourIsSymbol(columnIndex int, row []string) (bool, rune) {
	char, _ := utf8.DecodeRune([]byte(row[columnIndex]))
	if char == '.' {
		return false, char
	}

	if unicode.IsSymbol(char) || unicode.IsPunct(char) {
		return true, char
	}

	return false, char
}

func canCheckTop(currentRowIndex int, currentColumnIndex, n int, m int) bool {
	return currentRowIndex-1 >= 0
}

func canCheckLeft(currentRowIndex int, currentColumnIndex, n int, m int) bool {
	return currentColumnIndex-1 >= 0
}

func canCheckRight(currentRowIndex int, currentColumnIndex, n int, m int) bool {
	return currentColumnIndex+1 <= m-1
}

func canCheckBottom(currentRowIndex int, currentColumnIndex, n int, m int) bool {
	return currentRowIndex+1 <= n-1
}
