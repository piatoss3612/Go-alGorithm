package level02

import "fmt"

// 문제: https://school.programmers.co.kr/learn/courses/30/lessons/12939
// 분류: 구현
func solution(s string) string {
	arr := []int{}

	var n int
	var minus bool
	for _, c := range s {
		if c == ' ' {
			if minus {
				n = -n
			}
			arr = append(arr, n)
			n = 0
			minus = false
		} else if c == '-' {
			minus = true
		} else {
			n = n*10 + int(c-'0')
		}
	}

	if minus {
		n = -n
	}
	arr = append(arr, n)

	min, max := arr[0], arr[0]

	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	return fmt.Sprintf("%d %d", min, max)
}
