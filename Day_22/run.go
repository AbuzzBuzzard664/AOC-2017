package Day_22

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

func Run() {
	input, start := get_input()
	fmt.Println("Part 1:", part_1(input, start))
	input, start = get_input()
	fmt.Println("Part 2:", part_2(input, start))
}

func part_1(grid map[Point]bool, pos Point) int {
	curDir := 0
	res := 0
	for range 10000 {
		if grid[pos] {
			curDir = (curDir + 1) % 4
			delete(grid, pos)
		} else {
			curDir = (curDir + 3) % 4
			grid[pos] = true
			res++
		}

		switch curDir {
		case 0:
			pos.y--
		case 1:
			pos.x++
		case 2:
			pos.y++
		case 3:
			pos.x--
		}
	}
	return res
}

func part_2(grid map[Point]bool, pos Point) int {
	newGrid := map[Point]int{}
	for point := range grid {
		newGrid[point] = 2
	}

	curDir := 0
	res := 0
	for range 10000000 {
		switch newGrid[pos] {
		case 0:
			curDir = (curDir + 3) % 4
			newGrid[pos] = 1
		case 1:
			newGrid[pos] = 2
			res++
		case 2:
			curDir = (curDir + 1) % 4
			newGrid[pos] = 3
		case 3:
			curDir = (curDir + 2) % 4
			delete(newGrid, pos)
		}

		switch curDir {
		case 0:
			pos.y--
		case 1:
			pos.x++
		case 2:
			pos.y++
		case 3:
			pos.x--
		}
	}
	return res
}

func get_input() (map[Point]bool, Point) {
	input, _ := os.ReadFile("./Day_22/Day_22.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	grid := map[Point]bool{}
	for y, line := range lines {
		for x, char := range strings.TrimSpace(line) {
			if char == '#' {
				grid[Point{x: x, y: y}] = true
			}
		}
	}
	return grid, Point{len(strings.TrimSpace(lines[0])) / 2, len(lines) / 2}
}
