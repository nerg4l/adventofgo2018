package main

import (
	"strings"
	"testing"
)

func TestSimplifyPolymer(t *testing.T) {
	r := strings.NewReader(`dabAcCaCBAcCcaDA
`)
	if got, _ := SimplifyPolymer(r); got != len("dabCBAcaDA") {
		t.Error("Wrong polymer")
	}
}

func TestImprovePolymer(t *testing.T) {
	r := strings.NewReader("dabAcCaCBAcCcaDA")
	if got, _ := ImprovePolymer(r); got != 4 {
		t.Error("Wrong polymer")
	}
}
