package Day_18

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	instructions := get_input()
	fmt.Println("Part 1: ", part_1(instructions))
	fmt.Println("Part 2: ", part_2(instructions))
}

func part_1(instructions []string) int {
	return run_program(instructions)
}

func run_program(instructions []string) int {
	registers := map[string]int{}
	var snd int
	for i := 0; i < len(instructions); i++ {
		parts := strings.Fields(instructions[i])
		op := parts[0]
		x := parts[1]
		y := ""
		if len(parts) > 2 {
			y = parts[2]
		}
		switch op {
		case "snd":
			snd = getValue(x, registers)
		case "set":
			registers[x] = getValue(y, registers)
		case "add":
			registers[x] += getValue(y, registers)
		case "mul":
			registers[x] *= getValue(y, registers)
		case "mod":
			registers[x] %= getValue(y, registers)
		case "rcv":
			if getValue(x, registers) != 0 {
				return snd
			}
		case "jgz":
			if getValue(x, registers) > 0 {
				i += getValue(y, registers) - 1
			}
		}
	}
	return -1
}

func getValue(operand string, registers map[string]int) int {
	num, err := strconv.Atoi(operand)
	if err == nil {
		return num
	}
	return registers[operand]
}

type state struct {
	registers map[string]int
	ip        int
	queue     []int
	waiting   bool
	sent      int
}

func part_2(instructions []string) int {
	p0 := prog(0)
	p1 := prog(1)

	for {
		blocked0 := step(&p0, &p1, instructions)
		blocked1 := step(&p1, &p0, instructions)
		if blocked0 && blocked1 {
			break
		}
	}
	return p1.sent
}

func step(cur, other *state, instructions []string) bool {
	if cur.ip < 0 || cur.ip >= len(instructions) {
		cur.waiting = true
		return true
	}
	parts := strings.Fields(instructions[cur.ip])
	op := parts[0]
	x := parts[1]
	y := ""
	if len(parts) > 2 {
		y = parts[2]
	}
	switch op {
	case "snd":
		val := getValue(x, cur.registers)
		other.queue = append(other.queue, val)
		cur.sent++
	case "set":
		cur.registers[x] = getValue(y, cur.registers)
	case "add":
		cur.registers[x] += getValue(y, cur.registers)
	case "mul":
		cur.registers[x] *= getValue(y, cur.registers)
	case "mod":
		cur.registers[x] %= getValue(y, cur.registers)
	case "rcv":
		if len(cur.queue) == 0 {
			cur.waiting = true
			return true
		}
		cur.registers[x] = cur.queue[0]
		cur.queue = cur.queue[1:]
		cur.waiting = false
	case "jgz":
		if getValue(x, cur.registers) > 0 {
			cur.ip += getValue(y, cur.registers) - 1
		}
	}
	cur.ip++
	return false
}

func prog(id int) state {
	return state{
		registers: map[string]int{"p": id},
		ip:        0,
		queue:     []int{},
		waiting:   false,
		sent:      0,
	}
}

func get_input() []string {
	data, _ := os.ReadFile("./Day_18/Day_18.txt")
	return strings.Split(strings.TrimSpace(string(data)), "\n")
}
