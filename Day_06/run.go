package Day_06

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Run() {
	fmt.Println("Part 1:", part_1())
	fmt.Println("Part 2:", part_2())
}

func part_1() int {
	bank := get_input()
	return len(debug(bank)) - 1
}

func debug(bank []int) [][]int {
	seen := [][]int{}
	for {
		seen = append(seen, slices.Clone(bank))
		max := 0
		for i := 1; i < len(bank); i++ {
			if bank[i] > bank[max] {
				max = i
			}
		}
		redist := bank[max]
		bank[max] = 0
		for i := (max + 1) % len(bank); redist > 0; i = (i + 1) % len(bank) {
			bank[i]++
			redist--
		}
		for _, val := range seen {
			if slices.Equal(val, bank) {
				seen = append(seen, slices.Clone(bank))
				return seen
			}
		}
	}
}

func part_2() int {
	bank := get_input()
	seen := debug(bank)
	for i, s := range seen {
		if slices.Equal(s, seen[len(seen)-1]) {
			return len(seen) - i - 1
		}
	}
	return -1
}

func get_input() []int {
	input, _ := os.ReadFile("./Day_06/Day_06.txt")
	banks := []int{}
	for _, val := range strings.Split(string(input), " ") {
		val = strings.TrimSpace(val)
		if val != "" {
			num, _ := strconv.Atoi(val)
			banks = append(banks, num)
		}
	}
	return banks
}
