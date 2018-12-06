package day_06

import (
	"strings"
	"testing"
)

func TestFindLargestArea(t *testing.T) {
	reader := strings.NewReader(`1, 1
1, 6
8, 3
3, 4
5, 5
8, 9`)
	if FindLargestArea(reader) != 17 {
		t.Error("Wrong area")
	}
}

func TestFindRegionNearManyCoordinates(t *testing.T) {
	reader := strings.NewReader(`1, 1
1, 6
8, 3
3, 4
5, 5
8, 9`)
	if FindRegionNearManyCoordinates(reader, 32) != 16 {
		t.Error("Wrong area")
	}
}
