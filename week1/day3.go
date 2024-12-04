package week1

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseDay3() []string {
	r, _ := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)

	content, _ := os.ReadFile("week1/inputs/day3.txt")

	return r.FindAllString(string(content), -1)
}

func parseDay3Pt2() []string {
	r, _ := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)

	content, _ := os.ReadFile("week1/inputs/day3.txt")

	return r.FindAllString(string(content), -1)
}

func Day3Part1() int {
	regexResult := parseDay3()
	totalSum := 0

	for _, val := range regexResult {
		newStr := strings.Replace(val, "mul(", "", -1)
		newStr = strings.Replace(newStr, ")", "", -1)

		nums := strings.Split(newStr, ",")

		numA, _ := strconv.Atoi(nums[0])
		numB, _ := strconv.Atoi(nums[1])

		totalSum += numA * numB
	}

	return totalSum
}

func Day3Part2() int {
	regexResult := parseDay3Pt2()
	totalSum := 0
	mulEnabled := true

	for _, val := range regexResult {
		if val == "do()" {
			mulEnabled = true
			continue
		}
		if !mulEnabled {
			continue
		}
		if val == "don't()" {
			mulEnabled = false
			continue
		}

		newStr := strings.Replace(val, "mul(", "", -1)
		newStr = strings.Replace(newStr, ")", "", -1)

		nums := strings.Split(newStr, ",")

		numA, _ := strconv.Atoi(nums[0])
		numB, _ := strconv.Atoi(nums[1])

		totalSum += numA * numB
	}

	return totalSum
}
