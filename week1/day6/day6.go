package day6

import (
	"os"
	"strings"
)

type Direction int

const (
	Up      Direction = iota
	Down    Direction = iota
	Left    Direction = iota
	Right   Direction = iota
	Invalid Direction = iota
)

func parseInput() []string {
	content, _ := os.ReadFile("week1/day6/input.txt")
	inputArr := strings.Split(string(content), "\n")

	return inputArr
}

func getInitialDirection(guard string) Direction {
	switch guard {
	case "^":
		return Up
	case "v":
		return Down
	case "<":
		return Left
	case ">":
		return Right
	default:
		return Invalid
	}
}

func findInitialPosition(rows []string) (row int, col int, direction Direction) {
	direction = Invalid

	for i, currentRow := range rows {
		for j, currentColumn := range currentRow {
			direction = getInitialDirection(string(currentColumn))

			if direction != Invalid {
				col = j
				row = i
				return
			}
		}
	}

	return
}

func replaceField(rows []string, col int, row int) {
	out := []rune(rows[row])
	out[col] = 'X'
	rows[row] = string(out)
}

func checkUpwards(rows []string, col int, row int, direction Direction) int {
	hasVisited := 0

	if string(rows[row][col]) != "X" {
		hasVisited++
		replaceField(rows, col, row)
	}

	if row-1 < 0 {
		return hasVisited
	}

	nextRow := row - 1

	if string(rows[nextRow][col]) == "#" {
		return checkRight(rows, col+1, row, Right) + hasVisited
	}

	return checkUpwards(rows, col, nextRow, direction) + hasVisited
}

func checkDownwards(rows []string, col int, row int, direction Direction) int {
	hasVisited := 0

	if string(rows[row][col]) != "X" {
		hasVisited++
		replaceField(rows, col, row)
	}

	if row+1 == len(rows) {
		return hasVisited
	}

	nextRow := row + 1

	if string(rows[nextRow][col]) == "#" {
		return checkLeft(rows, col-1, row, Left) + hasVisited
	}

	return checkDownwards(rows, col, nextRow, direction) + hasVisited
}

func checkLeft(rows []string, col int, row int, direction Direction) int {
	hasVisited := 0

	if string(rows[row][col]) != "X" {
		hasVisited++
		replaceField(rows, col, row)
	}

	if col-1 < 0 {
		return hasVisited
	}

	nextCol := col - 1

	if string(rows[row][nextCol]) == "#" {
		return checkUpwards(rows, col, row-1, Up) + hasVisited
	}

	return checkLeft(rows, nextCol, row, direction) + hasVisited
}

func checkRight(rows []string, col int, row int, direction Direction) int {
	hasVisited := 0

	if string(rows[row][col]) != "X" {
		hasVisited++
		replaceField(rows, col, row)
	}

	if col+1 >= len(rows[row]) {
		return hasVisited
	}

	nextCol := col + 1

	if string(rows[row][nextCol]) == "#" {
		return checkDownwards(rows, col, row+1, Down) + hasVisited
	}

	return checkRight(rows, nextCol, row, direction) + hasVisited
}

func Part1() int {
	rows := parseInput()
	startingRow, startingCol, startingDirection := findInitialPosition(rows)

	totalFields := 0

	if startingDirection == Invalid {
		totalFields = 0
	}

	if startingDirection == Down {
		totalFields = checkDownwards(rows, startingCol, startingRow, startingDirection)
	}

	if startingDirection == Left {
		totalFields = checkLeft(rows, startingCol, startingRow, startingDirection)
	}

	if startingDirection == Right {
		totalFields = checkRight(rows, startingCol, startingRow, startingDirection)
	}

	totalFields = checkUpwards(rows, startingCol, startingRow, startingDirection)

	return totalFields
}
