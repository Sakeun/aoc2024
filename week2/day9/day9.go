package day9

import (
	"github.com/Sakeun/aoc2024/datastructs"
	"os"
	"strconv"
)

const input = "2333133121414131402"

func Part1() int {
	deq := datastructs.Deque[string]{}
	content, _ := os.ReadFile("week2/day9/input.txt")
	addNum := true
	i := 0

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
	i = 0
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
