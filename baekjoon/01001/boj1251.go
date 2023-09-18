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

	S   []byte
	ans string
)

// 난이도: Silver 5
// 메모리: 1216KB
// 시간: 4ms
// 분류: 구현, 문자열, 브루트포스 알고리즘
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	S = scanBytes()
}

func Solve() {
	bruteForce()
	fmt.Fprintln(writer, ans)
}

func bruteForce() {
	n := len(S)
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			first := make([]byte, i+1)
			second := make([]byte, j-i)
			third := make([]byte, n-j-1)

			copy(first, S[0:i+1])
			copy(second, S[i+1:j+1])
			copy(third, S[j+1:n])

			reverseBytes(first)
			reverseBytes(second)
			reverseBytes(third)

			reversed := fmt.Sprintf("%s%s%s", string(first), string(second), string(third))
			if ans == "" || strings.Compare(ans, reversed) > 0 {
				ans = reversed
			}
		}
	}
}

func reverseBytes(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
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