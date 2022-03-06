package main

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
	"time"
)

type sleightAssemblyStep struct {
	letter         byte
	followingSteps map[byte]struct{}
	preSteps       map[byte]struct{}
}

func newSleightAssemblyStep(letter byte) sleightAssemblyStep {
	return sleightAssemblyStep{
		letter:         letter,
		followingSteps: make(map[byte]struct{}),
		preSteps:       make(map[byte]struct{}),
	}
}

func parseSteps(r io.Reader) (map[byte]sleightAssemblyStep, error) {
	s := bufio.NewScanner(r)
	steps := make(map[byte]sleightAssemblyStep)
	for s.Scan() {
		text := s.Text()
		var a, b byte
		_, err := fmt.Sscanf(text, "Step %c must be finished before step %c can begin.", &a, &b)
		if err != nil {
			return nil, err
		}
		if _, ok := steps[a]; !ok {
			steps[a] = newSleightAssemblyStep(a)
		}
		steps[a].followingSteps[b] = struct{}{}
		if _, ok := steps[b]; !ok {
			steps[b] = newSleightAssemblyStep(b)
		}
		steps[b].preSteps[a] = struct{}{}
	}
	return steps, s.Err()
}

func OrderSteps(r io.Reader) (interface{}, error) {
	lookup, err := parseSteps(r)
	if err != nil {
		return nil, err
	}
	steps := make([]sleightAssemblyStep, 0, len(lookup))
	for _, value := range lookup {
		steps = append(steps, value)
	}
	var result strings.Builder
	for len(steps) > 0 {
		sort.Slice(steps, func(i, j int) bool {
			a := (len(steps[i].preSteps) << 8) + int(steps[i].letter)
			b := (len(steps[j].preSteps) << 8) + int(steps[j].letter)
			return a < b
		})
		next := steps[0].letter
		result.WriteByte(next)
		for following := range lookup[next].followingSteps {
			delete(lookup[following].preSteps, next)
		}
		steps = steps[1:]
	}
	return result.String(), nil
}

type worker struct {
	letter    byte
	busyUntil time.Duration
}

func (w *worker) busy(d time.Duration) bool {
	return w.busyUntil > d
}

type workerSchedule struct {
	available map[*worker]struct{}
	busy      map[*worker]struct{}
}

func (s workerSchedule) free(w *worker) {
	s.available[w] = struct{}{}
	delete(s.busy, w)
}

func (s workerSchedule) occupy(w *worker, letter byte, until time.Duration) {
	w.letter = letter
	w.busyUntil = until
	s.busy[w] = struct{}{}
	delete(s.available, w)
}

func ParallelWorkTime(r io.Reader, workersN int, stepTime int) (interface{}, error) {
	lookup, err := parseSteps(r)
	if err != nil {
		return nil, err
	}
	schedule := workerSchedule{
		available: make(map[*worker]struct{}, workersN),
		busy:      make(map[*worker]struct{}, workersN),
	}
	for i := 0; i < workersN; i++ {
		schedule.available[&worker{}] = struct{}{}
	}
	steps := make([]sleightAssemblyStep, 0, len(lookup))
	for _, value := range lookup {
		steps = append(steps, value)
	}
	duration := time.Duration(0)
	for ; len(steps) > 0; duration += time.Second {
		for w := range schedule.busy {
			if !w.busy(duration) {
				for following := range lookup[w.letter].followingSteps {
					delete(lookup[following].preSteps, w.letter)
				}
				schedule.free(w)
			}
		}
		if len(schedule.available) == 0 {
			continue
		}
		sort.Slice(steps, func(i, j int) bool {
			a := (len(steps[i].preSteps) << 8) + int(steps[i].letter)
			b := (len(steps[j].preSteps) << 8) + int(steps[j].letter)
			return a < b
		})
		for len(steps) > 0 {
			c := steps[0]
			if len(c.preSteps) > 0 || len(schedule.available) == 0 {
				break
			}
			for w := range schedule.available {
				if w.busy(duration) {
					continue
				}
				steps = steps[1:]
				schedule.occupy(w, c.letter, duration+(time.Duration(stepTime)+time.Duration(c.letter-'A'+1))*time.Second)
				break
			}
		}
	}
	// Shortcut for jumping to last second.
	max := duration
	for w := range schedule.busy {
		if max < w.busyUntil {
			max = w.busyUntil
		}
	}
	return int(max / time.Second), nil
}
