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
	N, K    int
)

// 1052번: 물병
// hhttps://www.acmicpc.net/problem/1052
// 난이도: 골드 5
// 메모리: 860 KB
// 시간: 172 ms
// 분류: 그리디 알고리즘, 수학, 비트마스킹
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, K = scanInt(), scanInt()
}

func Solve() {
	var ans int
	for {
		// N의 이진수 표현에서 1의 개수가 K보다 작거나 같으면 종료
		// == 물병을 K개 이하로 가질 수 있는 최소의 N을 찾음
		if countOneBits(N) <= K {
			break
		}

		N++ // 물병을 하나 구매하고 N을 증가시킴
		ans++
	}
	fmt.Fprintln(writer, ans)
}

func countOneBits(n int) int {
	var cnt int
	for n > 0 {
		if n&1 == 1 {
			cnt++
		}
		n >>= 1
	}
	return cnt
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
