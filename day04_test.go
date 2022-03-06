package main

import (
	"strings"
	"testing"
)

func TestMostMinuteAsleepOpportunityChecksum(t *testing.T) {
	r := strings.NewReader(`[1518-11-01 00:00] Guard #10 begins shift
[1518-11-01 00:05] falls asleep
[1518-11-01 00:25] wakes up
[1518-11-01 00:30] falls asleep
[1518-11-01 00:55] wakes up
[1518-11-01 23:58] Guard #99 begins shift
[1518-11-02 00:40] falls asleep
[1518-11-02 00:50] wakes up
[1518-11-03 00:05] Guard #10 begins shift
[1518-11-03 00:24] falls asleep
[1518-11-03 00:29] wakes up
[1518-11-04 00:02] Guard #99 begins shift
[1518-11-04 00:36] falls asleep
[1518-11-04 00:46] wakes up
[1518-11-05 00:03] Guard #99 begins shift
[1518-11-05 00:45] falls asleep
[1518-11-05 00:55] wakes up`)
	got, err := MostMinuteAsleepOpportunityChecksum(r)
	if err != nil {
		t.Errorf("MostMinuteAsleepOpportunityChecksum() error = %v", err)
		return
	}
	if want := 240; got != want {
		t.Errorf("MostMinuteAsleepOpportunityChecksum() got = %v, want %v", got, want)
	}
}

func TestMostFrequentlyAsleepOpportunityChecksum(t *testing.T) {
	r := strings.NewReader(`[1518-11-01 00:00] Guard #10 begins shift
[1518-11-01 00:05] falls asleep
[1518-11-01 00:25] wakes up
[1518-11-01 00:30] falls asleep
[1518-11-01 00:55] wakes up
[1518-11-01 23:58] Guard #99 begins shift
[1518-11-02 00:40] falls asleep
[1518-11-02 00:50] wakes up
[1518-11-03 00:05] Guard #10 begins shift
[1518-11-03 00:24] falls asleep
[1518-11-03 00:29] wakes up
[1518-11-04 00:02] Guard #99 begins shift
[1518-11-04 00:36] falls asleep
[1518-11-04 00:46] wakes up
[1518-11-05 00:03] Guard #99 begins shift
[1518-11-05 00:45] falls asleep
[1518-11-05 00:55] wakes up`)
	got, err := MostFrequentlyAsleepOpportunityChecksum(r)
	if err != nil {
		t.Errorf("MostFrequentlyAsleepOpportunityChecksum() error = %v", err)
		return
	}
	if want := 4455; got != want {
		t.Errorf("MostFrequentlyAsleepOpportunityChecksum() got = %v, want %v", got, want)
	}
}
