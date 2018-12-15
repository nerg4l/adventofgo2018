package day_14

import (
	"strings"
	"testing"
)

func TestFindUltimateHotChocolateRecipe(t *testing.T) {
	reader := strings.NewReader(`9
`)

	if FindUltimateHotChocolateRecipe(reader) != "5158916779" {
		t.Error("Wrong score")
	}
}

func TestFindUltimateHotChocolateRecipe2(t *testing.T) {
	reader := strings.NewReader(`5
`)

	if FindUltimateHotChocolateRecipe(reader) != "0124515891" {
		t.Error("Wrong score")
	}
}

func TestFindUltimateHotChocolateRecipe3(t *testing.T) {
	reader := strings.NewReader(`18
`)

	if FindUltimateHotChocolateRecipe(reader) != "9251071085" {
		t.Error("Wrong score")
	}
}

func TestFindUltimateHotChocolateRecipe4(t *testing.T) {
	reader := strings.NewReader(`2018
`)

	if FindUltimateHotChocolateRecipe(reader) != "5941429882" {
		t.Error("Wrong score")
	}
}

func TestFindUltimateHotChocolateRecipeBackward(t *testing.T) {
	reader := strings.NewReader(`51589
`)

	if FindUltimateHotChocolateRecipeBackward(reader) != 9 {
		t.Error("Wrong score")
	}
}

func TestFindUltimateHotChocolateRecipeBackward2(t *testing.T) {
	reader := strings.NewReader(`01245
`)

	if FindUltimateHotChocolateRecipeBackward(reader) != 5 {
		t.Error("Wrong score")
	}
}

func TestFindUltimateHotChocolateRecipeBackward3(t *testing.T) {
	reader := strings.NewReader(`92510
`)

	if FindUltimateHotChocolateRecipeBackward(reader) != 18 {
		t.Error("Wrong score")
	}
}

func TestFindUltimateHotChocolateRecipeBackward4(t *testing.T) {
	reader := strings.NewReader(`59414
`)

	if FindUltimateHotChocolateRecipeBackward(reader) != 2018 {
		t.Error("Wrong score")
	}
}
