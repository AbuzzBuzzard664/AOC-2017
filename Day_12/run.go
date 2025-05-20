package Day_12

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	fmt.Println("Part 1: ", part_1())
	fmt.Println("Part 2: ", part_2())
}

func part_1() int {
	nodes := get_input()

	visited := map[int]bool{}
	return dfs(nodes, 0, visited)
}

func part_2() int {
	nodes := get_input()
	visited := map[int]bool{}
	groupCount := 0
	for id := range nodes {
		if !visited[id] {
			dfs(nodes, id, visited)
			groupCount++
		}
	}
	return groupCount
}

func dfs(nodes map[int][]int, current int, visited map[int]bool) int {
	if visited[current] {
		return 0
	}
	res := 1
	visited[current] = true
	for _, neighbour := range nodes[current] {
		res += dfs(nodes, neighbour, visited)
	}
	return res
}

func get_input() map[int][]int {
	input, _ := os.ReadFile("./Day_12/Day_12.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	nodes := map[int][]int{}
	for _, line := range lines {
		// Skip empty or malformed lines
		if strings.TrimSpace(line) == "" || !strings.Contains(line, " <-> ") {
			continue
		}
		parts := strings.Split(line, " <-> ")
		id, _ := strconv.Atoi(parts[0])
		neighbors := strings.Split(parts[1], ",")
		node := []int{}
		for _, neighbor := range neighbors {
			neighborID, _ := strconv.Atoi(strings.TrimSpace(neighbor))
			node = append(node, neighborID)
		}
		nodes[id] = node
	}
	return nodes
}
