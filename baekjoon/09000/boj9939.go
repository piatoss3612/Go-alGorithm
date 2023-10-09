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
	MaxBuf  = 1200000 // 스캐너 버퍼 크기
)

// 메모리: 7060KB
// 시간: 24ms
// 첫번째 문자열의 길이가 최대 1백만인 것에 주의합시다...
func main() {
	defer writer.Flush()
	scanner.Buffer(make([]byte, 0, MaxBuf), MaxBuf)
	scanner.Split(bufio.ScanWords)
	b1 := scanBytes() // 첫번째 문자열 바이트 슬라이스
	b2 := scanBytes() // 두번째 문자열 바이트 슬라이스
	b2Len := len(b2)  // 두번째 문자열 길이

	var stack []byte

	for i := 0; i < len(b1); i++ {
		stack = append(stack, b1[i]) // 스택에 첫번째 문자열의 1번째 문자 추가

		// 방금 스택에 추가된 첫번째 문자열의 i번째 문자와 두번째 문자열의 마지막 문자가 같으면서
		// 동시에 stack의 길이가 두번째 문자열 길이보다 길거나 같은 경우
		if b1[i] == b2[b2Len-1] && len(stack) >= len(b2) {
			if same(stack[len(stack)-b2Len:], b2, b2Len) { // 같은 문자열인지 체크
				stack = stack[:len(stack)-b2Len] // 같은 문자열이면 스택에서 제거
			}
		}
	}

	if len(stack) == 0 {
		fmt.Fprintln(writer, "FRULA")
	} else {
		fmt.Fprintln(writer, string(stack))
	}
}

func same(a, b []byte, length int) bool {
	for i := 0; i < length; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func scanBytes() []byte {
	scanner.Scan()
	return scanner.Bytes()
}
