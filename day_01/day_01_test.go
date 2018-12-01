package day_01

import (
	"strings"
	"testing"
)

func TestFrequencyResult1(t *testing.T) {
	reader := strings.NewReader("+1\n-2\n+3\n+1")
	result := HandleFrequencyDrift(reader)
	if result != 3 {
		t.Error("Wrong result")
	}
}

func TestFrequencyResult2(t *testing.T) {
	reader := strings.NewReader("+1\n+1\n+1")
	result := HandleFrequencyDrift(reader)
	if result != 3 {
		t.Error("Wrong result")
	}
}

func TestFrequencyResult3(t *testing.T) {
	reader := strings.NewReader("+1\n+1\n-2")
	result := HandleFrequencyDrift(reader)
	if result != 0 {
		t.Error("Wrong result")
	}
}

func TestFrequencyResult4(t *testing.T) {
	reader := strings.NewReader("-1\n-2\n-3")
	result := HandleFrequencyDrift(reader)
	if result != -6 {
		t.Error("Wrong result")
	}
}

func TestFrequencyRepeat1(t *testing.T) {
	reader := strings.NewReader("+1\n-2\n+3\n+1")
	result := FindFirstFrequencyReachedTwice(reader)
	if result != 2 {
		t.Error("Wrong result")
	}
}

func TestFrequencyRepeat2(t *testing.T) {
	reader := strings.NewReader("+1\n-1")
	result := FindFirstFrequencyReachedTwice(reader)
	if result != 0 {
		t.Error("Wrong result")
	}
}

func TestFrequencyRepeat3(t *testing.T) {
	reader := strings.NewReader("+3\n+3\n+4\n-2\n-4")
	result := FindFirstFrequencyReachedTwice(reader)
	if result != 10 {
		t.Error("Wrong result")
	}
}

func TestFrequencyRepeat4(t *testing.T) {
	reader := strings.NewReader("-6\n+3\n+8\n+5\n-6")
	result := FindFirstFrequencyReachedTwice(reader)
	if result != 5 {
		t.Error("Wrong result")
	}
}

func TestFrequencyRepeat5(t *testing.T) {
	reader := strings.NewReader("+7\n+7\n-2\n-7\n-4")
	result := FindFirstFrequencyReachedTwice(reader)
	if result != 14 {
		t.Error("Wrong result")
	}
}
