package solutions

func TwoSum(nums []int, target int) []int {
	lookingNums := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := lookingNums[target-nums[i]]; ok {
			return []int{lookingNums[target-nums[i]], i}
		}
		lookingNums[nums[i]] = i
	}

	return []int{}
}
