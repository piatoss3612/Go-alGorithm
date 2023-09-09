package main

import "strings"

// 1807. Evaluate the Bracket Pairs of a String
// Runtime: 184ms
// Memory Usage: 27.72MB
// Time Complexity: O(n)
// Space Complexity: O(n)
func evaluate(s string, knowledge [][]string) string {
	match := make(map[string]string)

	// convert knowledge 2d slice to map
	for _, v := range knowledge {
		match[v[0]] = v[1]
	}

	var res strings.Builder
	var key strings.Builder

	for i := 0; i < len(s); i++ {
		// if s[i] is '(', then we need to find the key
		if s[i] == '(' {
			i++

			// find the key
			for s[i] != ')' {
				key.WriteByte(s[i])
				i++
			}

			// if key exists in map, then append the value to res
			// if not, append '?'
			if val, ok := match[key.String()]; ok {
				res.WriteString(val)
			} else {
				res.WriteByte('?')
			}

			key.Reset() // reset key
		} else {
			res.WriteByte(s[i]) // append s[i] to res
		}
	}

	return res.String() // return res
}
