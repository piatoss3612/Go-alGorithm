package main

// Runtime: 6ms
// Memory Usage: 3.4MB
// Time complexity: O(n)
// Space complexity: O(n)
func duplicateZeros(arr []int) {
	n := len(arr)
	res := make([]int, 0, n)

	for i := 0; i < n; i++ {
		if len(res) == n {
			break
		}

		if arr[i] == 0 {
			res = append(res, arr[i])
			if len(res) == n {
				break
			}
			res = append(res, arr[i])
		} else {
			res = append(res, arr[i])
		}
	}

	copy(arr, res)
}
