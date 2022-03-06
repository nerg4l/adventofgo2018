package main

import (
	"bufio"
	"bytes"
	"container/ring"
	"fmt"
	"io"
)

const (
	left option = iota
	straight
	right
)

const (
	directionL = '<'
	directionR = '>'
	directionU = '^'
	directionD = 'v'
)

type option int

type cart struct {
	orientation *ring.Ring
	position
	option
}

func (c *cart) intersection() {
	switch c.option {
	case left:
		c.orientation = c.orientation.Prev()
		c.option = straight
	case straight:
		c.option = right
	case right:
		c.orientation = c.orientation.Next()
		c.option = left
	}
}

func (c *cart) move(r rune) {
	switch {
	case c.orientation.Value == directionR && r == '\\':
		c.orientation = c.orientation.Next()
	case c.orientation.Value == directionR && r == '/':
		c.orientation = c.orientation.Prev()
	case c.orientation.Value == directionD && r == '/':
		c.orientation = c.orientation.Next()
	case c.orientation.Value == directionD && r == '\\':
		c.orientation = c.orientation.Prev()
	case c.orientation.Value == directionL && r == '\\':
		c.orientation = c.orientation.Next()
	case c.orientation.Value == directionL && r == '/':
		c.orientation = c.orientation.Prev()
	case c.orientation.Value == directionU && r == '/':
		c.orientation = c.orientation.Next()
	case c.orientation.Value == directionU && r == '\\':
		c.orientation = c.orientation.Prev()
	case r == '+':
		c.intersection()
	}
}

func (c *cart) nextPosition() position {
	switch c.orientation.Value {
	case directionR:
		c.position.x++
	case directionL:
		c.position.x--
	case directionD:
		c.position.y++
	case directionU:
		c.position.y--
	}
	return c.position
}

func newOrientation(original rune) *ring.Ring {
	r := ring.New(4)
	r.Value = directionR
	r = r.Next()
	r.Value = directionD
	r = r.Next()
	r.Value = directionL
	r = r.Next()
	r.Value = directionU
	for original != r.Value {
		r = r.Next()
	}
	return r
}

func parseMap(r io.Reader) (map[position]*cart, [][]rune, error) {
	s := bufio.NewScanner(r)
	carts := make(map[position]*cart)
	var track [][]rune
	for s.Scan() {
		text := s.Text()

		y := len(track)
		var lineBuffer bytes.Buffer
		for x, r := range text {
			switch r {
			case directionR, directionD, directionL, directionU:
				p := position{x, y}
				carts[p] = &cart{
					orientation: newOrientation(r),
					position:    p,
					option:      left,
				}
			}
			switch r {
			case directionR, directionL:
				lineBuffer.WriteRune('-')
			case directionD, directionU:
				lineBuffer.WriteRune('|')
			default:
				lineBuffer.WriteRune(r)
			}
		}

		track = append(track, []rune(lineBuffer.String()))
	}
	return carts, track, s.Err()
}

func LocationOfFirstCrash(r io.Reader) (interface{}, error) {
	carts, track, err := parseMap(r)
	if err != nil {
		return nil, err
	}
	for {
		cartsCopy := make(map[position]*cart)
		for k, v := range carts {
			cartsCopy[k] = v
		}
		for y, line := range track {
			for x := range line {
				p := position{x, y}
				var c *cart
				var ok bool
				if c, ok = carts[p]; !ok {
					continue
				}
				next := c.nextPosition()
				if _, ok = cartsCopy[next]; ok {
					return fmt.Sprintf("%d,%d", next.x, next.y), nil
				}
				c.move(track[next.y][next.x])
				delete(cartsCopy, p)
				cartsCopy[next] = c
			}
		}
		carts = cartsCopy
	}
}

func LocationOfLastCart(r io.Reader) (interface{}, error) {
	carts, track, err := parseMap(r)
	if err != nil {
		return nil, err
	}
	for {
		cartsCopy := make(map[position]*cart)
		for k, v := range carts {
			cartsCopy[k] = v
		}
		for y, line := range track {
			for x := range line {
				p := position{x, y}
				var c *cart
				var ok bool
				if c, ok = carts[p]; !ok {
					continue
				}
				next := c.nextPosition()
				if _, ok = cartsCopy[next]; ok {
					delete(cartsCopy, next)
					delete(cartsCopy, p)
					delete(carts, next)
					continue
				}
				c.move(track[next.y][next.x])
				delete(cartsCopy, p)
				cartsCopy[next] = c
			}
		}
		carts = cartsCopy
		if len(carts) == 1 {
			for _, v := range carts {
				return fmt.Sprintf("%d,%d", v.position.x, v.position.y), nil
			}
		}
	}
}
