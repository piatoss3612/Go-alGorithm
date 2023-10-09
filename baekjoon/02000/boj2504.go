package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)

	S string
)

// 난이도: Silver 1
// 메모리: 924KB
// 시간: 4ms
// 분류: 구현, 자료 구조, 스택
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	S = scanString()
}

func Solve() {
	stack := make([]byte, 0, len(S))
	memo := map[int]int{} // memo[depth] = 깊이가 depth일 때 괄호열의 값
	depth := 0 // 괄호열의 깊이
	prev := -1 // stack의 마지막 인덱스

	for i := 0; i < len(S); i++ {
		switch S[i] {
		case ')':
			// stack이 비어있거나, stack의 마지막 원소가 '('가 아니면 올바른 괄호열이 아님
			if prev == -1 || stack[prev] != '(' {
				// 올바른 괄호열이 아니면 0 출력 후 종료
				fmt.Fprintln(writer, 0)
				return
			}

			// stack에서 '('를 제거
			stack = stack[:prev]
			prev -= 1

			// memo[depth]가 0이 아니면, 부모 괄호열을 발견한 것이므로 memo[depth-1]에 2를 곱해줌
			if memo[depth] != 0 {
				memo[depth-1] += memo[depth] * 2
				memo[depth] = 0
			} else {
				memo[depth-1] += 2 // memo[depth]가 0이면, memo[depth-1]에 2를 더해줌
			}
			depth -= 1 // 괄호열의 깊이를 1 감소
		case ']':
			// stack이 비어있거나, stack의 마지막 원소가 '['가 아니면 올바른 괄호열이 아님
			if prev == -1 || stack[prev] != '[' {
				fmt.Fprintln(writer, 0)
				return
			}

			// stack에서 '['를 제거
			stack = stack[:prev]
			prev -= 1

			// memo[depth]가 0이 아니면, 부모 괄호열을 발견한 것이므로 memo[depth-1]에 3을 곱해줌
			if memo[depth] != 0 {
				memo[depth-1] += memo[depth] * 3
				memo[depth] = 0
			} else {
				memo[depth-1] += 3 // memo[depth]가 0이면, memo[depth-1]에 3을 더해줌
			}
			depth -= 1 // 괄호열의 깊이를 1 감소
		default:
			// stack에 '(' 또는 '['를 추가, 깊이와 stack의 마지막 인덱스를 갱신
			stack = append(stack, S[i])
			depth += 1
			prev += 1
		}
	}

	fmt.Fprintln(writer, memo[0]) // 최상위 괄호열의 값 출력
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	return mustParseInt(scanString())
}
