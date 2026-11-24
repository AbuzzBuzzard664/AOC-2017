package Day_23

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	lines := get_input()
	fmt.Println("Part 1:", part_1(lines))
	fmt.Println("Part 2:", part_2(lines))
}

func part_1(lines []string) int {
	registers := map[string]int{"a": 0, "b": 0, "c": 0, "d": 0, "e": 0, "f": 0, "g": 0, "h": 0}
	res := 0
	idx := 0
	for idx >= 0 && idx < len(lines) {
		fields := strings.Fields(lines[idx])
		if fields[0] == "mul" {
			res++
		}
		idx += 1 + read_line(lines[idx], registers)
	}
	return res
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n%2 == 0 {
		return n == 2
	}
	for d := 3; d*d <= n; d += 2 {
		if n%d == 0 {
			return false
		}
	}
	return true
}

func part_2(lines []string) int {
	parts := strings.Fields(lines[0])
	base, _ := strconv.Atoi(parts[2])

	b := base*100 + 100000
	c := b + 17000
	h := 0

	for ; b <= c; b += 17 {
		if !isPrime(b) {
			h++
		}
	}
	return h
}
func read_line(line string, registers map[string]int) int {
	splits := strings.Fields(line)
	switch splits[0] {
	case "set":
		registers[splits[1]] = getValue(splits[2], registers)
	case "sub":
		registers[splits[1]] -= getValue(splits[2], registers)
	case "mul":
		registers[splits[1]] *= getValue(splits[2], registers)
	case "jnz":
		if getValue(splits[1], registers) != 0 {
			return getValue(splits[2], registers) - 1
		}
	}
	return 0
}

func getValue(operand string, registers map[string]int) int {
	num, err := strconv.Atoi(operand)
	if err == nil {
		return num
	}
	return registers[operand]
}

func get_input() []string {
	data, _ := os.ReadFile("./Day_23/Day_23.txt")
	return strings.Split(strings.TrimSpace(string(data)), "\n")
}
