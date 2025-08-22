package Day_19

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

func Run() {
	diagram := get_input()
	fmt.Println("Part 1: ", part_1(diagram))
	fmt.Println("Part 2: ", part_2(diagram))
}

func part_1(diagram [][]rune) string {
	letters := ""
	start := getStartPoint(diagram)
	current := start
	direction := 'S'
	for {
		current = movePacket(current, direction)
		if !isValid(diagram, current) {
			break
		}
		cell := diagram[current.y][current.x]
		if strings.Contains("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", string(cell)) {
			letters += string(cell)
		}
		if cell == '+' {
			direction = findDirection(diagram, current, direction)
		}
	}
	return letters
}

func movePacket(current Point, direction rune) Point {
	var temp Point
	switch direction {
	case 'N':
		temp.y = current.y - 1
		temp.x = current.x
	case 'E':
		temp.y = current.y
		temp.x = current.x + 1
	case 'S':
		temp.y = current.y + 1
		temp.x = current.x
	case 'W':
		temp.y = current.y
		temp.x = current.x - 1
	}
	return temp
}

func isValid(diagram [][]rune, current Point) bool {
	if current.x < 0 || current.x >= len(diagram[0]) || current.y < 0 || current.y >= len(diagram) {
		return false
	}
	cell := diagram[current.y][current.x]
	return cell != ' '
}

func findDirection(diagram [][]rune, current Point, direction rune) rune {
	switch direction {
	case 'N', 'S':
		for _, d := range []rune{'E', 'W'} {
			next := movePacket(current, d)
			if isValid(diagram, next) {
				return d
			}
		}
	case 'E', 'W':
		for _, d := range []rune{'N', 'S'} {
			next := movePacket(current, d)
			if isValid(diagram, next) {
				return d
			}
		}
	}
	return direction
}

func getStartPoint(diagram [][]rune) Point {
	for i, char := range diagram[0] {
		if char == '|' {
			var res Point
			res.x = i
			res.y = 0
			return res
		}
	}
	return Point{0, 0}
}

func part_2(diagram [][]rune) int {
	start := getStartPoint(diagram)
	current := start
	direction := 'S'
	steps := 1
	for {
		current = movePacket(current, direction)
		if !isValid(diagram, current) {
			break
		}
		steps++
		cell := diagram[current.y][current.x]
		if cell == '+' {
			direction = findDirection(diagram, current, direction)
		}
	}
	return steps
}

func get_input() [][]rune {

	diagram := [][]rune{}
	data, _ := os.ReadFile("./Day_19/Day_19.txt")
	lines := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	maxLen := 0
	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}
	for _, line := range lines {
		temp := []rune(line)
		// Pad the line with spaces if it's shorter than maxLen
		for len(temp) < maxLen {
			temp = append(temp, ' ')
		}
		diagram = append(diagram, temp)
	}
	return diagram
}
