package Day_16

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	dances := get_input()
	fmt.Println("Part 1", part_1(dances))
	fmt.Println("Part 2", part_2(dances))
}

func part_1(dances []string) string {
	programs := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p'}
	programs = dance(programs, dances)
	return makeString(programs)
}

func part_2(dances []string) string {
	programs := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p'}
	seen := make(map[string]int)
	order := make([]string, 0)
	iterations := 1000000000

	for i := 0; i < iterations; i++ {
		key := makeString(programs)
		if prev, ok := seen[key]; ok {
			cycleLen := i - prev
			remaining := (iterations - i) % cycleLen
			return order[prev+remaining]
		}
		seen[key] = i
		order = append(order, key)
		programs = dance(programs, dances)
	}
	return makeString(programs)
}

func makeString(programs []rune) string {
	res := ""
	for _, c := range programs {
		res += string(rune(c))
	}
	return res
}

func dance(programs []rune, dances []string) []rune {
	for _, s := range dances {
		switch s[0] {
		case 's':
			num, _ := strconv.Atoi(s[1:])
			programs = spin(num, programs)
		case 'x':
			numbers := strings.Split(s[1:], "/")
			a, _ := strconv.Atoi(numbers[0])
			b, _ := strconv.Atoi(numbers[1])
			programs = exchange(a, b, programs)
		case 'p':
			numbers := strings.Split(s[1:], "/")
			a := rune(numbers[0][0])
			b := rune(numbers[1][0])
			programs = partner(a, b, programs)
		}
	}
	return programs
}

func spin(X int, programs []rune) []rune {
	n := len(programs)
	X = X % n
	return append(programs[n-X:], programs[:n-X]...)
}

func exchange(A, B int, programs []rune) []rune {
	programsCopy := make([]rune, len(programs))
	copy(programsCopy, programs)
	programsCopy[A], programsCopy[B] = programsCopy[B], programsCopy[A]
	return programsCopy
}
func partner(A, B rune, programs []rune) []rune {
	Aind := -1
	Bind := -1
	for i, c := range programs {
		if c == A {
			Aind = i
		}
		if c == B {
			Bind = i
		}
	}
	if Aind == -1 || Bind == -1 {
		return programs
	}
	return exchange(Aind, Bind, programs)
}

func get_input() []string {
	data, _ := os.ReadFile("./Day_16/Day_16.txt")
	return strings.Split(strings.TrimSpace(string(data)), ",")
}
