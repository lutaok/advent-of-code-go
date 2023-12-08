package aoc2023day8

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Path() int {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	textInput, err := os.ReadFile(currentDir + "/2023/day8/day8-input.txt")

	if err != nil {
		log.Fatal(err)
	}

	buffer := textInput

	var directions []string
	var startingNodes []string

	nodes := make(map[string][]string)

	for {
		advance, token, err := bufio.ScanLines(buffer, true)

		if err != nil {
			log.Fatal(err)
		}

		if advance == 0 {
			break
		}
		if advance <= len(buffer) {

			if len(directions) == 0 {
				directions = strings.Split(string(token), "")

				buffer = buffer[advance:]
				continue
			}

			mapNodes := strings.Split(string(token), " = ")

			if len(mapNodes) == 2 {
				node := mapNodes[0]
				targetText := strings.ReplaceAll(mapNodes[1], "(", "")
				targetText = strings.ReplaceAll(targetText, ")", "")

				targets := strings.Split(targetText, ", ")

				// Part 2
				if strings.HasSuffix(node, "A") {
					startingNodes = append(startingNodes, node)
				}

				_, ok := nodes[node]

				if !ok {
					nodes[node] = targets
				}

			}

			buffer = buffer[advance:]
		}
	}

	var stepsCount int
	// endFound := false
	// Part 2
	currentNodes := startingNodes
	ends := make([]int, len(startingNodes))

	for k := range currentNodes {
		ends[k] = -1
	}

	allEnds := false

	var i int
	for i < len(directions) && !allEnds {
		stepsCount += 1
		currentDirection := directions[i]

		for j := 0; j < len(currentNodes); j++ {
			currentNode := currentNodes[j]
			paths := nodes[currentNode]

			var target string
			if currentDirection == "L" {
				target = paths[0]

			} else {
				target = paths[1]
			}

			currentNodes[j] = target

			if strings.HasSuffix(target, "Z") {
				ends[j] = stepsCount
			}
		}

		for _, endCount := range ends {
			if endCount == -1 {
				allEnds = false
				break
			}

			allEnds = true
		}

		i++

		if !allEnds && i == len(directions) {
			i = 0
		}
	}

	finalEnds := ends[2:]

	result := lcm(ends[0], ends[1], finalEnds...)

	return result
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}
