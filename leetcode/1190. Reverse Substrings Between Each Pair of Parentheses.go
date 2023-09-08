package main

// 1190. Reverse Substrings Between Each Pair of Parentheses
// Runtime: 1ms
// Memory Usage: 1.96MB
// Time complexity: O(n^2)
// Space complexity: O(n)
func reverseParentheses(s string) string {
	stack := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ')':
			for j := len(stack) - 1; j >= 0; j-- {
				if stack[j] == '(' {
					for k, l := j+1, len(stack)-1; k < l; k, l = k+1, l-1 {
						stack[k], stack[l] = stack[l], stack[k]
					}

					stack = append(stack[:j], stack[j+1:]...)

					break
				}
			}
		default:
			stack = append(stack, s[i])
		}
	}

	return string(stack)
}

func main() {
	s := "(u(love)i)"
	println(reverseParentheses(s))
}
