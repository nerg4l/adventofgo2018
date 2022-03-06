package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math"
)

type point struct {
	id                   int
	pos                  position
	velocityX, velocityY int
}

func (p *point) tick() {
	p.pos.x += p.velocityX
	p.pos.y += p.velocityY
}

func contains(haystack []*point, needle position) bool {
	for _, v := range haystack {
		if v.pos == needle {
			return true
		}
	}
	return false
}

func parsePoints(r io.Reader) (ps []*point, max, min position, err error) {
	s := bufio.NewScanner(r)
	var points []*point
	i := 0
	for s.Scan() {
		text := s.Text()
		p := &point{id: i}
		i++
		_, err := fmt.Sscanf(text, "position=<%d,%d> velocity=<%d,%d>", &p.pos.x, &p.pos.y, &p.velocityX, &p.velocityY)
		if err != nil {
			return nil, position{}, position{}, err
		}
		if p.pos.x > max.x {
			max.x = p.pos.x
		}
		if p.pos.y > max.y {
			max.y = p.pos.y
		}
		if p.pos.x < min.x {
			min.x = p.pos.x
		}
		if p.pos.y < min.y {
			min.y = p.pos.y
		}
		points = append(points, p)
	}
	return points, max, min, s.Err()
}

func timeForCorrectAlign(max, min position, points []*point) (int, position, position) {
	result := 0
	maxX, maxY, minX, minY := max.x, max.y, min.x, min.y
	for s := 0; true; s++ {
		oMaxX, oMaxY, oMinX, oMinY := maxX, maxY, minX, minY
		maxX, maxY, minX, minY = -math.MaxInt32, -math.MaxInt32, math.MaxInt32, math.MaxInt32
		for _, p := range points {
			p.tick()
			if p.pos.x > maxX {
				maxX = p.pos.x
			}
			if p.pos.y > maxY {
				maxY = p.pos.y
			}
			if p.pos.x < minX {
				minX = p.pos.x
			}
			if p.pos.y < minY {
				minY = p.pos.y
			}
		}
		if oMaxX < maxX || oMaxY < maxY || oMinX > minX || oMinY > minY {
			result = s
			maxX, maxY, minX, minY = oMaxX, oMaxY, oMinX, oMinY
			break
		}
	}
	return result, position{x: maxX, y: maxY}, position{x: minX, y: minY}
}

func TimeForCorrectAlign(r io.Reader) (interface{}, error) {
	points, max, min, err := parsePoints(r)
	if err != nil {
		return nil, err
	}
	correctTime, _, _ := timeForCorrectAlign(max, min, points)
	return correctTime, nil
}

func AlignTheStars(r io.Reader) (interface{}, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	rs := bytes.NewReader(b)
	points, max, min, err := parsePoints(rs)
	if err != nil {
		return nil, err
	}
	runTime, max, min := timeForCorrectAlign(max, min, points)
	_, _ = rs.Seek(0, 0)
	points, _, _, _ = parsePoints(rs)
	result := ""
	for s := 0; s <= runTime; s++ {
		if s == runTime {
			for y := min.y; y <= max.y; y++ {
				line := ""
				for x := min.x; x <= max.x; x++ {
					pos := position{x, y}
					if contains(points, pos) {
						line += "#"
					} else {
						line += "."
					}
				}
				result += fmt.Sprintln(line)
			}
			break
		}
		for _, p := range points {
			p.tick()
		}
	}
	return result, nil
}
