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

	str      string
	alphabet = [26]int{}
)

// 난이도: Silver 3
// 메모리: 920KB
// 시간: 8ms
// 분류: 구현, 문자열, 그리디 알고리즘
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	str = scanString()
	for _, c := range str {
		alphabet[c-'A']++
	}
}

func Solve() {
	leftBuilder := strings.Builder{}

	for i := 0; i < 26; i++ {
		for alphabet[i] > 1 {
			leftBuilder.WriteByte(byte(i + 'A'))
			alphabet[i] -= 2
		}
	}

	remain := 0
	for i := 0; i < 26; i++ {
		remain += alphabet[i]
	}

	if remain > 1 {
		fmt.Println("I'm Sorry Hansoo")
		return
	}

	left := leftBuilder.String()

	rightBuilder := strings.Builder{}

	for i := len(left) - 1; i >= 0; i-- {
		rightBuilder.WriteByte(left[i])
	}

	if remain == 1 {
		for i := 0; i < 26; i++ {
			if alphabet[i] == 1 {
				leftBuilder.WriteByte(byte(i + 'A'))
				break
			}
		}
	}

	fmt.Fprintf(writer, "%s%s\n", leftBuilder.String(), rightBuilder.String())
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
