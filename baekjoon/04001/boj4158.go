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
	N, M    int
	dup     map[int]int
)

// 난이도: Silver 5
// 메모리: 232756KB
// 시간: 1380ms
// 분류: 자료 구조, 해시를 사용한 집합과 맵
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
}

func Setup() {
	for {
		N, M = scanInt(), scanInt()
		if N == 0 && M == 0 {
			return
		}
		dup = make(map[int]int)
		Solve()
	}
}

func Solve() {
	cnt := 0
	for i := 1; i <= N; i++ {
		x := scanInt()
		dup[x]++
	}

	for i := 1; i <= M; i++ {
		x := scanInt()
		if dup[x] > 0 {
			cnt++
		}
	}

	fmt.Fprintln(writer, cnt)
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
