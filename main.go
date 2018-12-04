package main

import (
	"adventofcode/day_01"
	"adventofcode/day_02"
	"adventofcode/day_03"
	"adventofcode/day_04"
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
