package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("You must pass 2 arguments")
		os.Exit(1)
	}
	d, err := day(args)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	p, err := part(args)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	i := ((d - 1) * 2) + (p - 1)
	if i >= len(parts) {
		fmt.Println("Could not find implementation")
		os.Exit(1)
	}
	f := parts[i]
	if f == nil {
		fmt.Println("Could not find implementation")
		os.Exit(1)
	}
	result, err := f(os.Stdin)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	fmt.Printf("%v\n", result)
}

func part(args []string) (int, error) {
	part, err := strconv.Atoi(args[1])
	if err != nil {
		return 0, err
	}
	if part < 1 || part > 2 {
		return 0, errors.New("second argument must be between 1 and 2")
	}
	return part, nil
}

func day(args []string) (int, error) {
	day, err := strconv.Atoi(args[0])
	if err != nil {
		return 0, err
	}
	if day < 1 || day > 25 {
		return 0, errors.New("first argument must be between 1 and 25")
	}
	return day, nil
}

var parts = [](func(r io.Reader) (interface{}, error)){
	CalibrateResultingFrequency,
	CalibrateFirstRepeatingFrequency,
	BoxIDChecksum,
	BoxIDCommonLetters,
	CountFabricOverlap,
	NotOverlappingFabric,
	MostMinuteAsleepOpportunityChecksum,
	MostFrequentlyAsleepOpportunityChecksum,
	SimplifyPolymer,
	ImprovePolymer,
	LargestArea,
	func(r io.Reader) (interface{}, error) {
		return RegionNearManyCoordinates(r, 10000)
	},
	OrderSteps,
	func(r io.Reader) (interface{}, error) {
		return ParallelWorkTime(r, 5, 60)
	},
	SumMetadata,
	ValueOfNode,
	func(r io.Reader) (interface{}, error) {
		return WinningElfsScore(r, 1)
	},
	func(r io.Reader) (interface{}, error) {
		return WinningElfsScore(r, 100)
	},
	AlignTheStars,
	TimeForCorrectAlign,
	LargestTotalPowerWithDefaultSize,
	LargestTotalPowerOfAllSize,
	func(r io.Reader) (interface{}, error) {
		return SumNumbersOfAllPots(r, 20)
	},
	func(r io.Reader) (interface{}, error) {
		return SumNumbersOfAllPots(r, 50000000000)
	},
	LocationOfFirstCrash,
	LocationOfLastCart,
	UltimateHotChocolateRecipe,
	UltimateHotChocolateRecipeBackward,
}
