package day_09

import (
	"strings"
	"testing"
)

func TestCalcWinningElfsScore(t *testing.T) {
	reader := strings.NewReader(`9 players; last marble is worth 25 points
`)
	if CalcWinningElfsScore(reader, 1) != 32 {
		t.Error("Wrong sum")
	}
}

func TestCalcWinningElfsScore2(t *testing.T) {
	reader := strings.NewReader(`10 players; last marble is worth 1618 points
`)
	if CalcWinningElfsScore(reader, 1) != 8317 {
		t.Error("Wrong sum")
	}
}

func TestCalcWinningElfsScore3(t *testing.T) {
	reader := strings.NewReader(`13 players; last marble is worth 7999 points
`)
	if CalcWinningElfsScore(reader, 1) != 146373 {
		t.Error("Wrong sum")
	}
}

func TestCalcWinningElfsScore4(t *testing.T) {
	reader := strings.NewReader(`17 players; last marble is worth 1104 points
`)
	if CalcWinningElfsScore(reader, 1) != 2764 {
		t.Error("Wrong sum")
	}
}

func TestCalcWinningElfsScore5(t *testing.T) {
	reader := strings.NewReader(`21 players; last marble is worth 6111 points
`)
	if CalcWinningElfsScore(reader, 1) != 54718 {
		t.Error("Wrong sum")
	}
}

func TestCalcWinningElfsScore6(t *testing.T) {
	reader := strings.NewReader(`30 players; last marble is worth 5807 points
`)
	if CalcWinningElfsScore(reader, 1) != 37305 {
		t.Error("Wrong sum")
	}
}
