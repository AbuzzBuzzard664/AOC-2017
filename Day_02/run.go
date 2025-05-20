package Day_02

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	fmt.Println("Part 1: ", part_1())
	fmt.Println("Part 2: ", part_2())
}

func part_1() int {
	lines := get_input()
	res := 0
	for _, line := range lines {
		res += get_max(line) - get_min(line)
	}
	return res
}

func get_max(line []int) int {
	max := 0
	for _, i := range line {
		if i > max {
			max = i
		}
	}
	return max
}

func get_min(line []int) int {
	min := 100000000
	for _, i := range line {
		if i < min {
			min = i
		}
	}
	return min
}

func part_2() int {
	lines := get_input()
	res := 0
	for _, nums := range lines {
		found := false
		for i, num1 := range nums {
			for j, num2 := range nums {
				if i != j {
					if num2 % num1 == 0 {
						res += num2 / num1
						found = true
						break
					}
				}
			}
			if found {
				break
			}
		}
	}
	return res
}

func get_input() [][]int {
	lines := [][]int{}
	input, _ := os.ReadFile("./Day_02/Day_02.txt")
	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			continue
		}
		temp := []int{}
		myNum := ""
		for _, c := range line {
			if strings.Contains("0123456789", string(c)) {
				myNum += string(c)
			} else {
				if myNum != "" {
					if num, err := strconv.Atoi(myNum); err == nil {
						temp = append(temp, num)
					}
					myNum = ""
				}
			}
		}
		if myNum != "" {
			if num, err := strconv.Atoi(myNum); err == nil {
				temp = append(temp, num)
			}
		}
		lines = append(lines, temp)
	}
	return lines
}
