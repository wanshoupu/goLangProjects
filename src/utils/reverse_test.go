package utils

import (
	"fmt"
	"mygo/testutils"
	"testing"
)

func TestReverse(t *testing.T) {
	test := []int{1, 3, 4, 14}
	result := Reverse(test)
	fmt.Println(result)
	testutils.AssertEq(result, []int{14, 4, 3, 1})
}
