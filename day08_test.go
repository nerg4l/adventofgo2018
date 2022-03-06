package main

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestSumMetadata(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{args: args{r: strings.NewReader(`2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2`)}, want: 138},
		{args: args{r: strings.NewReader(`2 3 1 1 0 1 99 2 0 3 10 11 12 1 1 2`)}, want: 138},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumMetadata(tt.args.r)
			if err != nil {
				t.Errorf("SumMetadata() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumMetadata() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValueOfNode(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{args: args{r: strings.NewReader(`2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2`)}, want: 66},
		{args: args{r: strings.NewReader(`0 3 10 11 12`)}, want: 33},
		{args: args{r: strings.NewReader(`1 1 0 1 99 2`)}, want: 0},
		{args: args{r: strings.NewReader(`0 1 99`)}, want: 99},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValueOfNode(tt.args.r)
			if err != nil {
				t.Errorf("ValueOfNode() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValueOfNode() got = %v, want %v", got, tt.want)
			}
		})
	}
}
