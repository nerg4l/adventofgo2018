package day_05

import (
	"strings"
	"testing"
)

func TestSimplifyPolimer(t *testing.T) {
	reader := strings.NewReader(`dabAcCaCBAcCcaDA
`)
	if SimplifyPolymer(reader) != len("dabCBAcaDA") {
		t.Error("Wrong polymer")
	}
}

func TestImprovePolymer(t *testing.T) {
	reader := strings.NewReader("dabAcCaCBAcCcaDA")
	if ImprovePolymer(reader) != 4 {
		t.Error("Wrong polymer")
	}
}
