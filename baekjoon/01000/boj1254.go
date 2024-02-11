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
	S       string
)

// 1254번: 팰린드롬 만들기
// https://www.acmicpc.net/problem/1254
// 난이도: 실버 2
// 메모리: 860 KB
// 시간: 4 ms
// 분류: 브루트포스 알고리즘
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
	minLength := len(S) * 2

	for i := len(S) - 1; i >= 0; i-- {
		if isPalindrome(S[i:]) {
			minLength = len(S)*2 - (len(S) - i)
		}
	}

	fmt.Fprintln(writer, minLength)
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
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

// func readNums(b []byte) (res []int) {
// 	for i := 0; i < len(b); i++ {
// 		n := 0
// 		for ; i < len(b) && b[i] >= '0' && b[i] <= '9'; i++ {
// 			n = n*10 + int(b[i]-'0')
// 		}
// 		res = append(res, n)
// 	}
// 	return
// }
