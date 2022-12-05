package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	inp     []byte
	N       int
)

const INPUT_SIZE = 600000 // 입력값의 길이가 최대 50만인데 조금 넉넉하게 잡았다

// 난이도: Gold 5
// 메모리: 1420KB
// 시간: 12ms
// 분류: 문자열, 애드 혹
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, 0, INPUT_SIZE), INPUT_SIZE) // 이거 때매 무수한 틀렸습니다 세례를...
	Input()
	Solve()
}

func Input() {
	inp = scanBytes()
	N = len(inp)
}

func Solve() {
	isPalindrome, allSame := true, true
	for i := 0; i < N/2; i++ {
		if inp[i] != inp[N-1-i] {
			isPalindrome = false
			break
		}

		if inp[i] != inp[i+1] {
			allSame = false
		}
	}

	if !isPalindrome {
		// 1. 문자열이 회문이 아닌 경우
		// 문자열 전체가 회문이 아닌 가장 긴 부분문자열
		fmt.Fprintln(writer, N)
	} else if !allSame {
		// 2. 문자열이 회문이면서 모든 문자가 같은 경우
		fmt.Fprintln(writer, N-1)
	} else {
		// 3. 문자열이 회문이면서 모든 문자가 같지 않은 경우
		// 왼쪽 끝이든 오른쪽 끝이든 하나의 문자를 제거하면 가장 긴 부분문자열을 구할 수 있다
		fmt.Fprintln(writer, -1)
	}
}

func scanBytes() []byte {
	scanner.Scan()
	return scanner.Bytes()
}
