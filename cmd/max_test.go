package main

import (
	"testing"
)

func TestMax(t *testing.T) {
	//Initial values
	testTable := []struct {
		number   int
		expected bool
	}{
		{
			number:   22,
			expected: true,
		},
		{
			number:   -2,
			expected: true,
		},
		{
			number:   10,
			expected: true,
		},
	}

	//Act
	for _, testCase := range testTable {
		result := Max(testCase.number)

		if result != testCase.expected {
			t.Errorf("Calling Max(%v). Incorrect result. Expect %v, got %v.", testCase.number, testCase.expected, result)
		} else {
			t.Logf("Calling Max(%v). Correct result: %v.", testCase.number, result)
		}

	}
}
