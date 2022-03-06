package main

import (
	"strings"
	"testing"
)

func TestOrderSteps(t *testing.T) {
	r := strings.NewReader(`Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.`)
	if got, _ := OrderSteps(r); got != "CABDFE" {
		t.Error("Wrong order")
	}
}

func TestParallelWorkTime(t *testing.T) {
	r := strings.NewReader(`Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.`)
	if got, _ := ParallelWorkTime(r, 2, 0); got != 15 {
		t.Error("Wrong order")
	}
}
