package main

import "strconv"

type container map[int]string
type outputList []string

// FizzBuzz:: (container, int) -> outputList
// FizzBuzz is a function that takes an arbitrary number of integer to string transformations,
// applies them to the provided integer range denoted by length,
// and outputs a list of strings
func FizzBuzz(c container, length int) (output outputList) {
	if len(c) == 0 {
		return nil
	}
	var result string
	for i := 1; i < length+1; i++ {
		for place, rule := range c {
			if i%place == 0 {
				result = rule
			} else {
				result = strconv.Itoa(i)
			}
		}
		output = append(output, result)
	}
	return
}
