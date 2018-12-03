package day_03

import (
	"bufio"
	"fmt"
	"io"
)

type position struct {
	x, y int
}

type fabricClaim struct {
	id, x, y, width, height int
	overlapsWith            map[int]bool
}

func (c *fabricClaim) topRight() position {
	return position{c.x + c.width, c.y}
}

func (c *fabricClaim) bottomLeft() position {
	return position{c.x, c.y + c.height}
}

func newFabricClaim(str string) fabricClaim {
	c := fabricClaim{}
	_, err := fmt.Sscanf(
		str,
		"#%d @ %d,%d: %dx%d",
		&c.id, &c.x, &c.y, &c.width, &c.height,
	)
	if err != nil {
		panic("Invalid line")
	}
	c.overlapsWith = make(map[int]bool)
	return c
}

func CountFabricOverlap(reader io.Reader) int {
	scanner := bufio.NewScanner(reader)
	fabric := make(map[position]int)
	result := 0
	for scanner.Scan() {
		text := scanner.Text()
		claim := newFabricClaim(text)
		for x := claim.x; x < claim.topRight().x; x++ {
			for y := claim.y; y < claim.bottomLeft().y; y++ {
				p := position{x, y}
				fabric[p]++
				if fabric[p] == 2 {
					result++
				}
			}
		}
	}
	return result
}

func FindNotOverlappingFabric(reader io.Reader) int {
	scanner := bufio.NewScanner(reader)
	fabric := make(map[position][]fabricClaim)
	var claims []fabricClaim
	result := 0
	for scanner.Scan() {
		text := scanner.Text()
		claim := newFabricClaim(text)
		claims = append(claims, claim)
		for x := claim.x; x < claim.topRight().x; x++ {
			for y := claim.y; y < claim.bottomLeft().y; y++ {
				p := position{x, y}
				fabric[p] = append(fabric[p], claim)
				if len(fabric[p]) > 1 {
					for _, f := range fabric[p] {
						claim.overlapsWith[f.id] = true
						f.overlapsWith[claim.id] = true
					}
				}
			}
		}
	}
	for _, claim := range claims {
		if len(claim.overlapsWith) == 0 {
			result = claim.id
		}
	}
	return result
}
