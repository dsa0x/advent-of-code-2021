package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day2(input []string) int {

	x := 0
	y := 0
	aim := 0
	for _, v := range input {
		res := strings.Split(v, " ")
		dir := res[0]
		num, _ := strconv.Atoi(res[1])
		if dir == "forward" {
			x += num
			y += num * aim
		} else if dir == "up" {
			// y += num;
			aim -= num
		} else if dir == "down" {
			// y -= num;
			aim += num
		}
	}

	return x * y
}

func main() {
	text, err := os.ReadFile("./../input/day2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	input := strings.Split(string(text), "\n")

	fmt.Println(Day2(input))
}
