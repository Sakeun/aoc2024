package day5

import (
	"os"
	"reflect"
	"slices"
	"strconv"
	"strings"
)

func parseDay5() ([]string, []string) {
	content, _ := os.ReadFile("week1/day5/input.txt")
	inputArr := strings.Split(string(content), "\n")

	var ruleArr []string
	var pagesArr []string

	for _, val := range inputArr {
		if strings.Contains(val, "|") {
			ruleArr = append(ruleArr, val)
		}
		if strings.Contains(val, ",") {
			pagesArr = append(pagesArr, val)
		}
	}

	return ruleArr, pagesArr
}

func createRulesMap(arr []string) map[string][]string {
	ruleMap := make(map[string][]string)
	for _, val := range arr {
		currRule := strings.Split(val, "|")

		page := currRule[0]
		rule := currRule[1]

		ruleMap[page] = append(ruleMap[page], rule)
	}

	return ruleMap
}

func checkPreviousPages(pages []string, ruleMap map[string][]string, page string, i int) bool {
	validPage := true

	for j := i - 1; j >= 0; j-- {
		prevPage := pages[j]

		if slices.Contains(ruleMap[page], prevPage) {
			validPage = false
		}
	}

	return validPage
}

func Part1() int {
	ruleArr, pagesArr := parseDay5()

	ruleMap := createRulesMap(ruleArr)
	validSum := 0

	for _, pageBook := range pagesArr {
		pages := strings.Split(pageBook, ",")
		isValid := true

		for i, page := range pages {
			if i == 0 {
				continue
			}

			if !checkPreviousPages(pages, ruleMap, page, i) {
				isValid = false
				break
			}
		}

		if isValid {
			middlePage := pages[len(pages)/2]

			middlePageNum, _ := strconv.Atoi(middlePage)

			validSum += middlePageNum
		}
	}

	return validSum
}

func checkAndSwapPages(pages []string, ruleMap map[string][]string, page string, i int) ([]string, bool) {
	hasSwapped := false
	for j := i - 1; j >= 0; j-- {
		prevPage := pages[j]

		if slices.Contains(ruleMap[page], prevPage) {
			swapF := reflect.Swapper(pages)
			swapF(i, j)
			i = j
			hasSwapped = true
		}
	}

	return pages, hasSwapped
}

func Part2() int {
	ruleArr, pagesArr := parseDay5()

	ruleMap := createRulesMap(ruleArr)
	validSum := 0

	for _, pageBook := range pagesArr {
		pages := strings.Split(pageBook, ",")
		var hasSwapped bool

		for i, page := range pages {
			if i == 0 {
				continue
			}

			if !hasSwapped {
				pages, hasSwapped = checkAndSwapPages(pages, ruleMap, page, i)
			} else {
				pages, _ = checkAndSwapPages(pages, ruleMap, page, i)
			}
		}

		if hasSwapped {
			middlePage := pages[len(pages)/2]

			middlePageNum, _ := strconv.Atoi(middlePage)

			validSum += middlePageNum
		}
	}

	return validSum
}
