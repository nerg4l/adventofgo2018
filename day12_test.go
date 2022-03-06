package main

import (
	"strings"
	"testing"
)

func TestSumTheNumbersOfAllPots(t *testing.T) {
	r := strings.NewReader(`initial state: #..#.#..##......###...###

...## => #
..#.. => #
.#... => #
.#.#. => #
.#.## => #
.##.. => #
.#### => #
#.#.# => #
#.### => #
##.#. => #
##.## => #
###.. => #
###.# => #
####. => #
`)

	if got, _ := SumNumbersOfAllPots(r, 20); got != 325 {
		t.Error("Wrong number")
	}
}
