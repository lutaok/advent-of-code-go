package aoc2023day7

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CardGameSet struct {
	hand     []string
	handType int
	bid      int
}

var seeds = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	// "J": 11, // Part 1
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
	"J": 1, // Part 2
}

func CamelCards() int {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	textInput, err := os.ReadFile(currentDir + "/2023/day7/day7-input.txt")

	if err != nil {
		log.Fatal(err)
	}

	buffer := textInput

	var gameCards []CardGameSet

	for {
		advance, token, err := bufio.ScanLines(buffer, true)

		if err != nil {
			log.Fatal(err)
		}

		if advance == 0 {
			break
		}

		if advance <= len(buffer) {
			hand := strings.Split(string(token), " ")

			if len(hand) == 2 {
				cards := strings.Split(hand[0], "")
				bid, _ := strconv.Atoi(hand[1])

				handType := determineHandType(cards)

				gameCard := CardGameSet{
					hand:     cards,
					handType: handType,
					bid:      bid,
				}

				gameCards = append(gameCards, gameCard)
			}

			buffer = buffer[advance:]
		}
	}

	sort.SliceStable(gameCards, func(i, j int) bool {
		a := gameCards[i]
		b := gameCards[j]

		if a.handType == b.handType {

			for k := 0; k < len(a.hand); k++ {
				if a.hand[k] == b.hand[k] {
					continue
				}

				return compareCards(a.hand[k], b.hand[k])
			}
		}

		return b.handType > a.handType

	})

	var rankSum int

	for index, card := range gameCards {
		rankSum += (index + 1) * card.bid
	}

	return rankSum
}

func determineHandType(hand []string) int {
	var handType int

	seedsCount := make(map[string]int)

	for _, v := range hand {

		_, ok := seedsCount[v]

		if ok {
			seedsCount[v] += 1
		} else {
			seedsCount[v] = 1
		}
	}

	n := len(seedsCount)
	if n == 1 {
		handType = 7
	} else if n == 2 {
		handType = 6

		for _, v := range seedsCount {
			if v == 2 || v == 3 {
				handType = 5
			}
		}

		_, ok := seedsCount["J"]

		if ok {
			handType = 7
		}

	} else if n == 3 {
		handType = 4

		for _, v := range seedsCount {
			if v == 2 {
				handType = 3
			}
		}

		jokerCount, ok := seedsCount["J"]

		if ok {
			if jokerCount == 1 {
				if handType == 3 {
					handType = 5
				} else if handType == 4 {
					handType = 6
				}
			} else if jokerCount == 2 {
				if handType == 3 {
					handType = 6
				} else {
					handType = 5
				}
			} else if jokerCount == 3 {
				handType = 6
			}
		}
	} else if n == 4 {
		handType = 2

		_, ok := seedsCount["J"]

		if ok {
			handType = 4
		}

	} else if n == 5 {
		handType = 1

		_, ok := seedsCount["J"]

		if ok {
			handType = 2
		}
	}

	return handType
}

func compareCards(card1 string, card2 string) bool {
	card1Strength := seeds[card1]
	card2Strength := seeds[card2]

	return card2Strength > card1Strength
}
