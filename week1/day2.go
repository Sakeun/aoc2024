package week1

import (
	"strconv"
	"strings"
)

const input = "7 6 4 2 1\n1 2 7 8 9\n9 7 6 2 1\n1 3 2 4 5\n8 6 4 4 1\n1 3 6 7 9\n5 9 4 3 2"

func parseInputDay2() [][]int {
	rows := [][]int{}
	row := []int{}

	var currNumBuilder strings.Builder

	//content, _ := os.ReadFile("week1/inputs/day2.txt")

	for _, val := range input {
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

func DayTwoPuzzleOne() int {
	rows := parseInputDay2()
	amountSafe := 0

	for _, row := range rows {
		isInc := false
		isSafe := false

		for i, num := range row {
			if i+1 >= len(row) {
				break
			}

			if i == 0 {
				isInc = num < row[i+1]
			} else if isInc && num > row[i+1] {
				isSafe = false
				break
			} else if !isInc && num < row[i+1] {
				isSafe = false
				break
			}

			isSafe = checkSafe(isInc, row, i)

			if !isSafe {
				break
			}
		}

		if isSafe {
			amountSafe++
		}
	}

	return amountSafe
}

func remove(slice []int, i int) []int {
	return append(slice[:i+1], slice[i+2:]...)
}

func DayTwoPuzzleTwo() int {

	rows := parseInputDay2()
	amountSafe := 0

	for _, row := range rows {
		isInc := false
		isSafe := false
		passUsed := false

		for i := 0; i <= len(row); {
			if i+1 >= len(row) {
				break
			}

			if i != 0 && !passUsed && (!checkIncSafe(isInc, row, i) || !checkSafe(isInc, row, i)) {
				passUsed = true
				row = remove(row, i)
				continue
			}

			if i+1 >= len(row) {
				break
			}

			if i == 0 {
				isInc = row[i] < row[i+1]
			}

			isSafe = checkSafe(isInc, row, i) || checkIncSafe(isSafe, row, i)

			if !isSafe {
				break
			}
			i++
		}

		if isSafe {
			amountSafe++
		}
	}

	return amountSafe
}
