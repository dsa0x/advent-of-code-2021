package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1(input []int) int {

	max := input[0]
	count := 0
	for i := 1; i < len(input); i++ {
		num := input[i]
		if num > max {
			count++
		}
		max = num
	}
	return count
}

func Part2(input []int) int {
	startQueue := []int{}
	result := []int{}
	for i := 0; i < 3; i++ {
		startQueue = append(startQueue, input[i])
	}
	val := sum(startQueue)
	result = append(result, val)

	for i := 3; i < len(input); i++ {
		val = val - input[i-3] + input[i]
		// startQueue = append(startQueue[1:], input[i])
		result = append(result, val)
	}

	return Part1(result)

}

func sum(input []int) int {
	sum := 0
	for _, num := range input {
		sum += num
	}
	return sum
}

func main() {
	text, err := os.ReadFile("./../input/day1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	input := strings.Split(string(text), "\n")
	inputInt := []int{}

	for _, line := range input {
		num, _ := strconv.Atoi(line)
		inputInt = append(inputInt, num)
	}

	fmt.Println(Part1(inputInt))
	fmt.Println(Part2(inputInt))
}
