package main

import (
	"adventofcode/day_01"
	"adventofcode/day_02"
	"adventofcode/day_03"
	"adventofcode/day_04"
	"adventofcode/day_05"
	"adventofcode/day_06"
	"adventofcode/day_07"
	"adventofcode/day_08"
	"adventofcode/day_09"
	"adventofcode/day_10"
	"adventofcode/day_11"
	"adventofcode/day_12"
	"adventofcode/day_13"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		panic("You must pass 3 arguments")
	}
	day, err := getDay(args)
	part := getPart(err, args)
	file := getFile(err, args)
	defer func() {
		_ = file.Close()
	}()
	var result interface{}
	switch day {
	case 1:
		switch part {
		case 1:
			result = day_01.HandleFrequencyDrift(file)
		case 2:
			result = day_01.FindFirstFrequencyReachedTwice(file)
		}
	case 2:
		switch part {
		case 1:
			result = day_02.CheckSum(file)
		case 2:
			result = day_02.FindTheBoxesFullOfPrototypeFabric(file)
		}
	case 3:
		switch part {
		case 1:
			result = day_03.CountFabricOverlap(file)
		case 2:
			result = day_03.FindNotOverlappingFabric(file)
		}
	case 4:
		switch part {
		case 1:
			result = day_04.FindMostMinuteAsleepOpportunityChecksum(file)
		case 2:
			result = day_04.FindMostFrequentlyAsleepOpportunityChecksum(file)
		}
	case 5:
		switch part {
		case 1:
			result = day_05.SimplifyPolymer(file)
		case 2:
			result = day_05.ImprovePolymer(file)
		}
	case 6:
		switch part {
		case 1:
			result = day_06.FindLargestArea(file)
		case 2:
			result = day_06.FindRegionNearManyCoordinates(file, 10000)
		}
	case 7:
		switch part {
		case 1:
			result = day_07.OrderSteps(file)
		case 2:
			result = day_07.ParallelWorkTime(file, 5, 60)
		}
	case 8:
		switch part {
		case 1:
			result = day_08.SumMetadata(file)
		case 2:
			result = day_08.CalcValueOfNode(file)
		}
	case 9:
		switch part {
		case 1:
			result = day_09.CalcWinningElfsScore(file, 1)
		case 2:
			result = day_09.CalcWinningElfsScore(file, 100)
		}
	case 10:
		switch part {
		case 1:
			result = day_10.AlignTheStars(file)
		case 2:
			result = day_10.CalcTimeForCorrectAlign(file)
		}
	case 11:
		switch part {
		case 1:
			result = day_11.FindTheLargestTotalPowerWithDefaultSize(file)
		case 2:
			result = day_11.FindTheLargestTotalPowerOfAllSize(file)
		}
	case 12:
		switch part {
		case 1:
			result = day_12.SumNumbersOfAllPots(file, 20)
		case 2:
			result = day_12.SumNumbersOfAllPots(file, 50000000000)
		}
	case 13:
		switch part {
		case 1:
			result = day_13.FindLocationOfFirstCrash(file)
		case 2:
			result = day_13.FindLocationOfLastCart(file)
		}
	default:
		panic("Can not find implementation")
	}
	fmt.Printf("%v", result)
}

func getFile(err error, args []string) *os.File {
	path, err := filepath.Abs(args[2])
	if err != nil {
		panic(err)
	}
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return file
}

func getPart(err error, args []string) int {
	part, err := strconv.Atoi(args[1])
	if err != nil {
		panic("Second argument must be a number")
	}
	if part < 1 || part > 2 {
		panic("Second argument must be between 1 and 2")
	}
	return part
}

func getDay(args []string) (int, error) {
	day, err := strconv.Atoi(args[0])
	if err != nil {
		panic("First argument must be a number")
	}
	if day < 1 || day > 25 {
		panic("First argument must be between 1 and 25")
	}
	return day, err
}
