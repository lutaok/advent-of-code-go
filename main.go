package main

import (
	"fmt"

	aoc2023day1 "github.com/lutaok/advent-of-code-go/2023/day1"
	aoc2023day2 "github.com/lutaok/advent-of-code-go/2023/day2"
	aoc2023day3 "github.com/lutaok/advent-of-code-go/2023/day3"
	aoc2023day4 "github.com/lutaok/advent-of-code-go/2023/day4"
	aoc2023day5 "github.com/lutaok/advent-of-code-go/2023/day5"
)

func main() {
	fmt.Println("Advent of code")

	fmt.Printf("Day 1 -> %d \n", aoc2023day1.CalibrationSum())
	fmt.Printf("Day 2 -> %d \n", aoc2023day2.CubeGames())
	fmt.Printf("Day 3 -> %d \n", aoc2023day3.CalculateEngineSchematicSum())
	fmt.Printf("Day 4 -> %d \n", aoc2023day4.CalculateWinningCardPoints())
	fmt.Printf("Day 5 -> %d \n", aoc2023day5.LowestSeedNumber())
}
