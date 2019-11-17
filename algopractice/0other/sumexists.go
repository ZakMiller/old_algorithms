package other

func sumExists(nums []int, sum int) bool {
	start := 0
	end := len(nums) - 1
	for nums[start]+nums[end] != sum && start != end {
		currentSum := nums[start] + nums[end]
		if currentSum > sum {
			end--
		} else {
			start++
		}
	}
	return start != end
}
