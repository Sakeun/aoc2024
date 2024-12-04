package day2

import (
	"os"
	"strconv"
	"strings"
)

func parseInput() [][]int {
	var rows [][]int
	var row []int

	var currNumBuilder strings.Builder

	content, _ := os.ReadFile("week1/day2/day2.txt")

	for _, val := range content {
		if val == 10 {
			num, _ := strconv.Atoi(currNumBuilder.String())
			row = append(row, num)

			rows = append(rows, row)

			currNumBuilder.Reset()
			row = []int{}
			continue
		}

		if val == 32 {
			num, _ := strconv.Atoi(currNumBuilder.String())
			row = append(row, num)
			currNumBuilder.Reset()
			continue
		}

		currNumBuilder.WriteString(string(val))
	}

	num, _ := strconv.Atoi(currNumBuilder.String())
	row = append(row, num)

	rows = append(rows, row)

	currNumBuilder.Reset()
	row = []int{}

	return rows
}

func checkSafe(isInc bool, row []int, i int) bool {
	if row[i+1] == row[i] {
		return false
	}

	if isInc {
		return row[i+1]-row[i] <= 3
	} else {
		return row[i]-row[i+1] <= 3
	}

}

func checkIncSafe(isInc bool, row []int, i int) bool {
	if isInc && row[i] > row[i+1] {
		return false
	} else if !isInc && row[i] < row[i+1] {
		return false
	}
	return true
}

func PuzzleOne() int {
	rows := parseInput()
	amountSafe := 0

	for _, row := range rows {
		if checkSpecificArrLoop(row) {
			amountSafe++
		}
	}

	return amountSafe
}

func remove(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}

func checkSpecificArrLoop(row []int) bool {
	isSafe := false
	isInc := false

	for i, num := range row {
		if i+1 >= len(row) {
			break
		}

		if i == 0 {
			isInc = num < row[i+1]
		}

		isSafe = checkSafe(isInc, row, i) && checkIncSafe(isInc, row, i)

		if !isSafe {
			break
		}
	}

	return isSafe
}

func PuzzleTwo() int {
	rows := parseInput()
	amountSafe := 0

	for _, row := range rows {
		if checkSpecificArrLoop(row) {
			amountSafe++
			continue
		}

		isSafe := false

		for i := 0; i < len(row); i++ {
			arr := make([]int, len(row))
			copy(arr, row)
			arr = remove(arr, i)

			isSafe = checkSpecificArrLoop(arr)

			if isSafe {
				amountSafe++
				break
			}
		}
	}

	return amountSafe
}
