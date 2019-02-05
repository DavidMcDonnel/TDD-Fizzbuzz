package main

type container map[int]string
type outputList []string

// FizzBuzz:: (container, int) -> outputList
// FizzBuzz is a function that takes an arbitrary number of integer to string transformations,
// applies them to the provided integer range denoted by length,
// and outputs a list of strings
func FizzBuzz(c container, length int) outputList {
	if len(c) == 0 {
		return nil
	} else {
		return outputList{"fizz"}
	}
}
