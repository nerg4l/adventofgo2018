package main

import (
	"strings"
	"testing"
)

func TestLargestArea(t *testing.T) {
	r := strings.NewReader(`1, 1
1, 6
8, 3
3, 4
5, 5
8, 9`)
	if got, _ := LargestArea(r); got != 17 {
		t.Error("Wrong area")
	}
}

func TestRegionNearManyCoordinates(t *testing.T) {
	r := strings.NewReader(`1, 1
1, 6
8, 3
3, 4
5, 5
8, 9`)
	if got, _ := RegionNearManyCoordinates(r, 32); got != 16 {
		t.Error("Wrong area")
	}
}
