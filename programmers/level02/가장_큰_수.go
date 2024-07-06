package main

import (
	"fmt"
	"sort"
	"strings"
)

// 문제: https://school.programmers.co.kr/learn/courses/30/lessons/42746
// 분류: 정렬
func solution(numbers []int) string {
	strs := make([]string, len(numbers))
	for i, n := range numbers {
		strs[i] = fmt.Sprintf("%d", n)
	}

	// 문자열을 조합하여 큰 수를 만들어야 하므로
	// 두 문자열을 합쳐서 비교했을 때 큰 수가 앞에 오도록 정렬한다.
	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] > strs[j]+strs[i]
	})

	answer := strings.Join(strs, "")

	// 예외 처리: 0이 여러 개일 경우 0을 반환한다.
	if answer[0] == '0' {
		return "0"
	}

	return answer
}

func main() {
	answer := solution([]int{6, 10, 2})
	fmt.Println(answer) // "6210"

	answer = solution([]int{3, 30, 34, 5, 9})
	fmt.Println(answer) // "9534330"

	answer = solution([]int{0, 0, 0, 0})
	fmt.Println(answer) // "0"
}
