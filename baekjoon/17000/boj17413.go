package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	S       string
)

// 17413번: 단어 뒤집기 2
// 난이도: 실버 3
// 메모리: 2236 KB
// 시간: 8 ms
// 분류: 구현, 자료 구조, 문자열, 스택
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanLines)
	scanner.Buffer(make([]byte, 200000, 200000), 200000)

	Setup()
	Solve()
}

func Setup() {
	S = scanString()
}

func Solve() {
	builder := strings.Builder{}
	stack := []byte{}
	isTag := false

	for i := 0; i < len(S); i++ {
		switch S[i] {
		case '<':
			isTag = true
			for len(stack) > 0 {
				builder.WriteByte(stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, S[i])
		case '>':
			for len(stack) > 0 {
				builder.WriteByte(stack[0])
				stack = stack[1:]
			}
			builder.WriteByte('>')
			isTag = false
		case ' ':
			if isTag {
				stack = append(stack, S[i])
			} else {
				for len(stack) > 0 {
					builder.WriteByte(stack[len(stack)-1])
					stack = stack[:len(stack)-1]
				}
				builder.WriteByte(' ')
			}
		default:
			stack = append(stack, S[i])
		}
	}

	for len(stack) > 0 {
		builder.WriteByte(stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	fmt.Fprintln(writer, builder.String())
}

func scanBytes() []byte {
	scanner.Scan()
	return scanner.Bytes()
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
