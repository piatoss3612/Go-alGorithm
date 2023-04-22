package main

// 1st solution
// Runtime: 26ms
// Memory Usage: 7.3MB
// Time complexity: O(n)
// Space complexity: O(n)
func decode(encoded []int, first int) []int {
	res := make([]int, 0, len(encoded)+1)
	res = append(res, first)

	for i := 0; i < len(encoded); i++ {
		first ^= encoded[i]
		res = append(res, first)
	}

	return res
}

// 2nd solution
// Runtime: 28ms
// Memory Usage: 7.7MB
// Time complexity: O(n)
// Space complexity: O(n)
func decode2(encoded []int, first int) []int {
	res := make([]int, 0, len(encoded)+1)
	res = append(res, first)
	res = append(res, encoded...)

	for i := 1; i < len(res); i++ {
		res[i] ^= res[i-1]
	}

	return res
}
