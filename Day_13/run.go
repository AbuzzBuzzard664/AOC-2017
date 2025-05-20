package Day_13

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	scanners := get_input()
	fmt.Println("Part 1:", part_1(scanners))
	fmt.Println("Part 2:", part_2(scanners))
}

func part_1(scanners []int) int {
	severity := 0
	for time, length := range scanners {
		if length > 0 && time%(2*(length-1)) == 0 {
			severity += time * length
		}
	}
	return severity
}

func part_2(scanners []int) int {
	delay := 0
	for {
		caught := false
		for time, length := range scanners {
			if length > 0 && (time+delay)%(2*(length-1)) == 0 {
				caught = true
				break
			}
		}
		if !caught {
			return delay
		}
		delay++
	}
}

func get_input() []int {
	input, _ := os.ReadFile("./Day_13/Day_13.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	scanners := []int{}
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			continue
		}
		pos, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		length, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		for len(scanners) <= pos {
			scanners = append(scanners, 0)
		}
		scanners[pos] = length
	}
	return scanners
}
