package day_06

import (
	"bufio"
	"io"
	"math"
	"strconv"
	"strings"
)

func parseDeviceData(reader io.Reader) deviceData {
	dd := deviceData{}
	scanner := bufio.NewScanner(reader)
	dd.minX, dd.minY = math.MaxFloat32, math.MaxFloat32
	dd.maxX, dd.maxY = .0, .0
	for scanner.Scan() {
		text := scanner.Text()
		r := strings.SplitN(text, ", ", 2)
		x, _ := strconv.ParseFloat(r[0], 32)
		y, _ := strconv.ParseFloat(r[1], 32)
		dd.coordinates = append(dd.coordinates, coordinate{x, y})
		dd.minX = math.Min(dd.minX, x)
		dd.minY = math.Min(dd.minY, y)
		dd.maxX = math.Max(dd.maxX, x)
		dd.maxY = math.Max(dd.maxY, y)
	}
	return dd
}

type deviceData struct {
	minX, minY  float64
	maxX, maxY  float64
	coordinates []coordinate
}

type coordinate struct {
	x, y float64
}

func (a coordinate) manhattan(b coordinate) float64 {
	return math.Abs(a.x-b.x) + math.Abs(a.y-b.y)
}

func FindLargestArea(reader io.Reader) int {
	dd := parseDeviceData(reader)
	fieldClosestTo := make(map[coordinate]int)
	for x := dd.minX; x <= dd.maxX; x++ {
		for y := dd.minY; y <= dd.maxY; y++ {
			min := dd.maxX + dd.maxY
			c := coordinate{x, y}
			for i, v := range dd.coordinates {
				manhattan := v.manhattan(c)
				if manhattan < min {
					min = manhattan
					fieldClosestTo[c] = i
				} else if manhattan == min {
					fieldClosestTo[c] = -1
				}
			}
		}
	}
	var max float64
	for ddIndex, ddC := range dd.coordinates {
		if ddC.x == dd.minX || ddC.y == dd.minY || ddC.x == dd.maxX || ddC.y == dd.maxY {
			continue
		}
		var m float64
		for fieldC, index := range fieldClosestTo {
			if ddIndex == index {
				if fieldC.x == dd.minX || fieldC.y == dd.minY || fieldC.x == dd.maxX || fieldC.y == dd.maxY {
					m = 0
					break
				}
				m++
			}
		}
		max = math.Max(max, m)
	}
	return int(max)
}

func FindRegionNearManyCoordinates(reader io.Reader, lessThan float64) int {
	dd := parseDeviceData(reader)
	field := make(map[coordinate]float64)
	for x := dd.minX; x <= dd.maxX; x++ {
		for y := dd.minY; y <= dd.maxY; y++ {
			c := coordinate{x, y}
			var totalManhattan float64
			for _, v := range dd.coordinates {
				totalManhattan += v.manhattan(c)
			}
			field[c] = totalManhattan
		}
	}
	var result int
	for _, totalManhattan := range field {
		if totalManhattan < lessThan {
			result++
		}
	}
	return result
}
