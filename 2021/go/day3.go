package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1(input []string) int {

	binLength := len(input[0])
	gamma := ""
	epsilon := ""

	for i := 0; i < binLength; i++ {
		store := make(map[string]int)
		for _, v := range input {
			_, ok := store[string(v[i])]
			if ok {
				store[string(v[i])]++
			} else {
				store[string(v[i])] = 1
			}
		}
		if store["1"] > store["0"] {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}
	g, _ := strconv.ParseInt(gamma, 2, 64)
	e, _ := strconv.ParseInt(epsilon, 2, 64)

	return int(g * e)
}

func Part2(input []string) int {

	return calculate(input, 1, 0) * calculate(input, 0, 1)
}

func calculate(input []string, max, min int) int {
	val := 0
	i := 0
	for len(input) > 1 {
		store := make(map[string]int)
		for _, v := range input {
			_, ok := store[string(v[i])]
			if ok {
				store[string(v[i])]++
			} else {
				store[string(v[i])] = 1
			}
		}
		if store["1"] >= store["0"] {
			val = max
		} else {
			val = min
		}
		res := []string{}
		for _, v := range input {
			if string(v[i]) == strconv.Itoa(val) {
				res = append(res, v)
			}
		}
		input = res
		i++
	}
	ox, _ := strconv.ParseInt(input[0], 2, 64)

	return int(ox)
}

func main() {
	text, err := os.ReadFile("./../input/day3.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	input := strings.Split(string(text), "\n")

	fmt.Println(Part1(input))
	fmt.Println(Part2(input))
}
