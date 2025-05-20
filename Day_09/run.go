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
		if groups[pos] == '!' {
			pos++
		} else if groups[pos] == '<' {
			pos = skip_Garbage(pos, groups)
		} else if groups[pos] == '{' {
			res, pos = get_group_score(pos+1, 1, groups)
		}
	}
	return res
}

func get_group_score(pos, layer int, groups string) (int, int) {
	var num int
	for ; pos < len(groups); pos++ {
		if groups[pos] == '!' {
			pos++
		} else if groups[pos] == '<' {
			pos = skip_Garbage(pos, groups)
		} else if groups[pos] == '{' {
			var temp int
			temp, pos = get_group_score(pos+1, layer+1, groups)
			num += temp
		} else if groups[pos] == '}' {
			return num + layer, pos
		}
	}
	return 0, 0
}

func skip_Garbage(pos int, groups string) int {
	for ; pos < len(groups); pos++ {
		if groups[pos] == '!' {
			pos++
		} else if groups[pos] == '>' {
			return pos
		}
	}
	return -1
}

func part_2(groups string) int {
	res := 0
	for pos := 0; pos < len(groups); pos++ {
		if groups[pos] == '!' {
			pos++
		} else if groups[pos] == '<' {
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
		if groups[pos] == '!' {
			pos++
		} else if groups[pos] == '>' {
			return res, pos + 1
		} else {
			res++
		}
	}
	return -1, -1
}

func get_input() string {
	input, _ := os.ReadFile("./Day_09/Day_09.txt")
	return string(input)
}
