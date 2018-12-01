package day_01

import (
	"bufio"
	"io"
	"strconv"
)

func HandleFrequencyDrift(file io.Reader) int {
	scanner := bufio.NewScanner(file)
	result := 0
	for scanner.Scan() {
		text := scanner.Text()
		i, err := strconv.Atoi(text)
		if err != nil {
			panic(err)
		}

		result += i
	}

	return result
}

func FindFirstFrequencyReachedTwice(file io.ReadSeeker) int {
	results := make(map[int]bool)
	result := 0
	for {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			results[result] = true
			text := scanner.Text()
			i, err := strconv.Atoi(text)
			if err != nil {
				panic(err)
			}

			result += i
			if results[result] {
				return result
			}
		}
		_, _ = file.Seek(0, 0)
	}
}
