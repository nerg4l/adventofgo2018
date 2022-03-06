package main

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"sort"
	"time"
)

const dateFormat = "2006-01-02 15:04"

type sleepLog struct {
	t    time.Time
	text string
}

func (l sleepLog) Before(ll sleepLog) bool {
	return l.t.Before(ll.t)
}

func parseSleepLog(r io.Reader) ([]sleepLog, error) {
	s := bufio.NewScanner(r)
	exp := regexp.MustCompile(`\[(.*?)] (.*)`)
	var logs []sleepLog
	for s.Scan() {
		text := s.Text()
		matches := exp.FindStringSubmatch(text)
		l := sleepLog{}
		var err error
		if l.t, err = time.Parse(dateFormat, matches[1]); err != nil {
			return nil, err
		}
		l.text = matches[2]
		logs = append(logs, l)
	}
	return logs, s.Err()
}

type sleepyGuard struct {
	id    int
	sleep [60]int
}

func (g sleepyGuard) sleepSum() int {
	n := 0
	for i := range g.sleep {
		n += g.sleep[i]
	}
	return n
}

func (g sleepyGuard) sleepiestMinute() int {
	var max, m int
	for i, v := range g.sleep {
		if v > max {
			max = v
			m = i
		}
	}
	return m
}

func parseSleepyGuards(records []sleepLog) (map[int]sleepyGuard, error) {
	guards := make(map[int]sleepyGuard)
	var g sleepyGuard
	var start time.Time
	for _, record := range records {
		switch record.text {
		case "falls asleep":
			start = record.t
		case "wakes up":
			t := start
			for t.Before(record.t) {
				g.sleep[t.Minute()]++
				t = t.Add(time.Minute)
			}
			guards[g.id] = g
		default:
			var id int
			_, err := fmt.Sscanf(record.text, "Guard #%d begins shift", &id)
			if err != nil {
				return nil, err
			}
			var ok bool
			g, ok = guards[id]
			if !ok {
				g = sleepyGuard{id: id}
				guards[id] = g
			}
		}
	}
	return guards, nil
}

func MostMinuteAsleepOpportunityChecksum(r io.Reader) (interface{}, error) {
	logs, err := parseSleepLog(r)
	if err != nil {
		return nil, err
	}
	sort.Slice(logs, func(i, j int) bool {
		return logs[i].t.Before(logs[j].t)
	})
	guards, err := parseSleepyGuards(logs)
	if err != nil {
		return nil, err
	}
	var selected sleepyGuard
	for i, v := range guards {
		if i == 0 || selected.sleepSum() < v.sleepSum() {
			selected = v
		}
	}
	min := selected.sleepiestMinute()
	return selected.id * min, nil
}

func MostFrequentlyAsleepOpportunityChecksum(r io.Reader) (interface{}, error) {
	logs, err := parseSleepLog(r)
	if err != nil {
		return nil, err
	}
	sort.Slice(logs, func(i, j int) bool {
		return logs[i].Before(logs[j])
	})
	guards, err := parseSleepyGuards(logs)
	if err != nil {
		return nil, err
	}
	var selected sleepyGuard
	for _, v := range guards {
		if selected.sleep[selected.sleepiestMinute()] < v.sleep[v.sleepiestMinute()] {
			selected = v
		}
	}
	min := selected.sleepiestMinute()
	return selected.id * min, nil
}
