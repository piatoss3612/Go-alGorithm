// 2348. Number of Zero-Filled Subarrays
// Runtime: 129ms
// Memory Usage: 8.92MB
// TimeComplexity: O(n)
// SpaceComplexity: O(1)
func zeroFilledSubarray(nums []int) (ans int64) {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			ans++

			for j := i + 1; j < len(nums); j++ {
				if nums[j] == 0 {
					ans += int64(j - i + 1)
					if j == len(nums)-1 {
						i = j
						break
					}
				} else {
					i = j
					break
				}
			}
		}
	}

	return
}
