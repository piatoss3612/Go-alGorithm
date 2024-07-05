package main

import "fmt"

func solution(number string, k int) string {
	stack := make([]rune, 0, len(number)) // 스택 초기화
	length := 0                           // 스택의 길이
	cnt := 0                              // 제거한 수의 개수

	for _, c := range number {
		// 스택에서 c보다 작은 수를 제거하기
		for cnt < k && length > 0 {
			top := stack[length-1]
			if top >= c {
				break
			}

			stack = stack[:length-1]
			length--
			cnt++
		}

		// 스택에 c추가
		stack = append(stack, c)
		length++
	}

	// 혹시 모를 예외: 항상 k개의 수를 제거해야 함
	if cnt < k {
		stack = stack[:length-(k-cnt)]
	}

	return string(stack)
}

func main() {
	a := solution("4177252841", 4)
	fmt.Println(a)
}
