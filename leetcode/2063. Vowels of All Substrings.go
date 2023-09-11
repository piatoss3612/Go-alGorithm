package main

import "fmt"

// 2063. Vowels of All Substrings
// Runtime: 2ms
// Memory Usage: 6.23MB
// TimeComplexity: O(n)
// SpaceComplexity: O(1)
// Ref: https://leetcode.com/problems/vowels-of-all-substrings/solutions/1563780/java-c-python-easy-and-concise-o-n/
func countVowels(word string) int64 {
	ans := 0
	n := len(word)

	for i := 0; i < n; i++ {
		if isVowel(word[i]) {
			// there could be substring start with word[x] and end with word[y]
			// where 0 <= x <= i and i <= y < n
			// there are (i + 1) choices for x and (n - i) choices for y
			// so there are (i + 1) * (n - i) substrings containing word[i] as vowel
			ans += (i + 1) * (n - i)
		}
	}

	return int64(ans)
}

func isVowel(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u'
}

func main() {
	word := "aba"
	fmt.Println(countVowels(word))
}
