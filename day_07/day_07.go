package day_07

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"sort"
)

type step struct {
	letter       rune
	requirements []rune
}

func (t *step) append(r rune) {
	t.requirements = append(t.requirements, r)
}

func (t *step) remove(r rune) {
	for i, s := range t.requirements {
		if s == r {
			t.requirements = append(t.requirements[:i], t.requirements[i+1:]...)
		}
	}
}

func parseSteps(reader io.Reader) map[rune]*step {
	scanner := bufio.NewScanner(reader)
	steps := make(map[rune]*step)
	for scanner.Scan() {
		text := scanner.Text()
		var a, b rune
		_, _ = fmt.Sscanf(text, "Step %c must be finished before step %c can begin.", &a, &b)
		if _, ok := steps[b]; !ok {
			steps[b] = &step{letter: b}
		}
		if _, ok := steps[a]; !ok {
			steps[a] = &step{letter: a}
		}
		steps[b].append(a)
	}
	return steps
}

func OrderSteps(reader io.Reader) string {
	steps := parseSteps(reader)
	result := ""
	for len(steps) > 0 {
		copied := make([]*step, 0, len(steps))
		for _, value := range steps {
			copied = append(copied, value)
		}
		sort.Slice(copied, func(i, j int) bool {
			a := len(copied[i].requirements)*math.MaxInt32 + int(copied[i].letter)
			b := len(copied[j].requirements)*math.MaxInt32 + int(copied[j].letter)
			return a < b
		})
		next := copied[0].letter
		result += string(next)
		for _, s := range steps {
			s.remove(next)
		}
		delete(steps, next)
	}
	return result
}

var currentTime = 0

type worker struct {
	doing     *step
	available int
}

func (w *worker) isAvailable() bool {
	return w.doing == nil
}

func (w *worker) willBeAvailable() bool {
	return w.available < currentTime
}

func (w *worker) applyJob(s *step, stepTime int) {
	w.doing = s
	w.available = currentTime + stepTime + int(s.letter-'A')
}

func ParallelWorkTime(reader io.Reader, workersN int, stepTime int) int {
	steps := parseSteps(reader)
	var workers []*worker
	for i := 0; i < workersN; i++ {
		workers = append(workers, &worker{})
	}
	for ; true; currentTime++ {
		for _, w := range workers {
			if w.willBeAvailable() && w.doing != nil {
				for _, s := range steps {
					s.remove(w.doing.letter)
				}
				w.doing = nil
			}
		}
		if len(steps) == 0 {
			working := false
			for _, w := range workers {
				if w.doing != nil {
					working = true
					break
				}
			}
			if !working {
				break
			}
		}
		copied := make([]*step, 0, len(steps))
		for _, value := range steps {
			copied = append(copied, value)
		}
		sort.Slice(copied, func(i, j int) bool {
			a := len(copied[i].requirements)*math.MaxInt32 + int(copied[i].letter)
			b := len(copied[j].requirements)*math.MaxInt32 + int(copied[j].letter)
			return a < b
		})
		for i := 0; i < len(copied); i++ {
			c := copied[i]
			if len(c.requirements) > 0 {
				break
			}
			for _, w := range workers {
				if w.isAvailable() {
					w.applyJob(c, stepTime)
					delete(steps, c.letter)
					break
				}
			}
		}
	}
	return currentTime
}
