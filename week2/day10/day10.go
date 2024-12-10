package day10

import (
	"os"
	"strconv"
	"strings"
)

const input = "89010123\n78121874\n87430965\n96549874\n45678903\n32019012\n01329801\n10456732\n"

func parseInput() [][]int {
	content, _ := os.ReadFile("week2/day10/input.txt")
	inputArr := strings.Split(string(content), "\n")

	inputArr = inputArr[:len(inputArr)-1]

	numArr := make([][]int, len(inputArr))

	for i, s := range inputArr {
		numArr[i] = make([]int, len(s))
		for j, val := range s {
			numArr[i][j], _ = strconv.Atoi(string(val))
		}
	}

	return numArr
}

func validatePath(arr [][]int, x int, y int, isPartOne bool) int {
	if arr[x][y] == 9 {
		if isPartOne {
			arr[x][y] = 'X'
		}
		return 1
	}

	totalAmount := 0
	currNum := arr[x][y]

	if y+1 < len(arr[x]) && arr[x][y+1]-currNum == 1 {
		totalAmount += validatePath(arr, x, y+1, isPartOne)
	}

	if y-1 >= 0 && arr[x][y-1]-currNum == 1 {
		totalAmount += validatePath(arr, x, y-1, isPartOne)
	}

	if x+1 < len(arr) && arr[x+1][y]-currNum == 1 {
		totalAmount += validatePath(arr, x+1, y, isPartOne)
	}

	if x-1 >= 0 && arr[x-1][y]-currNum == 1 {
		totalAmount += validatePath(arr, x-1, y, isPartOne)
	}

	return totalAmount
}

func Part1() int {
	arr := parseInput()
	totalAmount := 0

	for x, _ := range arr {
		for y, val := range arr[x] {
			if val == 0 {
				arr2 := parseInput()
				totalAmount += validatePath(arr2, x, y, true)
			}
		}
	}

	return totalAmount
}

func Part2() int {
	arr := parseInput()
	totalAmount := 0

	for x, _ := range arr {
		for y, val := range arr[x] {
			if val == 0 {
				arr2 := parseInput()
				totalAmount += validatePath(arr2, x, y, false)
			}
		}
	}

	return totalAmount
}
