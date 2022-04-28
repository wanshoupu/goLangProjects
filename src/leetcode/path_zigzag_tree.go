package main

import (
	"fmt"
	"mygo/utils"
)

/**
https://leetcode.com/problems/path-in-zigzag-labelled-binary-tree/
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
	var labels []int
	for label > 0 {
		labels = append(labels, label)
		label = utils.ParentLabel(label)
	}
	return utils.Reverse(labels)
}

func main() {
	tests := []struct {
		Input  int
		Output []int
	}{
		{4, []int{1, 3, 4}},
		{3, []int{1, 3}},
		{14, []int{1, 3, 4, 14}},
		{1, []int{1}},
		{19, []int{1, 3, 4, 14, 19}},
	}
	for _, test := range tests {
		path := pathInZigZagTree(test.Input)
		fmt.Println(test.Input, path)
		utils.AssertEq(test.Output, path)
	}
}
