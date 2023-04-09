package main

// 1st solution
// Runtime: 2ms
// Memory: 2MB
// Time complexity: O(n)
// Space complexity: O(n)
func isValid(s string) bool {
	stack := []rune{}

	for _, c := range s {
		if len(stack) == 0 {
			stack = append(stack, c)
			continue
		}

		top := len(stack) - 1

		switch c {
		case ')':
			if stack[top] == '(' {
				stack = stack[:top]
			} else {
				stack = append(stack, c)
			}
		case '}':
			if stack[top] == '{' {
				stack = stack[:top]
			} else {
				stack = append(stack, c)
			}
		case ']':
			if stack[top] == '[' {
				stack = stack[:top]
			} else {
				stack = append(stack, c)
			}
		default:
			stack = append(stack, c)
		}
	}

	return len(stack) == 0
}

// 2nd solution -> code refactoring
// Runtime: 2ms
// Memory: 2.1MB
// Time complexity: O(n)
// Space complexity: O(n)
func isValid2(s string) bool {
	stack := []rune{}

	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, c := range s {
		switch c {
		case ')', '}', ']':
			top := len(stack) - 1
			if top < 0 || stack[top] != pairs[c] {
				stack = append(stack, c)
				continue
			}
			stack = stack[:top]
		default:
			stack = append(stack, c)
		}
	}

	return len(stack) == 0
}
