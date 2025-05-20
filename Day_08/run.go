package Day_08

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	ref, compare, opref, op string
	opval, val              int
}

func Run() {
	fmt.Println("Part 1: ", part_1())
	fmt.Println("Part 2: ", Part_2())
}

func part_1() int {
	registers, instructions := get_input()
	res := 0
	for _, instruction := range instructions {
		isTrue := false
		switch instruction.compare {
		case ">":
			if registers[instruction.opref] > instruction.opval {
				isTrue = true
			}
		case "<":
			if registers[instruction.opref] < instruction.opval {
				isTrue = true
			}
		case ">=":
			if registers[instruction.opref] >= instruction.opval {
				isTrue = true
			}
		case "<=":
			if registers[instruction.opref] <= instruction.opval {
				isTrue = true
			}
		case "!=":
			if registers[instruction.opref] != instruction.opval {
				isTrue = true
			}
		case "==":
			if registers[instruction.opref] == instruction.opval {
				isTrue = true
			}
		}
		if isTrue {
			if instruction.op == "inc" {
				registers[instruction.ref] += instruction.val
			} else {
				registers[instruction.ref] -= instruction.val
			}
		}
	}
	for _, val := range registers {
		if val > res {
			res = val
		}
	}
	return res
}

func Part_2() int {
	registers, instructions := get_input()
	res := 0
	for _, instruction := range instructions {
		isTrue := false
		switch instruction.compare {
		case ">":
			if registers[instruction.opref] > instruction.opval {
				isTrue = true
			}
		case "<":
			if registers[instruction.opref] < instruction.opval {
				isTrue = true
			}
		case ">=":
			if registers[instruction.opref] >= instruction.opval {
				isTrue = true
			}
		case "<=":
			if registers[instruction.opref] <= instruction.opval {
				isTrue = true
			}
		case "!=":
			if registers[instruction.opref] != instruction.opval {
				isTrue = true
			}
		case "==":
			if registers[instruction.opref] == instruction.opval {
				isTrue = true
			}
		}
		if isTrue {
			if instruction.op == "inc" {
				registers[instruction.ref] += instruction.val
			} else {
				registers[instruction.ref] -= instruction.val
			}
		}
		for _, val := range registers {
			if val > res {
				res = val
			}
		}
	}
	return res
}

func get_input() (map[string]int, []Instruction) {
	registers := map[string]int{}
	instructions := []Instruction{}
	input, _ := os.ReadFile("./Day_08/Day_08.txt")
	for _, line := range strings.Split(string(input), "\n") {
		if strings.TrimSpace(line) == "" {
			continue // Skip empty lines
		}
		parts := strings.Split(line, " ")
		var temp Instruction
		registers[parts[0]] = 0
		temp.ref = parts[0]
		temp.op = parts[1]
		num, _ := strconv.Atoi(strings.TrimSpace(parts[2]))
		temp.val = num
		registers[parts[4]] = 0
		temp.opref = parts[4]
		temp.compare = parts[5]
		num, _ = strconv.Atoi(strings.TrimSpace(parts[6]))
		temp.opval = num
		instructions = append(instructions, temp)
	}
	return registers, instructions
}
