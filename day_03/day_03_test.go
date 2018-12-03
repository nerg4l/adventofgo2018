package day_03

import (
	"strings"
	"testing"
)

func TestCountFabricOverlap(t *testing.T) {
	reader := strings.NewReader(`#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2`)
	if CountFabricOverlap(reader) != 4 {
		t.Error("Wrong overlap")
	}
}

func TestFindNotOverlappingFabric(t *testing.T) {
	reader := strings.NewReader(`#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2`)
	if FindNotOverlappingFabric(reader) != 3 {
		t.Error("Wrong overlap")
	}
}
