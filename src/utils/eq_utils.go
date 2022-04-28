package utils

import "fmt"

func AssertEq(a, b []int) {
	if len(a) != len(b) {
		panic("unequal size")
	}
	for i := range a {
		if a[i] != b[i] {
			panic(fmt.Sprintf("unequal elements: %d != %d", a[i], b[i]))
		}
	}
}
