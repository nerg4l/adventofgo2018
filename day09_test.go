package main

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestWinningElfsScore(t *testing.T) {
	type args struct {
		r          io.Reader
		multiplier uint
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{args: args{r: strings.NewReader(`9 players; last marble is worth 25 points`), multiplier: 1}, want: 32},
		{args: args{r: strings.NewReader(`10 players; last marble is worth 1618 points`), multiplier: 1}, want: 8317},
		{args: args{r: strings.NewReader(`13 players; last marble is worth 7999 points`), multiplier: 1}, want: 146373},
		{args: args{r: strings.NewReader(`17 players; last marble is worth 1104 points`), multiplier: 1}, want: 2764},
		{args: args{r: strings.NewReader(`21 players; last marble is worth 6111 points`), multiplier: 1}, want: 54718},
		{args: args{r: strings.NewReader(`30 players; last marble is worth 5807 points`), multiplier: 1}, want: 37305},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WinningElfsScore(tt.args.r, tt.args.multiplier)
			if err != nil {
				t.Errorf("WinningElfsScore() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WinningElfsScore() got = %v, want %v", got, tt.want)
			}
		})
	}
}
