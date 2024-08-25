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
	T       int
	arr     []string
)

// 17609번: 회문
// https://www.acmicpc.net/problem/17609
// 난이도: 골드 5
// 메모리: 11032 KB
// 시간: 44 ms
// 분류: 문자열, 두 포인터
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, 200000), 200000)

	Setup()
	Solve()
}

func Setup() {
	T = scanInt()
	arr = make([]string, T)
	for i := 0; i < T; i++ {
		arr[i] = scanString()
	}
}

func Solve() {
	for i := 0; i < T; i++ {
		fmt.Fprintln(writer, isPalindrome(i, 0, len(arr[i])-1, false))
	}
}

func isPalindrome(idx, start, end int, skipped bool) int {
	if start >= end {
		if skipped {
			return 1
		}
		return 0
	}

	if arr[idx][start] == arr[idx][end] {
		return isPalindrome(idx, start+1, end-1, skipped)
	}

	if skipped {
		return 2
	}

	return min(isPalindrome(idx, start+1, end, true), isPalindrome(idx, start, end-1, true))
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
