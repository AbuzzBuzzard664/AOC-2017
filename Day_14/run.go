package Day_14

import (
	"fmt"
	"os"
	"strconv"
)

func Run() {
	input := get_input()
	fmt.Println("Part 1:", part_1(input))
	fmt.Println("Part 2:", part_2(input))
}

func part_1(input string) int {
	total := 0
	for i := 0; i < 128; i++ {
		rowInput := input + "-" + strconv.Itoa(i)
		hash := knotHash(rowInput)
		total += countOnesInHex(hash)
	}
	return total
}

func knotHash(input string) string {
	lengths := []int{}
	for _, c := range input {
		lengths = append(lengths, int(c))
	}
	lengths = append(lengths, 17, 31, 73, 47, 23)

	circle := []int{}
	for i := 0; i < 256; i++ {
		circle = append(circle, i)
	}
	pos, skip := 0, 0
	for round := 0; round < 64; round++ {
		for _, l := range lengths {
			circle, pos, skip = hash(circle, pos, skip, l)
		}
	}

	dense := []int{}
	for i := 0; i < 16; i++ {
		xor := circle[i*16]
		for j := 1; j < 16; j++ {
			xor ^= circle[i*16+j]
		}
		dense = append(dense, xor)
	}

	hex := ""
	for _, v := range dense {
		hex += fmt.Sprintf("%02x", v)
	}
	return hex
}

func hash(circular []int, index, skip, length int) ([]int, int, int) {
	for i := 0; i < length/2; i++ {
		start := (index + i) % len(circular)
		end := (index + length - 1 - i) % len(circular)
		circular[start], circular[end] = circular[end], circular[start]
	}
	index = (index + length + skip) % len(circular)
	skip++
	return circular, index, skip
}

func countOnesInHex(hex string) int {
	count := 0
	for _, c := range hex {
		n, _ := strconv.ParseUint(string(c), 16, 4)
		for i := 0; i < 4; i++ {
			if (n>>i)&1 == 1 {
				count++
			}
		}
	}
	return count
}

func part_2(input string) int {
	grid := make([][]int, 128)
	for i := 0; i < 128; i++ {
		rowInput := input + "-" + strconv.Itoa(i)
		hash := knotHash(rowInput)
		grid[i] = hexToBinaryRow(hash)
	}

	visited := make([][]bool, 128)
	for i := range visited {
		visited[i] = make([]bool, 128)
	}

	regions := 0
	for i := 0; i < 128; i++ {
		for j := 0; j < 128; j++ {
			if grid[i][j] == 1 && !visited[i][j] {
				regions++
				floodFill(grid, visited, i, j)
			}
		}
	}

	return regions
}

func hexToBinaryRow(hex string) []int {
	row := []int{}
	for _, c := range hex {
		n, _ := strconv.ParseUint(string(c), 16, 4)
		for i := 3; i >= 0; i-- {
			row = append(row, int((n>>i)&1))
		}
	}
	return row
}

func floodFill(grid [][]int, visited [][]bool, x, y int) {
	if x < 0 || x >= 128 || y < 0 || y >= 128 || visited[x][y] || grid[x][y] == 0 {
		return
	}
	visited[x][y] = true
	floodFill(grid, visited, x+1, y)
	floodFill(grid, visited, x-1, y)
	floodFill(grid, visited, x, y+1)
	floodFill(grid, visited, x, y-1)
}

func get_input() string {
	input, _ := os.ReadFile("./Day_14/Day_14.txt")
	return string(input)
}
