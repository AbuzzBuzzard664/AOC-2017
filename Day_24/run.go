package Day_24

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Connection struct {
	aPort, bPort int
}

func Run() {
	input := get_input()
	fmt.Println("Part 1:", part_1(input))
	fmt.Println("Part 2:", part_2(input))
}

func part_1(input []Connection) int {
	res := 0
	for i, c := range input {
		if c.aPort == 0 || c.bPort == 0 {
			nextRemaining := removeAt(input, i)
			nextPort := c.bPort
			if c.bPort == 0 {
				nextPort = c.aPort
			}
			strength := c.aPort + c.bPort + helper(nextRemaining, nextPort)
			if strength > res {
				res = strength
			}
		}
	}
	return res
}

func helper(remaining []Connection, match int) int {
	best := 0
	for i, r := range remaining {
		if r.aPort == match || r.bPort == match {
			nextRemaining := removeAt(remaining, i)
			nextPort := r.bPort
			if r.bPort == match {
				nextPort = r.aPort
			}
			strength := r.aPort + r.bPort + helper(nextRemaining, nextPort)
			if strength > best {
				best = strength
			}
		}
	}
	return best
}

func removeAt(slice []Connection, index int) []Connection {
	out := make([]Connection, 0, len(slice)-1)
	out = append(out, slice[:index]...)
	out = append(out, slice[index+1:]...)
	return out
}

func part_2(input []Connection) int {
	bestStrength := 0
	bestLength := 0

	for i, c := range input {
		if c.aPort == 0 || c.bPort == 0 {
			nextRemaining := removeAt(input, i)
			nextPort := c.bPort
			if c.bPort == 0 {
				nextPort = c.aPort
			}

			subStrength, subLength := helper_2(nextRemaining, nextPort, 1)
			totalStrength := c.aPort + c.bPort + subStrength

			if subLength > bestLength || (subLength == bestLength && totalStrength > bestStrength) {
				bestLength = subLength
				bestStrength = totalStrength
			}
		}
	}

	return bestStrength
}

func helper_2(remaining []Connection, match int, length int) (int, int) {
	bestStrength := 0
	bestLength := length

	for i, r := range remaining {
		if r.aPort == match || r.bPort == match {
			nextRemaining := removeAt(remaining, i)
			nextPort := r.bPort
			if r.bPort == match {
				nextPort = r.aPort
			}

			subStrength, branchLength := helper_2(nextRemaining, nextPort, length+1)
			totalStrength := r.aPort + r.bPort + subStrength

			if branchLength > bestLength || (branchLength == bestLength && totalStrength > bestStrength) {
				bestLength = branchLength
				bestStrength = totalStrength
			}
		}
	}

	return bestStrength, bestLength
}

func get_input() []Connection {
	input, _ := os.ReadFile("./Day_24/Day_24.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	connections := []Connection{}

	for _, line := range lines {
		parts := strings.Split(strings.TrimSpace(line), "/")
		aPort, _ := strconv.Atoi(parts[0])
		bPort, _ := strconv.Atoi(parts[1])
		connections = append(connections, Connection{aPort: aPort, bPort: bPort})
	}
	return connections
}
