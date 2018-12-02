package main

import "shoupu/util"

/**
 *
 * BUG: iterating through an array is very dangerous, esp for int arrays.
 * One has to remember to use:
 *	for _, j := range breakable
 */
func wordBreak(s string, wordDict []string) bool {
	lookup := toSet(wordDict)
	breakable := []int{0}
	for i := 1; i <= len(s); i++ {
		broken := false
		for _, j := range breakable {
			if _, ok := lookup[s[j:i]]; ok {
				broken = true
				break
			}
		}
		if broken {
			breakable = append(breakable, i)
		}
	}
	return breakable[len(breakable)-1] == len(s)
}

func toSet(wordDict []string) map[string]bool {
	lookup := make(map[string]bool)
	for _, w := range wordDict {
		lookup[w] = false
	}
	return lookup
}

func main() {
	type Test struct {
		str  string
		dict []string
		ans  bool
	}
	tests := []Test{
		{str: "leetcode", dict: []string{"leet", "code"}, ans: true},
		{str: "applepenapple", dict: []string{"apple", "pen"}, ans: true},
		{str: "catsandog", dict: []string{"cats", "dog", "sand", "and", "cat"}, ans: false},
	}
	for _, test := range tests {
		ans := wordBreak(test.str, test.dict)
		util.AssertEquals(ans, test.ans)
	}
}
