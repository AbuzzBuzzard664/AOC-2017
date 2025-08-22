package Day_15

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Run() {
	a, b := get_input()
	fmt.Println("Part 1:", part_1(a, b))
	fmt.Println("Part 2:", part_2(a, b))
}

func part_1(a, b int64) int {
	count := 0
	for i := 0; i < 40000000; i++ {
		a = (a * 16807) % 2147483647
		b = (b * 48271) % 2147483647

		if (a & 65535) == (b & 65535) {
			count++
		}
	}
	return count
}

func part_2(a, b int64) int {
	count := 0
	for i := 0; i < 5000000; i++ {
		a = genA(a)
		b = genB(b)
		if (a & 65535) == (b & 65535) {
			count++
		}
	}
	return count
}

func genA(a int64) int64 {
	for {
		a = (a * 16807) % 2147483647
		if a%4 == 0 {
			return a
		}
	}
}

func genB(b int64) int64 {
	for {
		b = (b * 48271) % 2147483647
		if b%8 == 0 {
			return b
		}
	}
}

func get_input() (int64, int64) {
	data, _ := os.ReadFile("./Day_15/Day_15.txt")
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	re := regexp.MustCompile(`\d+`)
	genA := 0
	genB := 0
	if len(lines) >= 2 {
		if matches := re.FindString(lines[0]); matches != "" {
			genA, _ = strconv.Atoi(matches)
		}
		if matches := re.FindString(lines[1]); matches != "" {
			genB, _ = strconv.Atoi(matches)
		}
	}
	return int64(genA), int64(genB)
}
