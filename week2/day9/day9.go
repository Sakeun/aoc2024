package day9

import (
	"github.com/Sakeun/aoc2024/datastructs"
	"os"
	"strconv"
)

const input = "2333133121414131402"

func createCodeDeque(content string) datastructs.Deque[string] {
	addNum := true
	i := 0

	deq := datastructs.Deque[string]{}
	for _, val := range string(content) {
		amount, _ := strconv.Atoi(string(val))
		if addNum {
			for j := 0; j < amount; j++ {
				deq.PushBack(strconv.Itoa(i))
			}
			i++
			addNum = false
		} else {
			for j := 0; j < amount; j++ {
				deq.PushBack(".")
			}
			addNum = true
		}
	}

	return deq
}

func Part1() int {
	content, _ := os.ReadFile("week2/day9/input.txt")

	deq := createCodeDeque(string(content))

	x := deq.Size() - 1
	for j := 0; j < deq.Size(); j++ {
		if j >= x {
			break
		}

		if deq[j] != "." {
			continue
		}

		for deq[x] == "." {
			x--
		}

		deq[j] = deq[x]
		deq[x] = "."
	}

	totalSum := 0
	i := 0
	for deq.Size() != 0 {
		valA, _ := deq.PopFront()
		if valA == "." {
			continue
		}

		num, _ := strconv.Atoi(valA)
		totalSum += num * i
		i++
	}

	return totalSum
}

type vals struct {
	num     string
	indices datastructs.Deque[int]
}

func Part2() int {
	content, _ := os.ReadFile("week2/day9/input.txt")
	deq := createCodeDeque(string(content))
	newDeq := createCodeDeque(string(content))

	for deq.Size() != 0 {
		valueToGrab := deq.PeekBack()
		if valueToGrab == "." {
			_, _ = deq.PopBack()
			continue
		}
		nums := vals{num: valueToGrab, indices: make(datastructs.Deque[int], 0)}
		for i := len(deq) - 1; deq.PeekBack() == valueToGrab; i-- {
			if i < 0 {
				break
			}
			_, _ = deq.PopBack()
			nums.indices.PushBack(i)
		}

		if deq.Size() == 0 {
			break
		}

		matchingPattern := make([]int, 0)

		for i := 0; i < deq.Size(); i++ {
			if newDeq[i] == "." {
				if len(matchingPattern) == 0 {
					matchingPattern = append(matchingPattern, i)
				} else {
					if i-matchingPattern[len(matchingPattern)-1] >= 2 {
						matchingPattern = make([]int, 0)
					}

					matchingPattern = append(matchingPattern, i)
				}

				if len(matchingPattern) == nums.indices.Size() {
					break
				}
			}
		}

		if len(matchingPattern) != nums.indices.Size() {
			continue
		}

		for _, val := range matchingPattern {
			newDeq[val] = nums.num
		}

		for _, val := range nums.indices {
			newDeq[val] = "."
		}

	}

	totalSum := 0
	i := 0
	for newDeq.Size() != 0 {
		valA, _ := newDeq.PopFront()
		if valA == "." {
			i++
			continue
		}

		num, _ := strconv.Atoi(valA)
		totalSum += num * i
		i++
	}

	return totalSum
}
