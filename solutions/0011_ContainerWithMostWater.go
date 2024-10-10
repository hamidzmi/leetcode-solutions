package solutions

import "math"

func MaxArea(height []int) int {
	maxArea := 0
	p1 := 0
	p2 := len(height) - 1

	for p1 < p2 {
		h := int(math.Min(float64(height[p1]), float64(height[p2])))
		w := p2 - p1
		maxArea = int(math.Max(float64(h*w), float64(maxArea)))

		if height[p1] < height[p2] {
			p1++
		} else {
			p2--
		}
	}

	return maxArea
}
