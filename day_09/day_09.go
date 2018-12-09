package day_09

import (
	"bufio"
	"fmt"
	"io"
	"math"
)

type marble struct {
	val        int
	prev, next *marble
}

func marbleGame(players, lastMarble int) int {
	scores := make([]float64, players)
	c := &marble{val: 0}
	c.next = c
	c.prev = c

	for i := 1; i <= lastMarble; i++ {
		if i%23 == 0 {
			for j := 0; j < 7; j++ {
				c = c.prev
			}
			c.prev.next = c.next
			c.next.prev = c.prev
			scores[i%players] += float64(i + c.val)
			c = c.next
		} else {
			tmp := &marble{val: i}
			tmp.prev = c.next
			tmp.next = c.next.next
			tmp.prev.next = tmp
			tmp.next.prev = tmp
			c = tmp
		}
	}

	max := 0.0
	for _, v := range scores {
		max = math.Max(max, v)
	}
	return int(max)
}

func CalcWinningElfsScore(reader io.Reader, multiplier uint) int {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		text := scanner.Text()
		var players, lastMarble int
		_, _ = fmt.Sscanf(text, "%d players; last marble is worth %d points", &players, &lastMarble)
		return marbleGame(players, lastMarble*int(multiplier))
	}
	return 0
}
