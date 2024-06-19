package level02

// 문제: https://school.programmers.co.kr/learn/courses/30/lessons/12951
// 분류: 구현
func solution(s string) string {
	b := []byte(s)

	for i := 0; i < len(b); i++ {
		if i == 0 || (i > 0 && b[i-1] == ' ') {
			// Capitalize the current character
			if b[i] >= 'a' && b[i] <= 'z' {
				b[i] -= 32
			}
		} else {
			// Lowercase the current character
			if b[i] >= 'A' && b[i] <= 'Z' {
				b[i] += 32
			}
		}
	}

	return string(b)
}
