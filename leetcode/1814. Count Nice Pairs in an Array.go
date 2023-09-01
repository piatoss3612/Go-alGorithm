package main

import "fmt"

const mod = 1e9 + 7

// 1814. Count Nice Pairs in an Array
// Runtime: 154ms
// MemoryUsage: 12.27MB
// TimeComplexity: O(n)
// SpaceComplexity: O(n)
func countNicePairs(nums []int) int {
	n := len(nums)

	// nums[i] + rev(nums[i]) = nums[j] + rev(nums[j]) -> nums[i] - rev(nums[i]) = nums[j] - rev(nums[j])
	for i := 0; i < n; i++ {
		nums[i] = nums[i] - rev(nums[i])
	}

	cnt := 0 // count of pairs

	included := make(map[int]int) // included[nums[i]] = count of nums[i] in nums

	for i := 0; i < n; i++ {
		// if nums[i] is already in nums, then there are included[nums[i]] pairs
		if included[nums[i]] > 0 {
			cnt = (cnt + included[nums[i]]) % mod
			included[nums[i]]++ // add one more nums[i] to included
		} else {
			included[nums[i]] = 1 // add nums[i] to included
		}
	}

	return cnt // return count of pairs
}

func rev(x int) int {
	var res int
	for x > 0 {
		res = res*10 + x%10
		x /= 10
	}
	return res
}

func main() {
	arr := []int{13, 10, 35, 24, 76}
	fmt.Println(countNicePairs(arr))
}
