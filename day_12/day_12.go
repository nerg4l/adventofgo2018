package day_12

import (
	"bufio"
	"fmt"
	"io"
	"reflect"
)

type rule struct {
	in  []bool
	out bool
}

func (r *rule) matches(a []bool) bool {
	return reflect.DeepEqual(r.in, a)
}

func SumNumbersOfAllPots(reader io.Reader, generations int) int {
	scanner := bufio.NewScanner(reader)
	state := []bool{false, false, false, false, false}
	var rules []*rule
	for i := 0; scanner.Scan(); i++ {
		text := scanner.Text()

		if i == 0 {
			state = append(state, readInitialState(text)...)
			continue
		} else if text != "" {
			rules = append(rules, readRule(text))
		}
	}
	state = append(state, false, false, false, false, false)
	shift := 5
	for i := 0; i < generations; i++ {
		var newState []bool
		for _, v := range state {
			newState = append(newState, v)
		}
		for j := 2; j < len(state)-2; j++ {
			in := state[j-2 : j+3]
			found := false
			for _, v := range rules {
				if v.matches(in) {
					newState[j] = v.out
					found = true
					break
				}
			}
			if !found {
				newState[j] = false
			}
		}
		oShift := shift
		for !reflect.DeepEqual(newState[:5], []bool{false, false, false, false, false}) {
			shift++
			newState = append([]bool{false}, newState...)
		}
		for reflect.DeepEqual(newState[:6], []bool{false, false, false, false, false, false}) {
			shift--
			newState = newState[1:]
		}
		for !reflect.DeepEqual(newState[len(newState)-5:], []bool{false, false, false, false, false}) {
			newState = append(newState, false)
		}
		for reflect.DeepEqual(newState[len(newState)-5:], []bool{false, false, false, false, false, false}) {
			newState = newState[:len(newState)-1]
		}
		if !reflect.DeepEqual(state, newState) {
			state = newState
			continue
		}
		shift += (shift - oShift) * (generations - i - 1)
		break
	}
	var result int
	for j := 2; j < len(state)-2; j++ {
		if state[j] {
			result += j - shift
		}
	}
	return result
}

func readRule(s string) *rule {
	var inS string
	var outR rune
	_, _ = fmt.Sscanf(s, "%s => %c", &inS, &outR)
	r := &rule{}
	for _, v := range inS {
		r.in = append(r.in, v == '#')
	}
	r.out = outR == '#'
	return r
}

func readInitialState(s string) []bool {
	var stateS string
	_, _ = fmt.Sscanf(s, "initial state: %s", &stateS)
	var state []bool
	for _, v := range stateS {
		state = append(state, v == '#')
	}
	return state
}
