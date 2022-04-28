package utils

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	test := []int{1, 3, 4, 14}
	result := Reverse(test)
	fmt.Println(result)
	AssertEq(result, []int{14, 4, 3, 1})
}

func TestReverseOdd(t *testing.T) {
	test := []int{1, 3, 14}
	result := Reverse(test)
	fmt.Println(result)
	AssertEq(result, []int{14, 3, 1})
}
