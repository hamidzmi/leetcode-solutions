package solutions

func TwoSum(nums []int, target int) []int {
	seenNumsMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := seenNumsMap[target-nums[i]]; ok {
			return []int{seenNumsMap[target-nums[i]], i}
		}
		seenNumsMap[nums[i]] = i
	}

	return []int{}
}
