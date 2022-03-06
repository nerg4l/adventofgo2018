package main

import (
	"strings"
	"testing"
)

func TestCountFabricOverlap(t *testing.T) {
	r := strings.NewReader(`#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2`)
	got, err := CountFabricOverlap(r)
	if err != nil {
		t.Errorf("CountFabricOverlap() error = %v", err)
		return
	}
	if want := 4; got != want {
		t.Errorf("CountFabricOverlap() got = %v, want %v", got, want)
	}
}

func TestNotOverlappingFabric(t *testing.T) {
	r := strings.NewReader(`#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2`)
	got, err := NotOverlappingFabric(r)
	if err != nil {
		t.Errorf("NotOverlappingFabric() error = %v", err)
		return
	}
	if want := 3; got != want {
		t.Errorf("NotOverlappingFabric() got = %v, want %v", got, want)
	}
}
