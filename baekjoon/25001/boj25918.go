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

	N int
	S string
)

// 난이도: Silver 1
// 메모리: 1372KB
// 시간: 8ms
// 분류: 스택
func main() {
	defer writer.Flush()
	scanner.Buffer(make([]byte, 0, 300000), 300000) // 문자열 길이가 최대 20만이므로 버퍼 크기를 충분히 크게 잡아야 함
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	S = scanString()
}

func Solve() {
	stack := make([]byte, 0, N)
	depth := 0 // 괄호의 깊이
	day := 0 // 문자열 S를 만드는데 필요한 최소 일수

	for _, c := range S {
		switch c {
		case '(':
			if len(stack) == 0 || stack[len(stack)-1] == '(' {
				stack = append(stack, byte(c))
				depth += 1
				day = max(day, depth)
				continue
			}

			// stack[len(stack)-1] == ')'
			stack = stack[:len(stack)-1]
			depth -= 1
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] == ')' {
				stack = append(stack, byte(c))
				depth += 1
				day = max(day, depth)
				continue
			}

			// stack[len(stack)-1] == '('
			stack = stack[:len(stack)-1]
			depth -= 1
		}
	}

	if len(stack) != 0 {
		fmt.Fprintln(writer, -1)
	} else {
		fmt.Fprintln(writer, day)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
