package main

import (
	"sort"
	"strconv"
)

type ruleset struct {
	Place int
	Rule  string
}
type rulesetArray []ruleset
type outputList []string

func (r rulesetArray) Less(i, j int) bool { return r[i].Place < r[j].Place }
func (r rulesetArray) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r rulesetArray) Len() int           { return len(r) }

// FizzBuzz:: (rulesetArray, int) -> outputList
// FizzBuzz is a function that takes an arbitrary number of integer to string transformations,
// applies them to the provided integer range denoted by length,
// and outputs a list of strings
func FizzBuzz(c rulesetArray, length int) (output outputList) {
	sort.Sort(c)
	for i := 1; i < length+1; i++ {
		var result string
		for _, rs := range c {
			if i%rs.Place == 0 {
				result += rs.Rule
			}
		}
		if result == "" {
			result = strconv.Itoa(i)
		}
		output = append(output, result)
	}
	return
}
