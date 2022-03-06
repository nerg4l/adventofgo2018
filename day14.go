package main

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"unicode"
)

func UltimateHotChocolateRecipe(r io.Reader) (interface{}, error) {
	var expectedScoreOfTen int
	_, err := fmt.Fscanf(r, "%d", &expectedScoreOfTen)
	if err != nil {
		return nil, err
	}
	numbers := []int{3, 7}
	workers := []int{0, 1}
	for len(numbers) <= expectedScoreOfTen+10 {
		var sum int
		for _, v := range workers {
			sum += numbers[v]
		}
		if sum > 9 {
			numbers = append(numbers, sum/10)
		}
		numbers = append(numbers, sum%10)
		for i, v := range workers {
			workers[i] = (v + numbers[v] + 1) % len(numbers)
		}
	}
	buf := new(bytes.Buffer)
	for i := 0; i < 10; i++ {
		j := expectedScoreOfTen + i
		buf.WriteRune(rune(numbers[j] + '0'))
	}
	s := buf.String()
	return s, nil
}

func UltimateHotChocolateRecipeBackward(r io.Reader) (interface{}, error) {
	var tmpString string
	_, err := fmt.Fscanf(r, "%s", &tmpString)
	if err != nil {
		return nil, err
	}
	var expected []int
	for _, v := range tmpString {
		if !unicode.IsDigit(v) {
			continue
		}
		expected = append(expected, int(v-'0'))
	}
	expectedLen := len(expected)
	numbers := []int{3, 7}
	workers := []int{0, 1}
	found := 0
	for i := 0; found == 0; i++ {
		var sum int
		for _, v := range workers {
			sum += numbers[v]
		}
		if sum > 9 {
			numbers = append(numbers, sum/10)
		}
		numbers = append(numbers, sum%10)
		if len(numbers) > expectedLen {
			for j := 0; j < 2; j++ {
				start := len(numbers) - expectedLen - j
				end := len(numbers) - j
				check := numbers[start:end]
				if reflect.DeepEqual(expected, check) {
					found = len(numbers) - expectedLen - j
				}
			}
		}
		for i, v := range workers {
			workers[i] = (v + numbers[v] + 1) % len(numbers)
		}
	}
	return found, nil
}
