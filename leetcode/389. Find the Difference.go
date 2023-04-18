package main

// 1st solution - count the number of each alphabet in s and t
// Runtime: 0ms
// Memory: 2.2MB
// Time Complexity: O(n) -> n is the length of s
// Space Complexity: O(n)
func findTheDifference(s string, t string) byte {
	origin := make([]int, 26)
	diff := make([]int, 26)

	for i := 0; i < len(s); i++ {
		origin[int(s[i]-'a')] += 1
		diff[int(t[i]-'a')] += 1
	}

	diff[int(t[len(t)-1]-'a')] += 1

	var ans byte

	for i := 0; i < 26; i++ {
		if diff[i] > origin[i] {
			ans = byte(i + 'a')
			break
		}
	}
	return ans
}

// 2nd solution - byte addition
// Runtime: 2ms
// Memory: 2.2MB
// Time Complexity: O(n)
// Space Complexity: O(1)
func findTheDifference2(s string, t string) byte {
	var origin, diff byte

	for i := 0; i < len(s); i++ {
		origin += s[i]
		diff += t[i]
	}

	diff += t[len(t)-1]

	return diff - origin
}
