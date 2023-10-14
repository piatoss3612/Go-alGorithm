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

// 9417번: 최대 GCD
// https://www.acmicpc.net/problem/9417
// 난이도: 실버 4
// 메모리: 860 KB
// 시간: 4 ms
// 분류: 수학, 정수론, 유클리드 호제법, 브루트포스 알고리즘
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanLines)

	Setup()
	Solve()
}

func Setup() {
	T = scanInt()
}

func Solve() {
	for i := 0; i < T; i++ {
		b := scanBytes()
		var arr []int

		n := 0
		for i := 0; i < len(b); i++ {
			if b[i] == ' ' {
				arr = append(arr, n)
				n = 0
				continue
			}

			n = n*10 + int(b[i]-'0')
		}

		arr = append(arr, n)

		max := 0

		for a := 0; a < len(arr)-1; a++ {
			for b := a + 1; b < len(arr); b++ {
				g := gcd(arr[a], arr[b])
				if g > max {
					max = g
				}
			}
		}

		fmt.Fprintln(writer, max)
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
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
