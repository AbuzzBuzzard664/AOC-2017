package Day_05

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	fmt.Println("Part 1: ", Part1())
	fmt.Println("Part 2: ", Part2())
}

func Part1() int {
	steps := get_input()
	res := 0
	index := 0
	for index < len(steps) {
		temp := steps[index]
		steps[index]++
		index += temp
		res++
	}
	return res
}

func Part2() int {
	steps := get_input()
	res := 0
	index := 0
	for index < len(steps) {
		temp := steps[index]
		if temp >= 3 {
			steps[index]--
		} else {
			steps[index]++
		}
		index += temp
		res++
	}
	return res
}

func get_input() []int {
	input, err := os.ReadFile("./Day_05/Day_05.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return []int{}
	}
	ints := []int{}
	for _, val := range strings.Split(string(input), "\n") {
		val = strings.Trim(val, "\r")
		if val != "" {
			i, err := strconv.Atoi(val)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				continue
			}
			ints = append(ints, i)
		}
	}
	return ints
}
