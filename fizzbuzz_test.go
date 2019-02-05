package main

import (
	"fmt"
	"reflect"
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

	isDeepEqual := func(want outputList) func(outputList) error {
		return func(have outputList) error {
			if !reflect.DeepEqual(have, want) {
				return fmt.Errorf("Expected equivalent list. Found %v, wanted %v", have, want)
			}
			return nil
		}
	}

	// test cases
	tests := [...]struct {
		rule   container
		length int
		check  checkFunc
	}{
		{container{}, 0, isEmptyList},
		{container{1: "fizz"}, 1, isDeepEqual(outputList{"fizz"})},
		{container{3: "fizz"}, 3, isDeepEqual(outputList{"1", "2", "fizz"})},
		{container{2: "fizz", 3: "buzz"}, 3, isDeepEqual(outputList{"1", "fizz", "buzz"})},
		{container{3: "fizz", 2: "buzz"}, 3, isDeepEqual(outputList{"1", "buzz", "fizz"})},
		{container{2: "fizz", 3: "buzz"}, 6, isDeepEqual(outputList{"1", "fizz", "buzz", "fizz", "5", "fizzbuzz"})},
		{container{3: "fizz", 2: "buzz"}, 6, isDeepEqual(outputList{"1", "buzz", "fizz", "buzz", "5", "buzzfizz"})},
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
