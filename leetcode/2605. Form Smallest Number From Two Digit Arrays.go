package main

// 1st solution
// Runtime: 4ms
// Memory Usage: 2.2MB
// Time Complexity: O(n)
// Space Complexity: O(n)
func minNumber(nums1 []int, nums2 []int) int {
	set := make(map[int]int)

	min1, min2 := 987654321, 987654321

	for i := 0; i < len(nums1); i++ {
		if nums1[i] < min1 {
			min1 = nums1[i]
		}
		set[nums1[i]] += 1
	}

	for i := 0; i < len(nums2); i++ {
		if nums2[i] < min2 {
			min2 = nums2[i]
		}
		set[nums2[i]] += 1
	}

	min3 := 987654321

	for k, v := range set {
		if v > 1 {
			if k < min3 {
				min3 = k
			}
		}
	}

	if min3 != 987654321 {
		return min3
	}

	if min1 > min2 {
		min1, min2 = min2, min1
	}

	return min1*10 + min2
}

// 2nd solution
// Runtime: 0ms
// Memory Usage: 2.1MB
// Time Complexity: O(n)
// Space Complexity: O(1)
func minNumber2(nums1 []int, nums2 []int) int {
	bits1, digit1 := numsToBits(nums1)
	bits2, digit2 := numsToBits(nums2)

	if bits1&bits2 > 0 {
		return extractMinDigit(bits1 & bits2)
	}

	if digit1 > digit2 {
		digit1, digit2 = digit2, digit1
	}

	return digit1*10 + digit2
}

func numsToBits(nums []int) (bits, digit int) {
	digit = 987654321
	for i := 0; i < len(nums); i++ {
		bits += 1 << nums[i]
		if nums[i] < digit {
			digit = nums[i]
		}
	}
	return
}

func extractMinDigit(bits int) int {
	for i := 1; i <= 9; i++ {
		if bits&(1<<i) > 0 {
			return i
		}
	}
	return 0
}
