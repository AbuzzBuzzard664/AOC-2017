package Day_03

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

type Point struct {
	x, y int
}

func Run() {
	fmt.Println("Part 1: ", part_1())
	fmt.Println("Part 2: ", part_2())
}

func part_1() int {
	n := get_input()
	i := 1
	// Find the layer of the spiral where n is located
	for i*i < n {
		i += 2
	}

	// Calculate the pivot points for the current layer
	pivots := []int{}
	for k := 0; k < 4; k++ {
		pivots = append(pivots, i*i-k*(i-1))
	}

	// Find the Manhattan distance
	for _, p := range pivots {
		dist := int(math.Abs(float64(p - n)))
		if dist <= (i-1)/2 {
			return (i - 1) - dist
		}
	}

	return -1 // Should not reach here
}

func part_2() int {
	n := get_input()
	x, y, dx, dy := 0, 0, 0, -1
	grid := make(map[Point]int)

	coords := []Point{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1} /*(0,0)*/, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for {
		total := 0
		for _, offset := range coords {
			ox, oy := offset.x, offset.y
			total += grid[Point{x + ox, y + oy}]
		}

		if total > n {
			return total
		}

		if x == 0 && y == 0 {
			grid[Point{0, 0}] = 1
		} else {
			grid[Point{x, y}] = total
		}

		// Change direction if at a corner
		if x == y || (x < 0 && x == -y) || (x > 0 && x == 1-y) {
			dx, dy = -dy, dx
		}

		x, y = x+dx, y+dy
	}
}

func get_input() int {
	input, _ := os.ReadFile("./Day_03/Day_03.txt")
	num, _ := strconv.Atoi(string(input))
	return num
}
