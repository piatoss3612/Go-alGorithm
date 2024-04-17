package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	scanner  = bufio.NewScanner(os.Stdin)
	writer   = bufio.NewWriter(os.Stdout)
	N, M     int
	alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// 17828번: 문자열 화폐
// hhttps://www.acmicpc.net/problem/17828
// 난이도: 골드 5
// 메모리: 31540 KB
// 시간: 108 ms
// 분류: 그리디 알고리즘, 문자열
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
}

func Solve() {
	builder := strings.Builder{}

	n, m := N, M

	// 뒤에서부터 가장 큰 가치의 문자를 채워넣으면서 앞으로 이동
	for n > 0 && m > 0 {
		x := m - (n - 1)
		if x > 0 {
			if x > 26 {
				x = 26
			}
			builder.WriteByte(alphabet[x-1])
			m -= x
		} else {
			builder.WriteByte(alphabet[0])
			m--
		}
		n--
	}

	s := []byte(builder.String())

	if len(s) != N || m > 0 {
		fmt.Fprintln(writer, "!")
		return
	}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	fmt.Fprintln(writer, string(s))
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
