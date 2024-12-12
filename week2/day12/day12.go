package day12

import (
	"github.com/Sakeun/aoc2024/datastructs"
	"os"
	"strings"
)

type Vector struct {
	x int
	y int
}

const input = "RRRRIICCFF\nRRRRIICCCF\nVVRRRCCFFF\nVVRCCCJFFF\nVVVVCJJCFE\nVVIVCCJJEE\nVVIIICJJEE\nMIIIIIJJEE\nMIIISIJEEE\nMMMISSJEEE\n"

func parseInput() [][]rune {
	content, _ := os.ReadFile("week2/day12/input.txt")
	stringArr := strings.Split(string(content), "\n")
	stringArr = stringArr[:len(stringArr)-1]

	inputArr := make([][]rune, len(stringArr))

	for i, s := range stringArr {
		inputArr[i] = []rune(s)
	}

	return inputArr
}

func getAdj(loc Vector, arr [][]rune) []Vector {
	adj := make([]Vector, 0)
	currValue := arr[loc.y][loc.x]

	if loc.x > 0 && arr[loc.y][loc.x-1] == currValue {
		newLoc := Vector{loc.x - 1, loc.y}
		adj = append(adj, newLoc)
	}

	if loc.x < len(arr[0])-1 && arr[loc.y][loc.x+1] == currValue {
		newLoc := Vector{loc.x + 1, loc.y}
		adj = append(adj, newLoc)
	}

	if loc.y > 0 && arr[loc.y-1][loc.x] == currValue {
		newLoc := Vector{loc.x, loc.y - 1}
		adj = append(adj, newLoc)
	}

	if loc.y < len(arr)-1 && arr[loc.y+1][loc.x] == currValue {
		newLoc := Vector{loc.x, loc.y + 1}
		adj = append(adj, newLoc)
	}

	return adj
}

var allVisitedPoints = make(map[Vector]bool)

func BFS(x int, y int, arr [][]rune) int {
	result := 0
	var q datastructs.Queue[Vector]
	q.Enqueue(Vector{x, y})
	visitedPoints := make(map[Vector]bool)

	for !q.IsEmpty() {
		curr, _ := q.Dequeue()
		if visitedPoints[curr] {
			continue
		}
		visitedPoints[curr] = true
		allVisitedPoints[curr] = true
		adj := getAdj(curr, arr)

		result += 4 - len(adj)
		for _, v := range adj {
			q.Enqueue(v)
		}
	}

	return result * len(visitedPoints)
}

func Part1() int {
	inputArr := parseInput()
	res := 0

	for i, row := range inputArr {
		for j := range row {
			if allVisitedPoints[Vector{i, j}] {
				continue
			}
			res += BFS(i, j, inputArr)
		}
	}

	return res
}
