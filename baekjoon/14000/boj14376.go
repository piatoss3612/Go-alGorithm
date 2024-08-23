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
)

// 14376번: The Last Word (Large)
// https://www.acmicpc.net/problem/14376
// 난이도: 실버 1
// 메모리: 4312 KB
// 시간: 8 ms
// 분류: 그리디 알고리즘, 문자열
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	T = scanInt()
}

func Solve() {
	for t := 1; t <= T; t++ {
		S := scanString()
		lastWord := makeLastWord(S)
		fmt.Fprintf(writer, "Case #%d: %s\n", t, lastWord)
	}
}

func makeLastWord(S string) string {
	b := []byte(S)
	lastWord := make([]byte, 0, len(b))
	for _, c := range b {
		if len(lastWord) == 0 {
			lastWord = append(lastWord, c)
			continue
		}

		if lastWord[0] > c {
			lastWord = append(lastWord, c)
		} else {
			lastWord = append([]byte{c}, lastWord...)
		}
	}

	return string(lastWord)
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
