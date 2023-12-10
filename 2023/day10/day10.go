package aoc2023day10

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strings"
)

var pipes = map[string][][]int{
	"|": {{0, 1}, {0, -1}},
	"-": {{1, 0}, {-1, 0}},
	"L": {{1, 0}, {0, -1}},
	"J": {{-1, 0}, {0, -1}},
	"7": {{-1, 0}, {0, 1}},
	"F": {{0, 1}, {1, 0}},
	"S": {{1, 0}, {0, 1}, {-1, 0}, {0, -1}},
}

func MazeFarthestPoint() int {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	textInput, err := os.ReadFile(currentDir + "/2023/day10/day10-input.txt")

	if err != nil {
		log.Fatal(err)
	}

	buffer := textInput

	var maze [][]string

	startingRowCount := -1
	startingColumnCount := -1
	rowCount := 0
	for {
		advance, token, err := bufio.ScanLines(buffer, true)

		if err != nil {
			log.Fatal(err)
		}

		if advance == 0 {
			break
		}
		if advance <= len(buffer) {
			mazePipes := strings.Split(strings.TrimSpace(string(token)), "")

			colCount := slices.Index(mazePipes, "S")

			if colCount > -1 {
				startingRowCount = rowCount
				startingColumnCount = colCount
			}
			rowCount += 1

			maze = append(maze, mazePipes)

			buffer = buffer[advance:]
		}
	}

	farthestPointSteps := bfs(maze, []int{startingRowCount, startingColumnCount})

	return farthestPointSteps
}

func bfs(maze [][]string, startIndices []int) int {
	farthestPoint := 0
	visited := make([][]bool, len(maze))
	for i := range visited {
		visited[i] = make([]bool, len(maze[i]))
	}

	visited[startIndices[0]][startIndices[1]] = true

	queue := [][]int{}
	queue = append(queue, startIndices)
	queue = append(queue, nil) // Level 0

	for len(queue) != 0 {
		currentIndices := queue[0]
		queue = queue[1:]

		if currentIndices == nil {
			queue = append(queue, nil)

			if queue[0] == nil {
				break
			}

			farthestPoint += 1
			continue
		}

		row := currentIndices[0]
		col := currentIndices[1]
		currentNode := maze[row][col]
		neighbours := pipes[currentNode]

		for _, v := range neighbours {
			nextRowOffset := v[1]
			nextColOffset := v[0]

			nextRow := nextRowOffset + row
			nextCol := nextColOffset + col
			nextNode := maze[nextRow][nextCol]

			nextNeighbours, ok := pipes[nextNode]

			if !ok {
				continue
			}
			isValid := slices.IndexFunc(nextNeighbours, func(elementIdxs []int) bool {
				nextElRowOffset := elementIdxs[1]
				nextElColOffset := elementIdxs[0]

				if row == nextRow+nextElRowOffset && col == nextCol+nextElColOffset {
					return true
				}

				return false
			})

			if !visited[nextRow][nextCol] && isValid > -1 {
				queue = append(queue, []int{nextRow, nextCol})
			}
		}

		visited[row][col] = true
	}

	return farthestPoint
}
