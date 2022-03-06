package main

import (
	"bufio"
	"io"
	"strconv"
)

type licenseNode struct {
	metadata []int
	children []*licenseNode
}

func parseLicenseTree(r io.Reader) (*licenseNode, error) {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)
	root := &licenseNode{}
	stack := []*licenseNode{root}
	// Using stack instead of recursion.
	for s.Scan() {
		text := s.Text()
		i, err := strconv.Atoi(text)
		if err != nil {
			panic(err)
		}
		n := stack[len(stack)-1]
		// Set root node children capacity
		if n.children == nil {
			n.children = make([]*licenseNode, 0, i)
			continue
		}
		// Set metadata capacity
		if n.metadata == nil {
			n.metadata = make([]int, 0, i)
			continue
		}
		// Push to stack with children capacity
		if len(n.children) < cap(n.children) {
			child := &licenseNode{children: make([]*licenseNode, 0, i)}
			n.children = append(n.children, child)
			stack = append(stack, child)
			continue
		}
		// Push metadata
		if len(n.metadata) < cap(n.metadata) {
			n.metadata = append(n.metadata, i)
		}
		// Pop from stack
		if len(n.metadata) == cap(n.metadata) {
			stack = stack[:len(stack)-1]
		}
	}
	return root, s.Err()
}

func SumMetadata(r io.Reader) (interface{}, error) {
	root, err := parseLicenseTree(r)
	if err != nil {
		return nil, err
	}
	// Using stack instead of recursion.
	stack := []*licenseNode{root}
	sum := 0
	for len(stack) > 0 {
		n := stack[len(stack)-1]
		for _, i := range n.metadata {
			sum += i
		}
		stack = append(stack[:len(stack)-1], n.children...)
	}
	return sum, nil
}

func ValueOfNode(r io.Reader) (interface{}, error) {
	root, err := parseLicenseTree(r)
	if err != nil {
		return nil, err
	}
	stack := []*licenseNode{root}
	val := 0
	// Using stack instead of recursion.
	for len(stack) > 0 {
		n := stack[0]
		if len(n.children) == 0 {
			for _, i := range n.metadata {
				val += i
			}
		} else {
			for _, i := range n.metadata {
				if i <= len(n.children) {
					stack = append(stack, n.children[i-1])
				}
			}
		}
		stack = stack[1:]
	}
	return val, nil
}
