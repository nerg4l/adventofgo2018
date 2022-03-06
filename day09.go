package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"io"
)

func marbleGame(playerN, lastMarble int) int {
	scores := make([]int, playerN)
	r := ring.New(1)
	r.Value = 0

	for i := 1; i <= lastMarble; i++ {
		if i%23 == 0 {
			for j := 0; j < 8; j++ {
				r = r.Prev()
			}

			u := r.Unlink(1)
			scores[i%playerN] += i + u.Value.(int)
			r = r.Next()
			continue
		}
		rr := &ring.Ring{Value: i}
		r.Next().Link(rr)
		r = rr
	}

	max := 0
	for _, v := range scores {
		max = Max(max, v)
	}
	return int(max)
}

func WinningElfsScore(r io.Reader, multiplier uint) (interface{}, error) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		text := s.Text()
		var playerN, lastMarble int
		_, err := fmt.Sscanf(text, "%d players; last marble is worth %d points", &playerN, &lastMarble)
		if err != nil {
			return nil, err
		}
		return marbleGame(playerN, lastMarble*int(multiplier)), nil
	}
	return nil, s.Err()
}
