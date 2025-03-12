package twosum

import (
	"reflect"
	"testing"
)

type TypeTest struct {
	arg1     []int
	arg2     int
	expected []int
}

var tests = []TypeTest{
	{arg1: []int{2, 7, 11, 15}, arg2: 9, expected: []int{0, 1}},
	{arg1: []int{1, 2, 11, 15}, arg2: 3, expected: []int{0, 1}},
}

func TestTwoSum(t *testing.T) {
	for _, test := range tests {
		got := twoSum(test.arg1, test.arg2)
		if !reflect.DeepEqual(got, test.expected) {
			t.Errorf("twoSum(%v, %d) = %v; want %v", test.arg1, test.arg2, got, test.expected)
		}
	}
}
