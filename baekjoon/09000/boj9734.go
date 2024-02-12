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
	S       string
)

// 9734번: 순환 소수
// https://www.acmicpc.net/problem/9734
// 난이도: 실버 1
// 메모리: 868 KB
// 시간: 4 ms
// 분류: 수학, 정수론, 유클리드 호제법
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
}

func Setup() {
	for scanner.Scan() {
		S = scanner.Text()
		Solve()
	}
}

func Solve() {
	parts := strings.Split(S, ".")
	integer := parts[0]
	decimal := parts[1]

	repeatingStart := strings.Index(decimal, "(")
	decimalBeforeRepeating := decimal[:repeatingStart]
	repeatingDecimal := decimal[repeatingStart+1 : len(decimal)-1]

	d1, d2 := 1, 1
	a, b := mustParseInt(integer+decimalBeforeRepeating+repeatingDecimal), mustParseInt(integer+decimalBeforeRepeating)

	for i := 0; i < len(decimalBeforeRepeating+repeatingDecimal); i++ {
		d1 *= 10
	}

	for i := 0; i < len(decimalBeforeRepeating); i++ {
		d2 *= 10
	}

	n, d := a-b, d1-d2

	g := gcd(n, d)

	fmt.Fprintf(writer, "%s = %d / %d\n", S, n/g, d/g)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
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
