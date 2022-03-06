package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
)

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// rectangle represents the coordinates of the opposite
// corners a rectangle
type rectangle struct {
	a, b position
}

func parseDestinations(r io.Reader) (rectangle, []position, error) {
	var positions []position
	var rect rectangle
	s := bufio.NewScanner(r)
	rect.a.x, rect.a.y = math.MaxInt, math.MaxInt
	rect.b.x, rect.b.y = 0, 0
	for s.Scan() {
		text := s.Text()
		var x, y int
		_, err := fmt.Sscanf(text, "%d, %d", &x, &y)
		if err != nil {
			return rectangle{}, nil, err
		}
		positions = append(positions, position{x, y})
		rect.a.x, rect.a.y = Min(rect.a.x, x), Min(rect.a.y, y)
		rect.b.x, rect.b.y = Max(rect.b.x, x), Max(rect.b.y, y)
	}
	return rect, positions, s.Err()
}

func manhattanDistance(a, b position) int {
	return int(math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y)))
}

func LargestArea(r io.Reader) (interface{}, error) {
	rect, positions, err := parseDestinations(r)
	if err != nil {
		return nil, err
	}
	areas := make(map[position]float64, len(positions))
	xx, yy := Min(rect.a.x, rect.b.x), Min(rect.a.y, rect.b.y)
	xn, ym := Max(rect.a.x, rect.b.x), Max(rect.a.y, rect.b.y)
	for x := xx; x <= xn; x++ {
		for y := yy; y <= ym; y++ {
			minmd := math.MaxInt
			p := position{x, y}
			index := -1
			for i := range positions {
				md := manhattanDistance(positions[i], p)
				if md < minmd {
					minmd = md
					index = i
				} else if md == minmd {
					index = -1
				}
			}
			if index < 0 {
				continue
			}
			aindex := positions[index]
			if p.x == xx || p.y == yy || p.x == xn || p.y == ym {
				areas[aindex] = math.Inf(-1)
				continue
			}
			areas[aindex]++
		}
	}
	var max float64
	for _, a := range areas {
		max = math.Max(max, a)
	}
	return int(max), nil
}

func RegionNearManyCoordinates(r io.Reader, lessThan int) (interface{}, error) {
	rect, positions, err := parseDestinations(r)
	if err != nil {
		return nil, err
	}
	var result int
	xx, yy := Min(rect.a.x, rect.b.x), Min(rect.a.y, rect.b.y)
	xn, ym := Max(rect.a.x, rect.b.x), Max(rect.a.y, rect.b.y)
	for x := xx; x <= xn; x++ {
		for y := yy; y <= ym; y++ {
			c := position{x, y}
			var totalmd int
			for _, v := range positions {
				totalmd += manhattanDistance(v, c)
			}
			if totalmd < lessThan {
				result++
			}
		}
	}
	return result, nil
}
