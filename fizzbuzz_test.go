package main

import (
	"fmt"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	type checkFunc func(outputList) error

	// checks
	isEmptyList := func(have outputList) error {
		if len(have) > 0 {
			return fmt.Errorf("Expected empty list. Found %v", have)
		}
		return nil
	}

	// test cases
	tests := [...]struct {
		rule   container
		length int
		check  checkFunc
	}{
		{container{}, 0, isEmptyList},
	}

	//execute tests
	for _, tc := range tests {
		t.Run(fmt.Sprintf("Fizzbuzz of (%v,%v)", tc.rule, tc.length), func(t *testing.T) {
			output := FizzBuzz(tc.rule, tc.length)
			if err := tc.check(output); err != nil {
				t.Error(err)
			}
		})
	}
}
