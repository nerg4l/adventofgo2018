package main

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
)

func CalibrateResultingFrequency(r io.Reader) (interface{}, error) {
	s := bufio.NewScanner(r)
	freq := 0
	for s.Scan() {
		i, err := strconv.Atoi(s.Text())
		if err != nil {
			return nil, err
		}
		freq += i
	}

	return freq, s.Err()
}

func CalibrateFirstRepeatingFrequency(r io.Reader) (interface{}, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}
	m := make(map[int]bool)
	freq := 0
	m[0] = true
	rs := bytes.NewReader(b)
	for {
		s := bufio.NewScanner(rs)
		for s.Scan() {
			i, err := strconv.Atoi(s.Text())
			if err != nil {
				return nil, err
			}

			freq += i
			if m[freq] {
				return freq, nil
			}
			m[freq] = true
		}
		if s.Err() != nil {
			return nil, s.Err()
		}
		_, _ = rs.Seek(0, 0)
	}
}
