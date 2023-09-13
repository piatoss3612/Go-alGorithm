package main

import "fmt"

// 1424. Diagonal Traverse II
// Runtime: 4435ms...
// Memory Usage: 13.13MB
// Time Complexity: O(n*m)
// Space Complexity: O(n*m)
func findDiagonalOrder(nums [][]int) (ans []int) {
	n := len(nums)
	m := 0

	// find max length of row
	for _, row := range nums {
		if len(row) > m {
			m = len(row)
		}
	}

	// find all elements in diagonal order from top-left to middle diagonal
	for k := 0; k < n; k++ {
		i, j := k, 0
		for i >= 0 && j < n {
			if j < len(nums[i]) {
				ans = append(ans, nums[i][j])
			}
			i, j = i-1, j+1
		}
	}

	// find all elements in diagonal order from middle diagonal to bottom-right
	for k := 1; k < m; k++ {
		i, j := n-1, k
		for i >= 0 && j < m {
			if j < len(nums[i]) {
				ans = append(ans, nums[i][j])
			}
			i, j = i-1, j+1
		}
	}

	return
}

// 2nd Solution
// Runtime: 1954ms
// Memory Usage: 14.55MB
// Time Complexity: O(n*m)
// Space Complexity: O(n*m)
func findDiagonalOrder2(nums [][]int) (ans []int) {
	n := len(nums)
	m := 0

	// find max length of row
	for _, row := range nums {
		if len(row) > m {
			m = len(row)
		}
	}

	temp := make([][]int, n+m-1)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if len(nums[i]) > j {
				temp[i+j] = append(temp[i+j], nums[i][j])
			}
		}
	}

	for _, row := range temp {
		for i := len(row) - 1; i >= 0; i-- {
			ans = append(ans, row[i])
		}
	}

	return
}

func main() {
	fmt.Println(findDiagonalOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}
