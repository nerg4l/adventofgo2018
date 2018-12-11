package day_11

import (
	"strings"
	"testing"
)

func TestFuelCell_PowerLevel(t *testing.T) {
	f := FuelCell{3, 5, 8}
	if f.PowerLevel() != 4 {
		t.Error("Wrong power level")
	}
}

func TestFuelCell_PowerLevel2(t *testing.T) {
	f := FuelCell{122, 79, 57}
	if f.PowerLevel() != -5 {
		t.Error("Wrong power level")
	}
}

func TestFuelCell_PowerLevel3(t *testing.T) {
	f := FuelCell{217, 196, 39}
	if f.PowerLevel() != 0 {
		t.Error("Wrong power level")
	}
}

func TestFuelCell_PowerLevel4(t *testing.T) {
	f := FuelCell{101, 153, 71}
	if f.PowerLevel() != 4 {
		t.Error("Wrong power level")
	}
}

func TestFindTheLargestTotalPowerWithDefaultSize(t *testing.T) {
	reader := strings.NewReader(`18
`)
	if FindTheLargestTotalPowerWithDefaultSize(reader) != "33,45" {
		t.Error("Wrong coordinates")
	}
}

func TestFindTheLargestTotalPowerWithDefaultSize2(t *testing.T) {
	reader := strings.NewReader(`42
`)
	if FindTheLargestTotalPowerWithDefaultSize(reader) != "21,61" {
		t.Error("Wrong coordinates")
	}
}

func TestFindTheLargestTotalPowerOfAllSize(t *testing.T) {
	reader := strings.NewReader(`18
`)
	if FindTheLargestTotalPowerOfAllSize(reader) != "90,269,16" {
		t.Error("Wrong coordinates and size")
	}
}

func TestFindTheLargestTotalPowerOfAllSize2(t *testing.T) {
	reader := strings.NewReader(`42
`)
	if FindTheLargestTotalPowerOfAllSize(reader) != "232,251,12" {
		t.Error("Wrong coordinates and size")
	}
}
