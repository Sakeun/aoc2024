package main

import (
	"fmt"
	"github.com/Sakeun/aoc2024/week1/day1"
	"github.com/Sakeun/aoc2024/week1/day2"
	"github.com/Sakeun/aoc2024/week1/day3"
	"github.com/Sakeun/aoc2024/week1/day4"
)

func main() {
	fmt.Println("Week 1:")
	fmt.Println("Day 1:")
	fmt.Printf("Puzzle one: %v \n", day1.PuzzleOne())
	fmt.Printf("Puzzle two: %v \n", day1.PuzzleTwo())
	fmt.Println()

	fmt.Println("Day 2:")
	fmt.Printf("Puzzle one: %v \n", day2.PuzzleOne())
	fmt.Printf("Puzzle two: %v \n", day2.PuzzleTwo())
	fmt.Println()

	fmt.Println("Day 3:")
	fmt.Printf("Puzzle one: %v \n", day3.Day3Part1())
	fmt.Printf("Puzzle two: %v \n", day3.Day3Part2())
	fmt.Println()

	fmt.Println("Day 4:")
	fmt.Printf("Puzzle one: %v \n", day4.Day4Part1())
	fmt.Printf("Puzzle two: %v \n", day4.Day4Part2())
}
