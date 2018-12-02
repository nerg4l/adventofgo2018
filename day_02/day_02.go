package day_02

import (
	"bufio"
	"io"
	"strings"
)

func CheckSum(file io.Reader) int {
	scanner := bufio.NewScanner(file)
	checkSumTwo := 0
	checkSumThree := 0
	for scanner.Scan() {
		text := scanner.Text()
		foundTwo, foundThree := false, false
		for _, r := range text {
			c := countRune(text, r)
			if !foundTwo && c == 2 {
				foundTwo = true
				checkSumTwo++
			} else if !foundThree && c == 3 {
				foundThree = true
				checkSumThree++
			}
			if foundTwo && foundThree {
				break
			}
		}
	}
	return checkSumTwo * checkSumThree
}

func countRune(s string, r rune) int {
	n := 0
	for {
		i := strings.IndexRune(s, r)
		if i == -1 {
			return n
		}
		n++
		s = s[i+1:]
	}
}

func FindTheBoxesFullOfPrototypeFabric(file io.Reader) string {
	scanner := bufio.NewScanner(file)
	result := ""
	var lines []string
	for scanner.Scan() {
		text := scanner.Text()
		for _, target := range lines {
			var differences []int
			for i := range text {
				if text[i] != target[i] {
					differences = append(differences, i)
				}
			}
			if len(differences) == 1 {
				i := differences[0]
				result = text[:i] + text[i+1:]
				break
			}
		}
		if result != "" {
			break
		}
		lines = append(lines, text)
	}
	return result
}
