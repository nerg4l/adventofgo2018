package day_09

import (
	"bufio"
	"container/ring"
	"fmt"
	"io"
	"math"
)

func marbleGame(players, lastMarble int) int {
	scores := make([]float64, players)
	r := ring.New(1)
	r.Value = 0

	for i := 1; i <= lastMarble; i++ {
		if i%23 == 0 {
			for j := 0; j < 8; j++ {
				r = r.Prev()
			}

			u := r.Unlink(1)
			scores[i%players] += float64(i + u.Value.(int))
			r = r.Next()
		} else {
			tmp := &ring.Ring{Value: i}
			r.Next().Link(tmp)
			r = tmp
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
