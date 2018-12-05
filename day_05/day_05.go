package day_05

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"unicode"
	"unicode/utf8"
)

func SimplifyPolymer(reader io.Reader) int {
	r := bufio.NewReader(reader)
	result := ""
	prev := '0'
	for {
		curr, _, err := r.ReadRune()
		if err == io.EOF {
			break
		} else if unicode.IsSpace(curr) {
			continue
		}
		pIsUpper := unicode.IsUpper(prev)
		cIsUpper := unicode.IsUpper(curr)
		if pIsUpper && !cIsUpper && prev == unicode.ToUpper(curr) || !pIsUpper && cIsUpper && prev == unicode.ToLower(curr) {
			last := len(result) - 1
			result = result[:last]
			prev, _ = utf8.DecodeLastRuneInString(result)
			continue
		} else {
			result += string(curr)
			prev = curr
		}
	}
	return len(result)
}

func ImprovePolymer(reader io.Reader) int {
	buf := new(bytes.Buffer)
	n, _ := buf.ReadFrom(reader)
	min := int(n)
	s := buf.String()
	for r := 'a'; r <= 'z'; r++ {
		rm := strings.Replace(s, string(r), "", -1)
		rm = strings.Replace(rm, string(unicode.ToUpper(r)), "", -1)
		i := SimplifyPolymer(strings.NewReader(rm))
		if min > i {
			min = i
		}
	}
	return min
}
