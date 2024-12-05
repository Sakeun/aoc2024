package day1

import (
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseInput() ([]int, []int) {
	var row1 []int
	var row2 []int
	var currNumBuilder strings.Builder

	content, _ := os.ReadFile("week1/day1/day1.txt")

	for i, val := range content {
		if val <= 47 || val >= 58 {
			continue
		}

		currNumBuilder.WriteString(string(val))

		if i+1 == len(content) || (content[i+1] <= 47 || content[i+1] >= 58) {
			num, _ := strconv.Atoi(currNumBuilder.String())

			if len(row1) == len(row2) {
				row1 = append(row1, num)
			} else {
				row2 = append(row2, num)
			}

			currNumBuilder.Reset()
			continue
		}
	}

	sort.Ints(row1)
	sort.Ints(row2)
	return row1, row2
}

func Part1() int {
	row1, row2 := parseInput()
	sum := 0

	for i, num := range row1 {
		if num > row2[i] {
			sum += num - row2[i]
		} else {
			sum += row2[i] - num
		}
	}

	return sum
}

func Part2() int {
	row1, row2 := parseInput()
	row2Map := make(map[int]int)
	totalCost := 0

	for _, num := range row2 {
		row2Map[num]++
	}

	for _, num := range row1 {
		totalCost += num * row2Map[num]
	}

	return totalCost
}
