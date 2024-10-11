package solutions

func Trap(height []int) int {
	if len(height) < 2 {
		return 0
	}

	totalWatter := 0
	leftPointer, rightPointer := 0, len(height)-1
	maxLeft, maxRight := height[0], height[len(height)-1]
	for leftPointer < rightPointer {
		if height[leftPointer] > maxLeft {
			maxLeft = height[leftPointer]
		}
		if height[rightPointer] > maxRight {
			maxRight = height[rightPointer]
		}

		if maxLeft <= maxRight {
			if height[leftPointer] < maxLeft {
				totalWatter += maxLeft - height[leftPointer]
			}
			leftPointer++
		} else {
			if height[rightPointer] < maxRight {
				totalWatter += maxRight - height[rightPointer]
			}
			rightPointer--
		}
	}

	return totalWatter
}
