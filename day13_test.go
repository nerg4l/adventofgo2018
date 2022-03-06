package main

import (
	"strings"
	"testing"
)

func TestLocationOfFirstCrash(t *testing.T) {
	r := strings.NewReader(`/->-\
|   |  /----\
| /-+--+-\  |
| | |  | v  |
\-+-/  \-+--/
\------/
`)

	if got, _ := LocationOfFirstCrash(r); got != "7,3" {
		t.Error("Wrong location")
	}
}

func TestLocationOfLastCart(t *testing.T) {
	r := strings.NewReader(`/>-<\  
|   |  
| /<+-\
| | | v
\>+</ |
  |   ^
  \<->/
`)

	if got, _ := LocationOfLastCart(r); got != "6,4" {
		t.Error("Wrong location")
	}
}
