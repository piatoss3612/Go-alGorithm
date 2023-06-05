package main

// 1st solution
// Runtime: 30ms
// Memory: 6.4MB
// Time Complexity: O(nlogn) -> sort.Ints
func missingNumber1(nums []int) int {
	sort.Ints(nums)

	target := 0

	for _, num := range nums {
		if num != target {
			return target
		}

		target++
	}

	return target
}

// 2nd solution
// Runtime: 21ms
// Memory: 6.5MB
// Time Complexity: O(n)
func missingNumber2(nums []int) int {
	check := make([]bool, len(nums)+1)
	for _, v := range nums {
		check[v] = true
	}

	for i, v := range check {
		if !v {
			return i
		}
	}
	return -1
}

// 3rd solution
// Runtime: 15ms
// Memory: 6.4MB
// Time Complexity: O(n)
func missingNumber3(nums []int) int {
	mask := 0

	for i := 0; i <= len(nums); i++ {
		mask ^= i
	}

	for i := 0; i < len(nums); i++ {
		mask ^= nums[i]
	}

	return mask
}

// 4th solution
// Runtime: 17ms
// Memory: 6.3MB
// Time Complexity: O(n)
func missingNumber4(nums []int) (sum int) {
	n := len(nums) + 1
	sum = (n * (n-1)) / 2

	for i := 0; i < len(nums); i++ {
		sum -= nums[i]
	}

	return
}