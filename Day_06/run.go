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
	banks := get_input()
	seen := [][]int{}
	for {
		seen = append(seen, slices.Clone(banks))
		max := 0
		for i := 1; i < len(banks); i++ {
			if banks[i] > banks[max] {
				max = i
			}
		}
		redist := banks[max]
		banks[max] = 0
		for i := (max + 1) % (len(banks) - 1); redist > 0; i = (i + 1) % (len(banks) - 1) {
			banks[i]++
			redist--
		}
		banks[max] = redist % (len(banks) - 1)
		fmt.Println(banks)
		for _, val := range seen {
			if slices.Equal(val, banks) {
				return len(seen) - 1
			}
		}
	}
}

func part_2() int {
	return 0
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
