package day4

import (
	"os"
	"strings"
)

func parseDay4() []string {
	content, _ := os.ReadFile("week1/day4/day4.txt")
	inputArr := strings.Split(string(content), "\n")

	return inputArr
}

func getNextLetter(currLetter string) string {
	switch currLetter {
	case "X":
		return "M"
	case "M":
		return "A"
	case "A":
		return "S"
	default:
		return ""
	}
}

func checkAround(letter string, col int, row int, arr []string, dir string) bool {
	if row < 0 || row >= len(arr) || col < 0 || col >= len(arr[row]) {
		return false
	}

	if string(arr[row][col]) != letter {
		return false
	}

	if letter == "S" {
		return true
	}

	switch dir {
	case "T":
		return checkAround(getNextLetter(letter), col, row-1, arr, dir)
	case "B":
		return checkAround(getNextLetter(letter), col, row+1, arr, dir)
	case "TL":
		return checkAround(getNextLetter(letter), col-1, row-1, arr, dir)
	case "TR":
		return checkAround(getNextLetter(letter), col+1, row-1, arr, dir)
	case "BL":
		return checkAround(getNextLetter(letter), col-1, row+1, arr, dir)
	case "BR":
		return checkAround(getNextLetter(letter), col+1, row+1, arr, dir)
	case "L":
		return checkAround(getNextLetter(letter), col-1, row, arr, dir)
	case "R":
		return checkAround(getNextLetter(letter), col+1, row, arr, dir)
	default:
		return false
	}
}

func Day4Part1() int {
	arr := parseDay4()
	initialLetter := "X"
	totalOcc := 0

	for i, row := range arr {
		for j, letter := range row {
			if string(letter) == "X" {
				if checkAround(getNextLetter(initialLetter), j, i+1, arr, "B") {
					totalOcc++
				}
				if checkAround(getNextLetter(initialLetter), j, i-1, arr, "T") {
					totalOcc++
				}
				if checkAround(getNextLetter(initialLetter), j-1, i-1, arr, "TL") {
					totalOcc++
				}
				if checkAround(getNextLetter(initialLetter), j+1, i-1, arr, "TR") {
					totalOcc++
				}
				if checkAround(getNextLetter(initialLetter), j-1, i+1, arr, "BL") {
					totalOcc++
				}
				if checkAround(getNextLetter(initialLetter), j+1, i+1, arr, "BR") {
					totalOcc++
				}
				if checkAround(getNextLetter(initialLetter), j-1, i, arr, "L") {
					totalOcc++
				}
				if checkAround(getNextLetter(initialLetter), j+1, i, arr, "R") {
					totalOcc++
				}
			}
		}
	}

	return totalOcc
}

func checkXmas(arr []string, row int, col int) bool {
	if row == 0 || col == 0 || row+1 >= len(arr) || col+1 >= len(arr[row]) || len(arr[row+1]) == 0 {
		return false
	}

	validXmas := false
	topLeftChar := string(arr[row-1][col-1])
	botRightChar := string(arr[row+1][col+1])
	topRightChar := string(arr[row-1][col+1])
	botLeftChar := string(arr[row+1][col-1])

	if ((botLeftChar == "M" && topRightChar == "S") || (botLeftChar == "S" && topRightChar == "M")) &&
		((topLeftChar == "M" && botRightChar == "S") || (topLeftChar == "S" && botRightChar == "M")) {
		validXmas = true
	}

	return validXmas
}

func Day4Part2() int {
	arr := parseDay4()
	initialLetter := "A"
	totalOcc := 0

	for i, row := range arr {
		for j, letter := range row {
			if string(letter) == initialLetter {
				if checkXmas(arr, i, j) {
					totalOcc++
				}
			}
		}
	}

	return totalOcc
}
