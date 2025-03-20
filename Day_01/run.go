package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Part 1: ", part_1())
	fmt.Println("Part 2: ", part_2())
}

func part_1() int {
	input := get_input()
	res := 0
	for i := range len(input) {
		if input[i] == input[(i+1)%len(input)] {
			num, _ := strconv.Atoi(string(input[i]))
			res += num
		}
	}
	return res
}

func part_2() int {
	input := get_input()
	res := 0
	step := len(input) / 2
	for i := range len(input) {
		if input[i] == input[(i+step)%len(input)] {
			num, _ := strconv.Atoi(string(input[i]))
			res += num
		}
	}
	return res
}

func get_input() string {
	input, _ := os.ReadFile("./Inputs/Day_01.txt")
	return string(input)
}
