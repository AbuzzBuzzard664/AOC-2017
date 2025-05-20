package Day_07

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Run() {
	res := part_1()
	fmt.Println("Part 1: ", res)
	fmt.Println("Part 2: ", part_2(res))
}

func part_1() string {
	adjList := get_input()
	start := "snslv"
	for {
		found := false
		for key, value := range adjList {
			if slices.Contains(value, start) {
				start = key
				found = true
			}

		}
		if !found {
			return start
		}
	}
}

func part_2(base string) int {
	weights := get_input_weights()
	adjlist := get_input()
	_, imbalance := find_imbalance(base, adjlist, weights)
	return imbalance
}

func find_imbalance(node string, adjlist map[string][]string, weights map[string]int) (int, int) {
	if adjlist[node] == nil {
		return weights[node], 0
	}
	childWeights := []int{}
	weightMap := map[int][]string{}
	for _, child := range adjlist[node] {
		childWeight, imbalance := find_imbalance(child, adjlist, weights)
		if imbalance != 0 {
			return 0, imbalance
		}
		childWeights = append(childWeights, childWeight)
		weightMap[childWeight] = append(weightMap[childWeight], child)
	}
	if len(weightMap) > 1 {
		var correctWeight, incorrectWeight int
		for weight, nodes := range weightMap {
			if len(nodes) == 1 {
				incorrectWeight = weight
			} else {
				correctWeight = weight
			}
		}
		incorrectNode := weightMap[incorrectWeight][0]
		weightDiff := correctWeight - incorrectWeight
		return 0, weights[incorrectNode] + weightDiff
	}
	totalWeight := weights[node]
	for _, w := range childWeights {
		totalWeight += w
	}
	return totalWeight, 0
}

func get_input() map[string][]string {
	input, _ := os.ReadFile("./Day_07/Day_07.txt")
	adjList := map[string][]string{}
	for _, line := range strings.Split(string(input), "\n") {
		line = strings.ReplaceAll(line, "\r", " ")
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if strings.Contains(line, "->") {
			parts := strings.Split(line, "->")
			left := strings.TrimSpace(parts[0])
			right := strings.Split(strings.TrimSpace(parts[1]), ", ")
			key := strings.TrimSpace(left[:strings.Index(left, "(")])
			adjList[key] = right
		} else {
			left := strings.TrimSpace(line[:strings.Index(line, "(")])
			adjList[left] = nil
		}
	}
	return adjList
}

func get_input_weights() map[string]int {
	input, _ := os.ReadFile("./Day_07/Day_07.txt")
	weights := map[string]int{}
	for _, line := range strings.Split(string(input), "\n") {
		line = strings.ReplaceAll(line, "\r", " ")
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.Split(line, "(")
		numstr := ""
		for _, r := range parts[1] {
			if strings.Contains("0123456789", string(r)) {
				numstr += string(r)
			} else {
				break
			}
		}
		num, _ := strconv.Atoi(numstr)
		weights[strings.TrimSpace(parts[0])] = num
	}
	return weights
}
