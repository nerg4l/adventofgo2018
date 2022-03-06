package main

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestUltimateHotChocolateRecipe(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{args: args{r: strings.NewReader(`9`)}, want: "5158916779"},
		{args: args{r: strings.NewReader(`5`)}, want: "0124515891"},
		{args: args{r: strings.NewReader(`18`)}, want: "9251071085"},
		{args: args{r: strings.NewReader(`2018`)}, want: "5941429882"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UltimateHotChocolateRecipe(tt.args.r)
			if err != nil {
				t.Errorf("UltimateHotChocolateRecipe() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UltimateHotChocolateRecipe() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUltimateHotChocolateRecipeBackward(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{args: args{r: strings.NewReader(`51589`)}, want: 9},
		{args: args{r: strings.NewReader(`01245`)}, want: 5},
		{args: args{r: strings.NewReader(`92510`)}, want: 18},
		{args: args{r: strings.NewReader(`59414`)}, want: 2018},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UltimateHotChocolateRecipeBackward(tt.args.r)
			if err != nil {
				t.Errorf("UltimateHotChocolateRecipeBackward() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UltimateHotChocolateRecipeBackward() got = %v, want %v", got, tt.want)
			}
		})
	}
}
