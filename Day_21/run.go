package Day_21

import (
	"fmt"
	"os"
	"strings"
)

func Run() {
	rules := get_input()
	fmt.Println("Part 1:", part_1(rules))
	fmt.Println("Part 2:", part_2(rules))
}

func part_1(rules map[string]string) int {
	return solve(rules, 5)
}

func part_2(rules map[string]string) int {
	return solve(rules, 18)
}

func solve(rules map[string]string, iter int) int {
	grid := []string{
		".#.",
		"..#",
		"###",
	}

	for range iter {
		n := len(grid)
		step := 3
		if n%2 == 0 {
			step = 2
		}

		blocks := n / step
		newStep := step + 1
		nextGrid := make([]string, blocks*newStep)

		for br := range blocks {
			for bc := range blocks {
				var sub []string
				for r := range step {
					row := grid[br*step+r]
					sub = append(sub, row[bc*step:(bc+1)*step])
				}
				key := strings.Join(sub, "/")

				outRows := strings.Split(rules[key], "/")

				for r := range newStep {
					nextGrid[br*newStep+r] += outRows[r]
				}
			}
		}

		grid = nextGrid
	}

	return count_hash(strings.Join(grid, "/"))
}

func get_flips_rotations(grid string) []string {
	rows := strings.Split(grid, "/")
	n := len(rows)
	curr := make([][]byte, n)
	for i := range rows {
		curr[i] = []byte(rows[i])
	}

	var results []string
	for range 4 {
		results = append(results, joinGrid(curr))
		results = append(results, joinGrid(flipHorizontal(curr)))
		curr = rotate90(curr)
	}

	return results
}

func rotate90(g [][]byte) [][]byte {
	res := make([][]byte, len(g))
	for i := range len(g) {
		res[i] = make([]byte, len(g))
		for j := range len(g) {
			res[i][j] = g[len(g)-1-j][i]
		}
	}
	return res
}

func flipHorizontal(g [][]byte) [][]byte {
	res := make([][]byte, len(g))
	for i := range len(g) {
		res[i] = make([]byte, len(g))
		for j := range len(g) {
			res[i][j] = g[i][len(g)-1-j]
		}
	}
	return res
}

func joinGrid(g [][]byte) string {
	rows := make([]string, len(g))
	for i, row := range g {
		rows[i] = string(row)
	}
	return strings.Join(rows, "/")
}

func count_hash(grid string) int {
	res := 0
	for _, char := range grid {
		if char == '#' {
			res++
		}
	}
	return res
}

func get_input() map[string]string {
	input, _ := os.ReadFile("./Day_21/Day_21.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	dict := map[string]string{}
	for _, line := range lines {
		splits := strings.Split(strings.TrimSpace(line), "=>")
		for _, temp := range get_flips_rotations(strings.TrimSpace(splits[0])) {
			dict[temp] = strings.TrimSpace(splits[1])
		}
	}
	return dict
}
