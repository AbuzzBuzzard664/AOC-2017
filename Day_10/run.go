package Day_10

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	input := get_input()
	fmt.Println("Part 1: ", part_1(get_num_list(input)))
	fmt.Println("Part 2: ", part_2(get_ascii_codes(input)))
}

func part_1(num_list []int) int {
	circular := []int{}
	for i := 0; i <= 255; i++ {
		circular = append(circular, i)
	}
	index, skip := 0, 0
	for _, num := range num_list {
		circular, index, skip = hash(circular, index, skip, num)
	}
	return circular[0] * circular[1]
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

func part_2(ascii_list []int) string {
	circular := []int{}
	for i := 0; i <= 255; i++ {
		circular = append(circular, i)
	}
	index, skip := 0, 0
	for i := 0; i < 64; i++ {
		for _, length := range ascii_list {
			circular, index, skip = hash(circular, index, skip, length)
		}
	}
	denseHash := dense_hash(circular)
	hashString := ""
	for _, num := range denseHash {
		hashString += fmt.Sprintf("%02x", num)
	}
	return hashString
}

func dense_hash(circular []int) []int {
	denseHash := []int{}
	for i := 0; i < 16; i++ {
		block := circular[i*16]
		for j := 1; j < 16; j++ {
			block ^= circular[i*16+j]
		}
		denseHash = append(denseHash, block)
	}
	return denseHash
}

func get_input() string {
	input, _ := os.ReadFile("./Day_10/Day_10.txt")
	return string(input)
}

func get_num_list(input string) []int {
	slice := strings.Split(strings.TrimSpace(string(input)), ",")
	num_list := []int{}
	for _, s := range slice {
		num, _ := strconv.Atoi(s)
		num_list = append(num_list, num)
	}
	return num_list
}

func get_ascii_codes(input string) []int {
	ascii := []int{}
	for _, r := range input {
		ascii = append(ascii, int(r))
	}
	ascii = append(ascii, 17, 31, 73, 47, 23)
	return ascii
}
