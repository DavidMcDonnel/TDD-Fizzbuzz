package main

import (
	"sort"
	"strconv"
)

type container map[int]string
type outputList []string

// FizzBuzz:: (container, int) -> outputList
// FizzBuzz is a function that takes an arbitrary number of integer to string transformations,
// applies them to the provided integer range denoted by length,
// and outputs a list of strings
func FizzBuzz(c container, length int) (output outputList) {
	keys := make([]int, 0, len(c))
	for key := range c {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for i := 1; i < length+1; i++ {
		var result string
		for _, place := range keys {
			if i%place == 0 {
				result += c[place]
			}
		}
		if result == "" {
			result = strconv.Itoa(i)
		}
		output = append(output, result)
	}
	return
}
