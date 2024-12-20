package main

import (
	"fmt"

	aoc2023day1 "github.com/lutaok/advent-of-code-go/2023/day1"
	aoc2023day10 "github.com/lutaok/advent-of-code-go/2023/day10"
	aoc2023day2 "github.com/lutaok/advent-of-code-go/2023/day2"
	aoc2023day3 "github.com/lutaok/advent-of-code-go/2023/day3"
	aoc2023day4 "github.com/lutaok/advent-of-code-go/2023/day4"
	aoc2023day5 "github.com/lutaok/advent-of-code-go/2023/day5"
	aoc2023day6 "github.com/lutaok/advent-of-code-go/2023/day6"
	aoc2023day7 "github.com/lutaok/advent-of-code-go/2023/day7"
	aoc2023day8 "github.com/lutaok/advent-of-code-go/2023/day8"
	aoc2023day9 "github.com/lutaok/advent-of-code-go/2023/day9"
)

func main() {
	// TODO: refactor this into separate modules
	fmt.Println("Advent of code")

	fmt.Printf("Day 1 -> %d \n", aoc2023day1.CalibrationSum())
	fmt.Printf("Day 2 -> %d \n", aoc2023day2.CubeGames())
	fmt.Printf("Day 3 -> %d \n", aoc2023day3.CalculateEngineSchematicSum())
	fmt.Printf("Day 4 -> %d \n", aoc2023day4.CalculateWinningCardPoints())
	fmt.Printf("Day 5 -> %d \n", aoc2023day5.LowestSeedNumber())
	fmt.Printf("Day 6 -> %d \n", aoc2023day6.BoatRaces())
	fmt.Printf("Day 7 -> %d \n", aoc2023day7.CamelCards())
	fmt.Printf("Day 8 -> %d \n", aoc2023day8.Path())
	fmt.Printf("Day 9 -> %d \n", aoc2023day9.SandDetection())
	fmt.Printf("Day 10 -> %d \n", aoc2023day10.MazeFarthestPoint())
}
