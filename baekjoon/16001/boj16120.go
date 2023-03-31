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
	input   []byte
)

// 난이도: Gold 4
// 메모리: 2796KB
// 시간: 148ms
// 분류: 문자열, 큐
// 풀이: 스택을 사용해서 PPAP 문자열을 하나의 P로 압축시켰다.
// 틀린 이유: PPAP 문자열을 여러 개 연결한 것도 PPAP 문자열이라고 생각해서 틀렸다. 여느 때처럼 문제를 주의 깊게 읽지 않아서 생긴 실수다.
// 주의점: 입력이 매우 기므로 buffer의 크기를 적절하게 늘려줘야 한다.
func main() {
	defer writer.Flush()
	scanner.Buffer(make([]byte, 0, 2000000), 2000000)
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	input = scanBytes()
}

func Solve() {
	stack := make([]byte, 0, 2000000)

	for len(input) > 0 {
		c := input[0]
		input = input[1:]
		stack = append(stack, c)

		// PPAP를 발견하면 스택에서 제거하고 P 하나로 변환한다
		n := len(stack)
		if n >= 4 {
			if string(stack[n-4:n]) == "PPAP" {
				stack = stack[:n-4]
				stack = append(stack, 'P')
			}
		}
	}

	// 스택에 P 하나만 남아있으면 PPAP 문자열이다
	if len(stack) == 1 && stack[0] == 'P' {
		fmt.Fprintln(writer, "PPAP")
	} else {
		fmt.Fprintln(writer, "NP")
	}
}

func scanBytes() []byte {
	scanner.Scan()
	return scanner.Bytes()
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
