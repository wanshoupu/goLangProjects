package main

import (
	"fmt"
	"reflect"
)

func shuffle(nums []int, n int) []int {
	res := make([]int, 2*n)
	for i := 0; i < 2*n; i++ {
		if i&1 == 0 {
			res[i] = nums[i>>1]
		} else {
			res[i] = nums[n+i>>1]
		}
	}
	return res
}

func main() {
	type Test struct {
		nums []int
		k    int
		ans  []int
	}
	tests := []Test{
		{nums: []int{1, 3}, k: 1, ans: []int{1, 3}},
		{nums: []int{2, 5, 1, 3, 4, 7}, k: 3, ans: []int{2, 3, 5, 4, 1, 7}},
		{nums: []int{1, 2, 3, 4, 4, 3, 2, 1}, k: 4, ans: []int{1, 4, 2, 3, 3, 2, 4, 1}},
	}
	for _, test := range tests {
		ans := shuffle(test.nums, test.k)
		fmt.Printf("%v\n", ans)
		if !reflect.DeepEqual(ans, test.ans) {
			panic(test)
		}
	}
}
