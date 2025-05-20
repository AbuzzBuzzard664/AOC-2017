package Day_11

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func Run() {
	fmt.Println("Part 1: ", part_1(get_input()))
	fmt.Println("Part 2: ", part_2(get_input()))
}

func part_1(directions []string) int {
	x, y, z := 0, 0, 0
	for _, direction := range directions {
		x, y, z = move(x, y, z, direction)
	}
	return distance(0, 0, 0, x, y, z)
}

func part_2(directions []string) int {
	x, y, z := 0, 0, 0
	best := 0
	for _, direction := range directions {
		x, y, z = move(x, y, z, direction)
		if num := distance(0, 0, 0, x, y, z); num > best {
			best = num
		}
	}
	return best
}

func move(x, y, z int, direction string) (int, int, int) {
	switch direction {
	case "n":
		y++
		z--
	case "s":
		y--
		z++
	case "ne":
		x++
		z--
	case "nw":
		x--
		y++
	case "se":
		x++
		y--
	case "sw":
		x--
		z++
	}
	return x, y, z
}

func distance(x1, y1, z1, x2, y2, z2 int) int {
	return int(math.Max(math.Abs(float64(x2-x1)), math.Max(math.Abs(float64(y2-y1)), math.Abs(float64(z2-z1)))))
}

func get_input() []string {
	input, _ := os.ReadFile("./Day_11/Day_11.txt")
	slice := strings.Split(strings.TrimSpace(string(input)), ",")
	return slice
}
