package main

// 1785. Minimum Elements to Add to Form a Given Sum
// Difficulty: Medium
// Runtime: 107ms
// Memory Usage: 9.68MB
// Time Complexity: O(n)
// space Complexity: O(1)
// Category: Greedy
func minElements(nums []int, limit int, goal int) int {
	// sum of nums
	sum := 0
	for _, num := range nums {
		sum += num
	}

	diff := abs(goal - sum) // difference between goal and sum in absolute value

	cnt := 0 // number of elements to add to nums

	for diff > 0 {
		// if diff is greater than limit
		if diff >= limit {
			cnt += diff / limit // add the quotient to cnt
			diff = diff % limit // update diff
		} else {
			cnt++ // add 1 to cnt
			break // break the loop
		}
	}

	return cnt
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
