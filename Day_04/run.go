package Day_04

import (
	"fmt"
	"os"
	"strings"
)

func Run() {
	fmt.Println("Part 1: ", part_1())
	fmt.Println("Part 2: ", part_2())
}

func part_1() int {
	res := 0
	passphrases := get_input()
	for _, line := range passphrases {
		if is_Valid1(line) {
			res++
		}
	}
	return res
}

func part_2() int {
	res := 0
	passphrases := get_input()
	for _, line := range passphrases {
		if is_Valid2(line) {
			res++
		}
	}
	return res
}

func is_Valid1(passcode []string) bool {
	for i, s1 := range passcode {
		for j, s2 := range passcode {
			if strings.TrimSpace(s1) == strings.TrimSpace(s2) && i != j {
				return false
			}
		}
	}
	return true
}

func is_Valid2(passcode []string) bool {
	for i, s1 := range passcode {
		for j, s2 := range passcode {
			if i != j {
				if sort(s1) == sort(s2) {
					return false
				}
			}
		}
	}
	return true
}

func sort(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		for j := i; j < len(runes); j++ {
			if runes[i] > runes[j] {
				runes[i], runes[j] = runes[j], runes[i]
			}
		}
	}
	return string(runes)
}

func get_input() [][]string {

	input, _ := os.ReadFile("./Day_04/Day_04.txt")
	passphrases := [][]string{}
	for _, line := range strings.Split(string(input), "\n") {
		line = strings.TrimSuffix(line, "\r")
		if line == "" {
			continue
		}
		temp := strings.Split(line, " ")
		passphrases = append(passphrases, temp)
	}
	return passphrases
}
