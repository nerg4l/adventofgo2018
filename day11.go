package main

import (
	"fmt"
	"io"
)

type FuelGrid [300][300]*FuelCell

func NewGrid(r io.Reader) (FuelGrid, error) {
	var serialNumber int
	_, err := fmt.Fscanf(r, "%d", &serialNumber)
	if err != nil {
		return FuelGrid{}, err
	}
	var grid FuelGrid
	for x := 1; x <= 300; x++ {
		for y := 1; y <= 300; y++ {
			xIndex, yIndex := x-1, y-1
			grid[xIndex][yIndex] = &FuelCell{x, y, serialNumber}
		}
	}
	return grid, nil
}

type FuelCell struct {
	x, y         int
	serialNumber int
}

func (f *FuelCell) RackId() int {
	return f.x + 10
}

func (f *FuelCell) PowerLevel() int {
	begin := f.RackId() * f.y
	increased := begin + f.serialNumber
	hundredsDigit := (increased * f.RackId() / 100) % 10
	return hundredsDigit - 5
}

func LargestTotalPowerWithDefaultSize(r io.Reader) (interface{}, error) {
	grid, err := NewGrid(r)
	if err != nil {
		return nil, err
	}
	size := 3
	var maxPowerLevel int
	var coordinates string
	for x := 1; x <= 300-size; x++ {
		for y := 1; y <= 300-size; y++ {
			xIndex, yIndex := x-1, y-1
			var sumPowerLevel int
			for xMod := 0; xMod < size; xMod++ {
				for yMod := 0; yMod < size; yMod++ {
					sumPowerLevel += grid[xIndex+xMod][yIndex+yMod].PowerLevel()
				}
			}
			if sumPowerLevel > maxPowerLevel {
				maxPowerLevel = sumPowerLevel
				coordinates = fmt.Sprintf("%d,%d", x, y)
			}
		}
	}
	return coordinates, nil
}

func LargestTotalPowerOfAllSize(r io.Reader) (interface{}, error) {
	grid, err := NewGrid(r)
	if err != nil {
		return nil, err
	}
	var maxPowerLevel int
	var size int
	var coordinates string
	for x := 1; x <= 300; x++ {
		for y := 1; y <= 300; y++ {
			xIndex, yIndex := x-1, y-1
			var sumPowerLevel int
			maxSize := 300 - max(x, y)
			for s := 1; s < maxSize; s++ {
				for xMod := 0; xMod < s; xMod++ {
					sumPowerLevel += grid[xIndex+xMod][yIndex+s-1].PowerLevel()
				}
				for yMod := 0; yMod < s-1; yMod++ {
					sumPowerLevel += grid[xIndex+s-1][yIndex+yMod].PowerLevel()
				}
				if sumPowerLevel > maxPowerLevel {
					maxPowerLevel = sumPowerLevel
					coordinates = fmt.Sprintf("%d,%d", x, y)
					size = s
				}
			}
		}
	}
	return fmt.Sprintf("%s,%d", coordinates, size), nil
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
