package main

import "shoupu/util"

func hindex(arr []int) int {
	return 0
}

func main() {
	type Test struct {
		arr []int
		ans int
	}
	tests := []Test{
		{arr: []int{1, 23, 4}, ans: 3,},
	}
	for _, test := range tests {
		ans := hindex(test.arr)
		util.AssertEquals(ans, test)
	}
}

