package main

import (
	"fmt"
	"mygo/utils"
)

/**
In an infinite binary tree where every node has two children,
the nodes are labelled in row order.

In the odd numbered rows (ie., the first, third, fifth,...), the labelling is left to right, while in the even numbered rows (second, fourth, sixth,...), the labelling is right to left.

Given the label of a node in this tree, return the labels in the path from the root of the tree to the node with that label.

Example 1:

Input: label = 14
Output: [1,3,4,14]
Example 2:

Input: label = 26
Output: [1,2,6,10,26]
*/
func pathInZigZagTree(label int) []int {
	layer := GetLayer(label)
	rank := label - (1 << layer)
	labels := []int{rank}
	for rank > 0 {
		rank >>= 1
		labels = append(labels, rank)
	}
	for i := range labels {
		if i&1 > 1 {
			labels[i] = 1<<i - labels[i]
		}
		labels[i]++
	}
	return utils.Reverse(labels)
}

func GetLayer(label int) int {
	var layer int
	for {
		if label == 0 {
			return layer
		}
		label &= label - 1
		layer++
	}
}

func main() {
	tests := []struct {
		Input  int
		Output []int
	}{
		{14, []int{1, 3, 4, 14}},
		{1, []int{}},
		{0, []int{}},
	}
	for _, test := range tests {
		label := pathInZigZagTree(test.Input)
		fmt.Printf("%s\n", label)
		//AssertEq(test.Output, label)
	}
}

func AssertEq(expected []int, actual []int) {
	if !testEq(expected, actual) {
		panic(fmt.Sprintf("not equal: '%s', '%s'", expected, actual))
	}
}

func testEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
