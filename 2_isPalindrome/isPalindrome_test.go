package ispalindrome

import (
	"reflect"
	"testing"
)

type TypeTest struct {
	arg1     int
	expected bool
}

var tests = []TypeTest{
	{arg1: 121, expected: true},
	{arg1: 123, expected: false},
	{arg1: 12321, expected: true},
	{arg1: 1000021, expected: false},
	{arg1: -121, expected: false},
}

func TestTwoSum(t *testing.T) {
	for _, test := range tests {
		got := ispalindrome(test.arg1)
		if !reflect.DeepEqual(got, test.expected) {
			t.Errorf("ispalindrome(%v) = %v; want %v", test.arg1, got, test.expected)
		}
	}
}
