package main

import (
	"bufio"
	"io"
	"strings"
)

func BoxIDChecksum(r io.Reader) (interface{}, error) {
	s := bufio.NewScanner(r)
	checksumTwo := 0
	checksumThree := 0
	for s.Scan() {
		text := s.Text()
		foundTwo, foundThree := false, false
		for len(text) > 0 {
			l := len(text)
			text = strings.ReplaceAll(text, string(text[0]), "")
			c := l - len(text)
			if !foundTwo && c == 2 {
				foundTwo = true
				checksumTwo++
			} else if !foundThree && c == 3 {
				foundThree = true
				checksumThree++
			}
			if foundTwo && foundThree {
				break
			}
		}
	}
	return checksumTwo * checksumThree, s.Err()
}

func BoxIDCommonLetters(r io.Reader) (interface{}, error) {
	s := bufio.NewScanner(r)
	diffs := make(map[string]map[string]bool)
	for s.Scan() {
		text := s.Text()
		for i := range text {
			lo := text[:i]
			hi := text[i+1:]
			ss, ok := diffs[lo]
			if !ok {
				ss = make(map[string]bool)
				diffs[lo] = ss
			}
			if ss[hi] {
				return lo + hi, nil
			}
			ss[hi] = true
		}
	}
	return nil, s.Err()
}
