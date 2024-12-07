package day7

import (
	"os"
	"strconv"
	"strings"
)

const input = "190: 10 19\n3267: 81 40 27\n83: 17 5\n156: 15 6\n7290: 6 8 6 15\n161011: 16 10 13\n192: 17 8 14\n21037: 9 7 18 13\n292: 11 6 16 20\n"

func parseInput() []string {
	content, _ := os.ReadFile("week1/day7/input.txt")
	inputArr := strings.Split(string(content), "\n")

	inputArr = inputArr[:len(inputArr)-1]

	return inputArr
}

func inputToMap(input []string) map[int][]int {
	inputMap := make(map[int][]int)
	for _, line := range input {
		allInput := strings.Split(line, ": ")

		solution := allInput[0]
		equation := strings.Split(allInput[1], " ")

		solutionNum, _ := strconv.Atoi(solution)

		for _, v := range equation {
			num, _ := strconv.Atoi(v)

			inputMap[solutionNum] = append(inputMap[solutionNum], num)
		}
	}

	return inputMap
}

func calculateNum(a int, b int, equationType rune) int {
	if equationType == '+' {
		return a + b
	}

	if equationType == '|' {
		num, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
		return num
	}

	return a * b
}

func traverseBinaryTree(currSum int, nums []int, solution int, eqType rune) bool {
	if len(nums) == 0 {
		if currSum == solution {
			return true
		}
		return false
	}

	currSum = calculateNum(currSum, nums[0], eqType)

	foundSolution := traverseBinaryTree(currSum, nums[1:], solution, '+') || traverseBinaryTree(currSum, nums[1:], solution, '*')

	if foundSolution {
		return true
	}

	return false
}

func Part1() int {
	totalSum := 0
	inputArr := parseInput()
	inputMap := inputToMap(inputArr)

	for k, v := range inputMap {
		if traverseBinaryTree(v[0], v[1:], k, '+') || traverseBinaryTree(v[0], v[1:], k, '*') {
			totalSum += k
		}
	}

	return totalSum
}

func traverseTree(currSum int, nums []int, solution int, eqType rune) bool {
	if len(nums) == 0 {
		if currSum == solution {
			return true
		}
		return false
	}

	currSum = calculateNum(currSum, nums[0], eqType)

	foundSolution := traverseTree(currSum, nums[1:], solution, '+') || traverseTree(currSum, nums[1:], solution, '|') || traverseTree(currSum, nums[1:], solution, '*')

	if foundSolution {
		return true
	}

	return false
}

func Part2() int {
	totalSum := 0
	inputArr := parseInput()
	inputMap := inputToMap(inputArr)

	for k, v := range inputMap {
		if traverseTree(v[0], v[1:], k, '+') || traverseTree(v[0], v[1:], k, '*') || traverseTree(v[0], v[1:], k, '|') {
			totalSum += k
		}
	}

	return totalSum
}
