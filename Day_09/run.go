package Day_09

import (
	"fmt"
	"os"
)

func Run() {
	groups := get_input()
	fmt.Println("Part 1: ", part_1(groups))
	fmt.Println("Part 2: ", part_2(groups))
}

func part_1(groups string) int {
	res := 0
	for pos := 0; pos < len(groups); pos++ {
		switch groups[pos] {
		case '!':
			pos++
		case '<':
			pos = skip_Garbage(pos, groups)
		case '{':
			res, pos = get_group_score(pos+1, 1, groups)
		}
	}
	return res
}

func get_group_score(pos, layer int, groups string) (int, int) {
	var num int
	for ; pos < len(groups); pos++ {
		switch groups[pos] {
		case '!':
			pos++
		case '<':
			pos = skip_Garbage(pos, groups)
		case '{':
			var temp int
			temp, pos = get_group_score(pos+1, layer+1, groups)
			num += temp
		case '}':
			return num + layer, pos
		}
	}
	return 0, 0
}

func skip_Garbage(pos int, groups string) int {
	for ; pos < len(groups); pos++ {
		switch groups[pos] {
		case '!':
			pos++
		case '>':
			return pos
		}
	}
	return -1
}

func part_2(groups string) int {
	res := 0
	for pos := 0; pos < len(groups); pos++ {
		switch groups[pos] {
		case '!':
			pos++
		case '<':
			var temp int
			temp, pos = count_Garbage(pos+1, groups)
			res += temp
		}
	}
	return res
}

func count_Garbage(pos int, groups string) (int, int) {
	res := 0
	for ; pos < len(groups); pos++ {
		switch groups[pos] {
		case '!':
			pos++
		case '>':
			return res, pos + 1
		default:
			res++
		}
	}
	return -1, -1
}

func get_input() string {
	input, _ := os.ReadFile("./Day_09/Day_09.txt")
	return string(input)
}
