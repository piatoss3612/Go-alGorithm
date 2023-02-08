package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	input   []byte
)

// 난이도: Gold 5
// 메모리: 1092KB
// 시간: 4ms
// 분류: 스택, 문자열, 파싱
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	input = scanBytes()
}

func Solve() {
	stack := []byte{}

	for len(input) > 0 {
		c := input[0]
		input = input[1:]

		switch c {
		case '{':
			fmt.Fprintln(writer, strings.Repeat(" ", len(stack)*2)+string(c)) // 스택에 남아있는 '{'의 개수 * 2만큼 공백 + '{' 출력 및 줄바꿈
			stack = append(stack, c)                                          // 스택에 '{' 추가

		case '}':
			// '{'와 '}' 사이의 문자열 탐색
			var s string
			for {
				n := len(stack)
				back := stack[n-1]
				if back == '{' {
					break
				} else {
					s = string(back) + s
					stack = stack[:n-1]
				}
			}

			// '{'와 '}' 사이에 문자열이 있는 경우
			if len(s) > 0 {
				fmt.Fprintln(writer, strings.Repeat(" ", len(stack)*2)+s) // '{'의 개수 * 2만큼의 공백 + 문자열 s 출력
			}

			stack = stack[:len(stack)-1] // '}'와 매칭되는 '{'를 스택에서 제거
			stack = append(stack, c)     // 스택에 '}' 추가

		case ',':
			// '{'와 ',' 사이의 문자열 탐색 ('}' 포함)
			var s string
			for {
				n := len(stack)
				back := stack[n-1]
				if back == '{' {
					break
				} else {
					s = string(back) + s
					stack = stack[:n-1]
				}
			}
			fmt.Fprint(writer, strings.Repeat(" ", len(stack)*2)+s) // '{'의 개수 * 2만큼의 공백 + 문자열 s 출력
			fmt.Fprint(writer, string(c))                           // ','출력
			fmt.Fprintln(writer)                                    // 줄바꿈

		default:
			stack = append(stack, c) // 스택에 c 추가
		}
	}
	fmt.Fprintln(writer, string(stack[0])) // 스택에 남아있는 '}' 출력
}

func scanBytes() []byte {
	scanner.Scan()
	return scanner.Bytes()
}
