package Day_20

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Vector struct {
	x, y, z int
}

type Particle struct {
	p, u, a Vector
}

func Run() {
	particles := get_input()
	fmt.Println("Part 1:", part_1(particles))
}

func part_1(particles []Particle) int {
	lowestAcc := math.MaxInt
	lowestVel := math.MaxInt
	lowestPos := math.MaxInt
	closestIndex := -1

	for i, particle := range particles {
		acc := manhattanDistance(particle.a)
		vel := manhattanDistance(particle.u)
		pos := manhattanDistance(particle.p)

		if acc < lowestAcc {
			lowestAcc = acc
			lowestVel = vel
			lowestPos = pos
			closestIndex = i
		} else if acc == lowestAcc && vel < lowestVel {
			lowestVel = vel
			lowestPos = pos
			closestIndex = i
		} else if acc == lowestAcc && vel == lowestVel && pos < lowestPos {
			lowestPos = pos
			closestIndex = i
		}
	}

	return closestIndex
}

func manhattanDistance(v Vector) int {
	return abs(v.x) + abs(v.y) + abs(v.z)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func get_input() []Particle {
	input, _ := os.ReadFile("./Day_20/Day_20.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	var particles []Particle
	for _, line := range lines {
		parts := strings.Split(line, ", ")

		p := Particle{
			p: parseVector(parts[0]),
			u: parseVector(parts[1]),
			a: parseVector(parts[2]),
		}
		particles = append(particles, p)
	}

	return particles
}

func parseVector(part string) Vector {
	start := strings.Index(part, "<")
	end := strings.Index(part, ">")
	coords := strings.Split(part[start+1:end], ",")
	var temp Vector

	temp.x, _ = strconv.Atoi(strings.TrimSpace(coords[0]))
	temp.y, _ = strconv.Atoi(strings.TrimSpace(coords[1]))
	temp.z, _ = strconv.Atoi(strings.TrimSpace(coords[2]))

	return temp
}
