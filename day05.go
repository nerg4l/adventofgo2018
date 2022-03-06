package main

import (
	"bufio"
	"io"
	"strings"
	"unicode"
)

func SimplifyPolymer(r io.Reader) (interface{}, error) {
	return simplifyPolymer(r)
}

func simplifyPolymer(r io.Reader) (int, error) {
	buf := bufio.NewReader(r)
	var result []rune
	for {
		r, _, err := buf.ReadRune()
		if err == io.EOF {
			break
		} else if err != nil {
			return 0, err
		} else if unicode.IsSpace(r) {
			continue
		}
		if len(result) < 1 {
			result = append(result, r)
			continue
		}
		prev := result[len(result)-1]
		if (r-prev) == 'A'-'a' || (prev-r) == 'A'-'a' {
			result = result[:len(result)-1]
			continue
		}
		result = append(result, r)
	}
	return len(result), nil
}

func ImprovePolymer(r io.Reader) (interface{}, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	s := string(b)
	min := len(s)
	for lr := 'a'; lr <= 'z'; lr++ {
		ur := unicode.ToUpper(lr)
		r := strings.NewReader(strings.Map(func(r rune) rune {
			if r == lr || r == ur {
				return -1
			}
			return r
		}, s))
		i, err := simplifyPolymer(r)
		if err != nil {
			return nil, err
		}
		if min > i {
			min = i
		}
	}
	return min, nil
}
