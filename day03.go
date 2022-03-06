package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
)

type position struct {
	x, y int
}

type fabricClaim struct {
	id, width, height int
	p                 position
}

func (c fabricClaim) X() int {
	return c.p.x
}

func (c fabricClaim) Y() int {
	return c.p.y
}

func fabricClaimFromString(str string) (fabricClaim, error) {
	c := fabricClaim{}
	_, err := fmt.Sscanf(
		str,
		"#%d @ %d,%d: %dx%d",
		&c.id, &c.p.x, &c.p.y, &c.width, &c.height,
	)
	if err != nil {
		return fabricClaim{}, err
	}
	return c, nil
}

func CountFabricOverlap(r io.Reader) (interface{}, error) {
	s := bufio.NewScanner(r)
	fabric := make(map[position][]fabricClaim)
	result := 0
	for s.Scan() {
		text := s.Text()
		claim, err := fabricClaimFromString(text)
		if err != nil {
			return nil, err
		}
		xx, yy := claim.X()+claim.width, claim.Y()+claim.height
		for x := claim.X(); x < xx; x++ {
			for y := claim.Y(); y < yy; y++ {
				p := position{x, y}
				fabric[p] = append(fabric[p], claim)
				if len(fabric[p]) == 2 {
					result++
				}
			}
		}
	}
	return result, s.Err()
}

func NotOverlappingFabric(r io.Reader) (interface{}, error) {
	s := bufio.NewScanner(r)
	claims := make(map[position][]fabricClaim)
	intact := make(map[int]bool)
	for s.Scan() {
		text := s.Text()
		claim, err := fabricClaimFromString(text)
		intact[claim.id] = true
		if err != nil {
			return nil, err
		}
		n, m := claim.X()+claim.width, claim.Y()+claim.height
		for x := claim.X(); x < n; x++ {
			for y := claim.Y(); y < m; y++ {
				p := position{x, y}
				claims[p] = append(claims[p], claim)
				if len(claims[p]) > 1 {
					for _, c := range claims[p] {
						delete(intact, c.id)
					}
				}
			}
		}
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	if len(intact) > 1 {
		return nil, errors.New("more than one found")
	}
	for id := range intact {
		return id, nil
	}
	return nil, errors.New("not found")
}
