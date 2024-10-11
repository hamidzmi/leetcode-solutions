package tests

import (
	"github.com/hmaidzmi/leetcode-solutions/solutions"
	"testing"
)

func TestTrap(t *testing.T) {
	tests := []struct {
		height []int
		expect int
	}{
		{[]int{}, 0},
		{[]int{1}, 0},
		{[]int{1, 2}, 0},
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
	}

	for _, test := range tests {
		result := solutions.Trap(test.height)
		if result != test.expect {
			t.Errorf("Trap(%v) = %v; expect %v", test.height, result, test.expect)
		}
	}
}
