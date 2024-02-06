package sum_test

import (
	"testing"

	"github.com/ar-sandbox3/level4/sum"
)

func TestInts(t *testing.T) {
	testCases := []struct {
		name    string
		numbers []int
		sum     int
	}{
		{"one to five", []int{1, 2, 3, 4, 5}, 15},
		{"nothing", nil, 0},
		{"one and minus one", []int{1, -1}, 0},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			s := sum.Ints(testCase.numbers...)
			if s != testCase.sum {
				t.Errorf("sum should be: %v, got: %v", testCase.sum, s)
			}
		})
	}
}
