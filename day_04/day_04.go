package day_04

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"sort"
	"time"
)

const dateFormat = "2006-01-02 15:04"

type log struct {
	time time.Time
	text string
}

type guard struct {
	id        int
	sleep     [60]int
	sleepTime int
}

func (g *guard) addSleep(start, end int) {
	g.sleepTime += end - start
	for start < end {
		g.sleep[start]++
		start += 1
	}
}

func (g *guard) sleepiest() (min, duration int) {
	for i, v := range g.sleep {
		if v > duration {
			duration = v
			min = i
		}
	}
	return min, duration
}

func parseRecords(reader io.Reader) []log {
	scanner := bufio.NewScanner(reader)
	r := regexp.MustCompile(`\[(.*?)\] (.*)`)
	var logs []log
	for scanner.Scan() {
		text := scanner.Text()
		matches := r.FindStringSubmatch(text)
		l := log{}
		var err error
		l.time, err = time.Parse(dateFormat, matches[1])
		if err != nil {
			panic(err)
		}
		l.text = matches[2]
		logs = append(logs, l)
	}
	sort.Slice(logs, func(i, j int) bool {
		return logs[i].time.Before(logs[j].time)
	})
	return logs
}

func parseGuards(records []log) map[int]*guard {
	guards := make(map[int]*guard)
	var g *guard
	var start int
	for _, record := range records {
		switch record.text {
		case "falls asleep":
			start = record.time.Minute()
		case "wakes up":
			g.addSleep(start, record.time.Minute())
		default:
			var id int
			_, _ = fmt.Sscanf(record.text, "Guard #%d begins shift", &id)
			if _, ok := guards[id]; !ok {
				guards[id] = &guard{id: id}
			}
			g = guards[id]
		}
	}
	return guards
}

func FindMostMinuteAsleepOpportunityChecksum(reader io.Reader) int {
	guards := parseGuards(parseRecords(reader))
	var selected *guard
	for _, v := range guards {
		if selected == nil || selected.sleepTime < v.sleepTime {
			selected = v
		}
	}
	min, _ := selected.sleepiest()
	return selected.id * min
}

func FindMostFrequentlyAsleepOpportunityChecksum(reader io.Reader) int {
	guards := parseGuards(parseRecords(reader))
	var selected *guard
	for _, v := range guards {
		if selected == nil {
			selected = v
			continue
		}
		_, durSelected := selected.sleepiest()
		_, dur := v.sleepiest()
		if durSelected < dur {
			selected = v
		}
	}
	min, _ := selected.sleepiest()
	return selected.id * min
}
