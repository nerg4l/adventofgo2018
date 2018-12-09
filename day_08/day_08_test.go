package day_08

import (
	"strings"
	"testing"
)

func TestSumMetadata(t *testing.T) {
	reader := strings.NewReader(`2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2
`)
	if SumMetadata(reader) != 138 {
		t.Error("Wrong sum")
	}
}

func TestSumMetadata2(t *testing.T) {
	// Child nodes swapped
	reader := strings.NewReader(`2 3 1 1 0 1 99 2 0 3 10 11 12 1 1 2
`)
	if SumMetadata(reader) != 138 {
		t.Error("Wrong sum")
	}
}

func TestCalcValueOfNode(t *testing.T) {
	reader := strings.NewReader(`2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2
`)
	if CalcValueOfNode(reader) != 66 {
		t.Error("Wrong sum")
	}
}

func TestCalcValueOfNode2(t *testing.T) {
	reader := strings.NewReader(`0 3 10 11 12
`)
	if CalcValueOfNode(reader) != 33 {
		t.Error("Wrong sum")
	}
}

func TestCalcValueOfNode3(t *testing.T) {
	reader := strings.NewReader(`1 1 0 1 99 2
`)
	if CalcValueOfNode(reader) != 0 {
		t.Error("Wrong sum")
	}
}

func TestCalcValueOfNode4(t *testing.T) {
	reader := strings.NewReader(`0 1 99
`)
	if CalcValueOfNode(reader) != 99 {
		t.Error("Wrong sum")
	}
}
