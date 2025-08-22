package Day_20

import (
	"fmt"
	"os"
	"regexp"
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
	fmt.Println("Part 1: ", part_1(particles))
}

func part_1(particles []Particle) int {
	closests := []int{}
	for _, _ = range particles {
		closests = append(closests, 0)
	}
	for i := 1; i <= 10000; i++ {
		best := get_Manhatten_Distance(particles[0])
		closest := 0
		for j, particle := range particles {
			if new := get_Manhatten_Distance(particle); best > new {
				best = new
				closest = j
			}
		}
		closests[closest]++
	}
	res := 0
	best := 0
	for i, num := range closests {
		if best < num {
			best = num
			res = i
		}
	}
	return res
}

func part_2() {

}

func get_Manhatten_Distance(p Particle) int {
	return abs(p.p.x) + abs(p.p.y) + abs(p.p.z)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func get_input() []Particle {
	input, _ := os.ReadFile("./Day_20.txt")
	lines := strings.Split(strings.TrimRight(string(input), "\n"), "\n")
	particles := []Particle{}
	for _, line := range lines {
		splits := strings.Split(line, ", ")
		var temp Particle
		temp.p = parseString(splits[0])
		temp.u = parseString(splits[1])
		temp.a = parseString(splits[2])
		particles = append(particles, temp)
	}
	return particles
}

func parseString(split string) Vector {
	r, _ := regexp.Compile(`-?\d+(,|>)`)
	var temp Vector
	numbers := r.FindAllString(split, -1)
	for i := range numbers {
		numbers[i] = strings.Trim(numbers[i], ",>")
	}
	temp.x, _ = strconv.Atoi(string(numbers[0]))
	temp.y, _ = strconv.Atoi(string(numbers[1]))
	temp.z, _ = strconv.Atoi(string(numbers[2]))
	return temp
}
