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
		rules  rulesetArray
		length int
		check  checkFunc
	}{
		//tests function signature
		{rulesetArray{}, 0, isEmptyList},
		//tests ability to apply a rule
		{rulesetArray{{Place: 1, Rule: "fizz"}}, 1, isDeepEqual(outputList{"fizz"})},
		//tests ability to run with arbitrary length without rules
		{rulesetArray{}, 3, isDeepEqual(outputList{"1", "2", "3"})},
		//tests ability to apply rule with arbitrary length
		{rulesetArray{{Place: 3, Rule: "fizz"}}, 3, isDeepEqual(outputList{"1", "2", "fizz"})},
		//tests ability to apply multiple rules
		{rulesetArray{{Place: 2, Rule: "fizz"}, {Place: 3, Rule: "buzz"}}, 3, isDeepEqual(outputList{"1", "fizz", "buzz"})},
		//tests ability to apply multiple rules in reverse
		{rulesetArray{{Place: 3, Rule: "fizz"}, {Place: 2, Rule: "buzz"}}, 3, isDeepEqual(outputList{"1", "buzz", "fizz"})},
		//tests ability to apply overlapping rules
		{rulesetArray{{Place: 2, Rule: "fizz"}, {Place: 3, Rule: "buzz"}}, 6, isDeepEqual(outputList{"1", "fizz", "buzz", "fizz", "5", "fizzbuzz"})},
		//tests ability to apply overlapping rules in reverse order
		{rulesetArray{{Place: 3, Rule: "fizz"}, {Place: 2, Rule: "buzz"}}, 6, isDeepEqual(outputList{"1", "buzz", "fizz", "buzz", "5", "buzzfizz"})},
		//tests ability to apply arbitrary number of rules
		{rulesetArray{{Place: 1, Rule: "fizz"}, {Place: 2, Rule: "buzz"}, {Place: 3, Rule: "foo"}}, 3, isDeepEqual(outputList{"fizz", "fizzbuzz", "fizzfoo"})},
		//tests ability to apply arbitrary number of rules in reverse order
		{rulesetArray{{Place: 1, Rule: "fizz"}, {Place: 3, Rule: "buzz"}, {Place: 2, Rule: "foo"}}, 3, isDeepEqual(outputList{"fizz", "fizzfoo", "fizzbuzz"})},
	}

	//execute tests
	for _, tc := range tests {
		t.Run(fmt.Sprintf("Fizzbuzz of (%v,%v)", tc.rules, tc.length), func(t *testing.T) {
			output := FizzBuzz(tc.rules, tc.length)
			if err := tc.check(output); err != nil {
				t.Error(err)
			}
		})
	}
}
