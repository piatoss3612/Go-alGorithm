package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	N       int
)

// 1013번: Contact
// hhttps://www.acmicpc.net/problem/1013
// 난이도: 골드 5
// 메모리: 6144 KB
// 시간: 20 ms
// 분류: 문자열, 정규 표현식
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
}

func Solve() {
	for i := 0; i < N; i++ {
		if IsContact(scanString()) {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
}

func IsContact(s string) bool {
	exp := "^(100+1+|01)+$"
	matched, _ := regexp.MatchString(exp, s)
	return matched
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
