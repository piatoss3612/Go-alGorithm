package main

// 2486. Append Characters to String to Make Subsequence
// Runtime: 6ms
// MemoryUsage: 5.62MB
// TimeComplexity: O(n)
// SpaceComplexity: O(1)
func appendCharacters(s string, t string) int {
	sIdx, tIdx := 0, 0

	for sIdx < len(s) && tIdx < len(t) {
		if s[sIdx] == t[tIdx] {
			tIdx++
		}
		sIdx++
	}

	return len(t) - tIdx
}
