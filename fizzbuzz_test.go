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
		//tests function signature
		{container{}, 0, isEmptyList},
		//tests ability to apply a rule
		{container{1: "fizz"}, 1, isDeepEqual(outputList{"fizz"})},
		//tests ability to run with arbitrary length without rules
		{container{}, 3, isDeepEqual(outputList{"1", "2", "3"})},
		//tests ability to apply rule with arbitrary length
		{container{3: "fizz"}, 3, isDeepEqual(outputList{"1", "2", "fizz"})},
		//tests ability to apply multiple rules
		{container{2: "fizz", 3: "buzz"}, 3, isDeepEqual(outputList{"1", "fizz", "buzz"})},
		//tests ability to apply multiple rules in reverse
		{container{3: "fizz", 2: "buzz"}, 3, isDeepEqual(outputList{"1", "buzz", "fizz"})},
		//tests ability to apply overlapping rules
		{container{2: "fizz", 3: "buzz"}, 6, isDeepEqual(outputList{"1", "fizz", "buzz", "fizz", "5", "fizzbuzz"})},
		//tests ability to apply overlapping rules in reverse order
		{container{3: "fizz", 2: "buzz"}, 6, isDeepEqual(outputList{"1", "buzz", "fizz", "buzz", "5", "buzzfizz"})},
		//tests ability to apply arbitrary number of rules
		{container{1: "fizz", 2: "buzz", 3: "foo"}, 3, isDeepEqual(outputList{"fizz", "fizzbuzz", "fizzfoo"})},
		//tests ability to apply arbitrary number of rules in reverse order
		{container{1: "fizz", 3: "buzz", 2: "foo"}, 3, isDeepEqual(outputList{"fizz", "fizzfoo", "fizzbuzz"})},
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
