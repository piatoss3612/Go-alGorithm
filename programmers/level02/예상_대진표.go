package main

// 문제: https://school.programmers.co.kr/learn/courses/30/lessons/12985
// 분류: 수학
func solution(n int, a int, b int) int {
	answer := 0

	for {
		answer++

		na := a / 2
		nb := b / 2

		if a%2 != 0 {
			na += 1
		}

		if b%2 != 0 {
			nb += 1
		}

		if na == nb {
			break
		}

		a, b = na, nb
	}

	return answer
}
