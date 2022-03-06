package main

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestFuelCell_PowerLevel(t *testing.T) {
	type fields struct {
		x            int
		y            int
		serialNumber int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{fields: fields{x: 3, y: 5, serialNumber: 8}, want: 4},
		{fields: fields{x: 122, y: 79, serialNumber: 57}, want: -5},
		{fields: fields{x: 217, y: 196, serialNumber: 39}, want: 0},
		{fields: fields{x: 101, y: 153, serialNumber: 71}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FuelCell{
				x:            tt.fields.x,
				y:            tt.fields.y,
				serialNumber: tt.fields.serialNumber,
			}
			if got := f.PowerLevel(); got != tt.want {
				t.Errorf("PowerLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLargestTotalPowerWithDefaultSize(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{args: args{r: strings.NewReader(`18`)}, want: "33,45"},
		{args: args{r: strings.NewReader(`42`)}, want: "21,61"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LargestTotalPowerWithDefaultSize(tt.args.r)
			if err != nil {
				t.Errorf("LargestTotalPowerWithDefaultSize() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LargestTotalPowerWithDefaultSize() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLargestTotalPowerOfAllSize(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{args: args{r: strings.NewReader(`18`)}, want: "90,269,16"},
		{args: args{r: strings.NewReader(`42`)}, want: "232,251,12"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := LargestTotalPowerOfAllSize(tt.args.r)
			if err != nil {
				t.Errorf("LargestTotalPowerOfAllSize() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LargestTotalPowerOfAllSize() got = %v, want %v", got, tt.want)
			}
		})
	}
}
