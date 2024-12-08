package day8

import (
	"os"
	"strings"
)

type Vector struct {
	x int
	y int
}

func parseInput() []string {
	content, _ := os.ReadFile("week2/day8/input.txt")
	inputArr := strings.Split(string(content), "\n")

	inputArr = inputArr[:len(inputArr)-1]

	return inputArr
}

func mapDataPoints(input []string) map[string][]Vector {
	pointMap := make(map[string][]Vector, 0)

	for x, line := range input {
		for y, val := range line {
			if val != '.' {
				pointMap[string(val)] = append(pointMap[string(val)], Vector{x, y})
			}
		}
	}

	return pointMap
}

func contains(arr []Vector, vec Vector) bool {
	for _, a := range arr {
		if a == vec {
			return true
		}
	}

	return false
}

func (a *Vector) add(b Vector) Vector {
	return Vector{a.x + b.x, a.y + b.y}
}

func (a *Vector) sub(b Vector) Vector {
	return Vector{a.x - b.x, a.y - b.y}
}

func checkIsValid(vec Vector, maxX int, maxY int) bool {
	if vec.x < 0 || vec.x >= maxX || vec.y < 0 || vec.y >= maxY {
		return false
	}

	return true
}

func Part1() int {
	inputArr := parseInput()
	pointMap := mapDataPoints(inputArr)
	hashtagArr := make([]Vector, 0)

	for _, v := range pointMap {
		for j, val := range v {
			for i := 0; i < len(v); i++ {
				if i == j {
					continue
				}
				valB := v[i]

				diffA := val.sub(valB)
				diffB := valB.sub(val)

				newPosA := val.add(diffA)
				newPosB := valB.add(diffB)

				if checkIsValid(newPosA, len(inputArr), len(inputArr[0])) && !contains(hashtagArr, newPosA) {
					hashtagArr = append(hashtagArr, newPosA)
				}

				if checkIsValid(newPosB, len(inputArr), len(inputArr[0])) && !contains(hashtagArr, newPosB) {
					hashtagArr = append(hashtagArr, newPosB)
				}
			}
		}
	}

	return len(hashtagArr)
}

func Part2() int {
	inputArr := parseInput()
	pointMap := mapDataPoints(inputArr)
	hashtagArr := make([]Vector, 0)

	for _, v := range pointMap {
		for j, val := range v {
			for i := 0; i < len(v); i++ {
				if i == j {
					continue
				}

				valB := v[i]

				diffA := val.sub(valB)
				diffB := valB.sub(val)

				if !contains(hashtagArr, val) {
					hashtagArr = append(hashtagArr, val)
				}

				if !contains(hashtagArr, v[i]) {
					hashtagArr = append(hashtagArr, v[i])
				}

				newPos := val.add(diffA)

				for checkIsValid(newPos, len(inputArr), len(inputArr[0])) {
					if !contains(hashtagArr, newPos) {
						hashtagArr = append(hashtagArr, newPos)
					}
					newPos = newPos.add(diffA)
				}

				newPos = valB.add(diffB)

				for checkIsValid(newPos, len(inputArr), len(inputArr[0])) {
					if !contains(hashtagArr, newPos) {
						hashtagArr = append(hashtagArr, newPos)
					}
					newPos = newPos.add(diffB)
				}

			}
		}
	}

	return len(hashtagArr)
}
