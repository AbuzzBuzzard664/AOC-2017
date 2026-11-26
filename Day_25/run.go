package Day_25

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Transition struct {
	Write     int
	Move      int
	NextState string
}

type StateRules struct {
	OnZero Transition
	OnOne  Transition
}

type TuringSpec struct {
	StartState string
	Steps      int
	States     map[string]*StateRules
}

func Run() {
	input := get_input()
	fmt.Println("Part 1:", part_1(input))
}

func part_1(input TuringSpec) int {
	tape := make(map[int]int)
	cursor := 0
	state := input.StartState
	for range input.Steps {
		rules := input.States[state]
		var tr Transition
		if tape[cursor] == 0 {
			tr = rules.OnZero
		} else {
			tr = rules.OnOne
		}

		if tr.Write == 0 {
			delete(tape, cursor)
		} else {
			tape[cursor] = tr.Write
		}
		cursor += tr.Move
		state = tr.NextState
	}
	return len(tape)
}

func get_input() TuringSpec {
	input, _ := os.ReadFile("./Day_25/Day_25.txt")
	spec := TuringSpec{
		States: make(map[string]*StateRules),
	}

	var currentState string
	var currentVal int // 0 or 1

	ensureState := func(name string) *StateRules {
		if _, ok := spec.States[name]; !ok {
			spec.States[name] = &StateRules{}
		}
		return spec.States[name]
	}

	currentTransition := func() *Transition {
		rules := ensureState(currentState)
		if currentVal == 0 {
			return &rules.OnZero
		}
		return &rules.OnOne
	}

	lines := strings.Split(string(input), "\n")
	for _, rawLine := range lines {
		line := strings.TrimSpace(rawLine)
		if line == "" {
			continue
		}

		switch {
		case strings.HasPrefix(line, "Begin in state "):
			spec.StartState = strings.TrimSuffix(strings.TrimPrefix(line, "Begin in state "), ".")

		case strings.HasPrefix(line, "Perform a diagnostic checksum after "):
			rest := strings.TrimPrefix(line, "Perform a diagnostic checksum after ")
			fields := strings.Fields(rest)
			steps, _ := strconv.Atoi(fields[0])
			spec.Steps = steps

		case strings.HasPrefix(line, "In state "):
			currentState = strings.TrimSuffix(strings.TrimPrefix(line, "In state "), ":")
			ensureState(currentState)

		case line == "If the current value is 0:":
			currentVal = 0

		case line == "If the current value is 1:":
			currentVal = 1

		case strings.HasPrefix(line, "- Write the value "):
			valStr := strings.TrimSuffix(strings.TrimPrefix(line, "- Write the value "), ".")
			val, _ := strconv.Atoi(valStr)
			currentTransition().Write = val

		case line == "- Move one slot to the right.":
			currentTransition().Move = 1

		case line == "- Move one slot to the left.":
			currentTransition().Move = -1

		case strings.HasPrefix(line, "- Continue with state "):
			next := strings.TrimSuffix(strings.TrimPrefix(line, "- Continue with state "), ".")
			currentTransition().NextState = next
		}
	}

	return spec
}
