package day_13

import (
	"strings"
	"testing"
)

func TestSumTheNumbersOfAllPots(t *testing.T) {
	reader := strings.NewReader(`/->-\
|   |  /----\
| /-+--+-\  |
| | |  | v  |
\-+-/  \-+--/
\------/
`)

	if FindLocationOfFirstCrash(reader) != "7,3" {
		t.Error("Wrong location")
	}
}

func TestFindLocationOfLastCart(t *testing.T) {
	reader := strings.NewReader(`/>-<\  
|   |  
| /<+-\
| | | v
\>+</ |
  |   ^
  \<->/
`)

	if FindLocationOfLastCart(reader) != "6,4" {
		t.Error("Wrong location")
	}
}
