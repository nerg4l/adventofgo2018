package main

import (
	"io"
	"strings"
	"testing"
)

func TestCalibrateResultingFrequency(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{r: strings.NewReader("+1\n-2\n+3\n+1")}, want: 3},
		{args: args{r: strings.NewReader("+1\n+1\n+1")}, want: 3},
		{args: args{r: strings.NewReader("+1\n+1\n-2")}, want: 0},
		{args: args{r: strings.NewReader("-1\n-2\n-3")}, want: -6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalibrateResultingFrequency(tt.args.r)
			if err != nil {
				t.Errorf("CalibrateResultingFrequency() error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("CalibrateResultingFrequency() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalibrateFirstRepeatingFrequency(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{r: strings.NewReader("+1\n-2\n+3\n+1")}, want: 2},
		{args: args{r: strings.NewReader("+1\n-1")}, want: 0},
		{args: args{r: strings.NewReader("+3\n+3\n+4\n-2\n-4")}, want: 10},
		{args: args{r: strings.NewReader("-6\n+3\n+8\n+5\n-6")}, want: 5},
		{args: args{r: strings.NewReader("+7\n+7\n-2\n-7\n-4")}, want: 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalibrateFirstRepeatingFrequency(tt.args.r)
			if err != nil {
				t.Errorf("CalibrateFirstRepeatingFrequency() error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("CalibrateFirstRepeatingFrequency() got = %v, want %v", got, tt.want)
			}
		})
	}
}
