package Day_17

import "fmt"

func Run() {
	input := get_input()
	fmt.Println("Part 1: ", part_1(input))
	fmt.Println("Part 2: ", part_2(input))
}

func part_1(steps int) int {
	circular := []int{0}
	k := 0
	for i := 1; i <= 2017; i++ {
		k = (k + steps) % len(circular)
		k++
		before := circular[:k]
		after := circular[k:]
		newElem := []int{i}
		circular = append(before, append(newElem, after...)...)
	}
	return circular[(k+1)%len(circular)]
}

func part_2(steps int) int {
	pos := 0
	result := 0
	for i := 1; i <= 50000000; i++ {
		pos = (pos + steps) % i
		if pos == 0 {
			result = i
		}
		pos++
	}
	return result
}

func get_input() int {
	return 344
}
