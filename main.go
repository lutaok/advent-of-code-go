package main

import (
	"fmt"

	aoc2023day1 "github.com/lutaok/advent-of-code-go/2023/day1"
	aoc2023day2 "github.com/lutaok/advent-of-code-go/2023/day2"
	aoc2023day3 "github.com/lutaok/advent-of-code-go/2023/day3"
)

func main() {
	fmt.Println("Advent of code")

	fmt.Printf("Day 1 -> %d \n", aoc2023day1.CalibrationSum())
	fmt.Printf("Day 2 -> %d \n", aoc2023day2.CubeGames())
	fmt.Printf("Day 3 -> %d \n", aoc2023day3.CalculateEngineSchematicSum())
}