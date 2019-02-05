package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	type checkFunc func([]string) error

	// checks
	isEmptyList := func(have []string) error {
		if len(have) > 0 {
			return fmt.Errorf("Expected empty list. Found %v", have)
		}
		return nil
	}

	isDeepEqual := func(want []string) func([]string) error {
		return func(have []string) error {
			if !reflect.DeepEqual(have, want) {
				return fmt.Errorf("Expected equivalent list. Found %v, wanted %v", have, want)
			}
			return nil
		}
	}

	// test cases
	tests := [...]struct {
		rules  ruleArray
		length int
		check  checkFunc
	}{
		//tests function signature
		{ruleArray{}, 0, isEmptyList},
		//tests ability to apply a rule
		{ruleArray{{Place: 1, Rule: "fizz"}}, 1, isDeepEqual([]string{"fizz"})},
		//tests ability to run with arbitrary length without rules
		{ruleArray{}, 3, isDeepEqual([]string{"1", "2", "3"})},
		//tests ability to apply rule with arbitrary length
		{ruleArray{{Place: 3, Rule: "fizz"}}, 3, isDeepEqual([]string{"1", "2", "fizz"})},
		//tests ability to apply multiple rules
		{ruleArray{{Place: 2, Rule: "fizz"}, {Place: 3, Rule: "buzz"}}, 3, isDeepEqual([]string{"1", "fizz", "buzz"})},
		//tests ability to apply multiple rules in reverse
		{ruleArray{{Place: 3, Rule: "fizz"}, {Place: 2, Rule: "buzz"}}, 3, isDeepEqual([]string{"1", "buzz", "fizz"})},
		//tests ability to apply overlapping rules
		{ruleArray{{Place: 2, Rule: "fizz"}, {Place: 3, Rule: "buzz"}}, 6, isDeepEqual([]string{"1", "fizz", "buzz", "fizz", "5", "fizzbuzz"})},
		//tests ability to apply overlapping rules in reverse order
		{ruleArray{{Place: 3, Rule: "fizz"}, {Place: 2, Rule: "buzz"}}, 6, isDeepEqual([]string{"1", "buzz", "fizz", "buzz", "5", "buzzfizz"})},
		//tests ability to apply arbitrary number of rules
		{ruleArray{{Place: 1, Rule: "fizz"}, {Place: 2, Rule: "buzz"}, {Place: 3, Rule: "foo"}}, 3, isDeepEqual([]string{"fizz", "fizzbuzz", "fizzfoo"})},
		//tests ability to apply arbitrary number of rules in reverse order
		{ruleArray{{Place: 1, Rule: "fizz"}, {Place: 3, Rule: "buzz"}, {Place: 2, Rule: "foo"}}, 3, isDeepEqual([]string{"fizz", "fizzfoo", "fizzbuzz"})},
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
