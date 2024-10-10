package tests

import (
	"github.com/hmaidzmi/leetcode-solutions/solutions"
	"testing"
)

func TestMaxArea(t *testing.T) {
	tests := []struct {
		height   []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{1}, 0},
		{[]int{1, 1}, 1},
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{4, 8, 1, 2, 3, 9}, 32},
	}

	for _, test := range tests {
		result := solutions.MaxArea(test.height)
		if result != test.expected {
			t.Errorf("max area = %v, expected = %v", result, test.expected)
		}
	}
}
