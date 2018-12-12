package day_12

import (
	"strings"
	"testing"
)

func TestSumTheNumbersOfAllPots(t *testing.T) {
	reader := strings.NewReader(`initial state: #..#.#..##......###...###

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

	if SumNumbersOfAllPots(reader, 20) != 325 {
		t.Error("Wrong number")
	}
}
