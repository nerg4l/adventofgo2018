package day_08

import (
	"bufio"
	"io"
	"strconv"
)

func parseNodes(node []int, position int) (result []int, end int, value int) {
	childNodes := node[position]
	metadataEntries := node[position+1]
	start := position + 2
	var values []int
	for i := 0; i < childNodes; i++ {
		var r []int
		var v int
		r, start, v = parseNodes(node, start)
		result = append(result, r...)
		values = append(values, v)
	}
	end = start + metadataEntries
	metadata := node[start:end]
	result = append(result, metadata...)
	if childNodes == 0 {
		for _, v := range metadata {
			value += v
		}
	} else {
		for _, v := range metadata {
			if v <= len(values) {
				value += values[v-1]
			}
		}
	}
	return
}

func SumMetadata(reader io.Reader) int {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	var numbers []int
	for scanner.Scan() {
		text := scanner.Text()
		i, err := strconv.Atoi(text)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, i)
	}
	results, _, _ := parseNodes(numbers, 0)
	sum := 0
	for _, i := range results {
		sum += i
	}
	return sum
}

func CalcValueOfNode(reader io.Reader) int {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	var numbers []int
	for scanner.Scan() {
		text := scanner.Text()
		i, err := strconv.Atoi(text)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, i)
	}
	_, _, val := parseNodes(numbers, 0)
	return val
}
