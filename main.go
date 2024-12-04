package main

import (
	"fmt"
	"github.com/Sakeun/aoc2024/week1"
)

func main() {
	fmt.Println("Week 1:")
	fmt.Println("Day 1:")
	fmt.Printf("Puzzle one: %v \n", week1.PuzzleOne())
	fmt.Printf("Puzzle two: %v \n", week1.PuzzleTwo())
	fmt.Println()

	fmt.Println("Day 2:")
	fmt.Printf("Puzzle one: %v \n", week1.DayTwoPuzzleOne())
	fmt.Printf("Puzzle two: %v \n", week1.DayTwoPuzzleTwo())
	fmt.Println()

	fmt.Println("Day 3:")
	fmt.Printf("Puzzle one: %v \n", week1.Day3Part1())
	fmt.Printf("Puzzle two: %v \n", week1.Day3Part2())
	fmt.Println()

	fmt.Println("Day 4:")
	fmt.Printf("Puzzle one: %v \n", week1.Day4Part1())
	fmt.Printf("Puzzle two: %v \n", week1.Day4Part2())
}
