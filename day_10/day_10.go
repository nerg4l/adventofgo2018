package day_10

import (
	"bufio"
	"fmt"
	"io"
	"math"
)

type position struct {
	x, y int
}

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

func parsePoints(reader io.Reader) ([]*point, int, int, int, int) {
	scanner := bufio.NewScanner(reader)
	var points []*point
	var maxX, maxY int
	var minX, minY int
	i := 0
	for scanner.Scan() {
		text := scanner.Text()
		p := &point{id: i}
		i++
		_, _ = fmt.Sscanf(text, "position=<%d,%d> velocity=<%d,%d>", &p.pos.x, &p.pos.y, &p.velocityX, &p.velocityY)
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
		points = append(points, p)
	}
	return points, maxX, maxY, minX, minY
}

func calcTimeForCorrectAlign(maxX int, maxY int, minX int, minY int, points []*point) (int, int, int, int, int) {
	result := 0
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
	return result, maxX, maxY, minX, minY
}

func CalcTimeForCorrectAlign(reader io.Reader) int {
	points, maxX, maxY, minX, minY := parsePoints(reader)
	correctTime, _, _, _, _ := calcTimeForCorrectAlign(maxX, maxY, minX, minY, points)
	return correctTime
}

func AlignTheStars(reader io.ReadSeeker) string {
	points, maxX, maxY, minX, minY := parsePoints(reader)
	runTime, maxX, maxY, minX, minY := calcTimeForCorrectAlign(maxX, maxY, minX, minY, points)
	_, _ = reader.Seek(0, 0)
	points, _, _, _, _ = parsePoints(reader)
	result := ""
	for s := 0; s <= runTime; s++ {
		if s == runTime {
			for y := minY; y <= maxY; y++ {
				line := ""
				for x := minX; x <= maxX; x++ {
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
	return result
}
