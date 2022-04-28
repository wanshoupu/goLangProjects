package main

func crackSafe(n int, k int) string {
	return ""
}

func main() {
	type Test struct {
		n   int
		k   int
		ans string
	}
	tests := []Test{
		{n: 1, k: 2, ans: "",},
		{n: 1, k: 2, ans: "",},
		{n: 1, k: 2, ans: "",},
	}
	for _, test := range tests {
		ans := crackSafe(test.n, test.k)
		if ans != test.ans {
			panic(test)
		}
	}
}
