package aoc2023day2

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	Id       int
	MinGreen int
	MinRed   int
	MinBlue  int
}

func CubeGames() int {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	textInput, err := os.ReadFile(currentDir + "/2023/day2/day2-input.txt")

	if err != nil {
		log.Fatal(err)
	}

	buffer := textInput
	var games []Game

	for {
		advance, token, err := bufio.ScanLines(buffer, true)
		if err != nil {
			log.Fatal(err)
		}
		if advance == 0 {
			break
		}
		var gameInfo Game
		if advance <= len(buffer) {
			result := strings.Replace(string(token), "Game ", "", 1)
			game := strings.Split(result, ":")
			if len(game) > 0 {
				id, err := strconv.Atoi(game[0])

				if err != nil {
					log.Fatal(err)
				}

				sets := strings.Split(game[1], ";")
				var minGreen int
				var minRed int
				var minBlue int
				isValidGame := true
				for _, set := range sets {
					cubes := strings.Split(set, ",")

					for _, cube := range cubes {
						cubeNoSpace := strings.TrimSpace(cube)

						var cubeCount string
						if strings.Contains(cubeNoSpace, "green") {
							cubeCount = strings.Replace(cubeNoSpace, " green", "", 1)
							greenCount, err := strconv.Atoi(cubeCount)

							if err != nil {
								log.Fatal(err)
							}

							if greenCount > 13 {
								isValidGame = false
							}

							if greenCount > minGreen {
								minGreen = greenCount
							}
						}

						if strings.Contains(cubeNoSpace, "red") {
							cubeCount = strings.Replace(cubeNoSpace, " red", "", 1)
							redCount, err := strconv.Atoi(cubeCount)

							if err != nil {
								log.Fatal(err)
							}

							if redCount > 12 {
								isValidGame = false
							}

							if redCount > minRed {
								minRed = redCount
							}
						}

						if strings.Contains(cubeNoSpace, "blue") {
							cubeCount = strings.Replace(cubeNoSpace, " blue", "", 1)

							blueCount, err := strconv.Atoi(cubeCount)

							if err != nil {
								log.Fatal(err)
							}

							if blueCount > 14 {
								isValidGame = false
							}

							if blueCount > minBlue {
								minBlue = blueCount
							}
						}
					}

				}

				if isValidGame {
					gameInfo.Id = id
				}
				gameInfo.MinGreen = minGreen
				gameInfo.MinRed = minRed
				gameInfo.MinBlue = minBlue
			} else {
				log.Fatal("no game info found")
			}
			games = append(games, gameInfo)
			buffer = buffer[advance:]
		}
	}

	var validGamesSum int
	var minCubeSetPows int

	for _, game := range games {
		validGamesSum += game.Id
		result := game.MinBlue * game.MinGreen * game.MinRed

		minCubeSetPows += result
	}

	// Commented to solve Day 2 Part 2
	// return validGamesSum
	return minCubeSetPows
}
