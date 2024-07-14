package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	N       int
	answer  [][2]int
)

// 1914번: 하노이 탑
// hhttps://www.acmicpc.net/problem/1914
// 난이도: 골드 5
// 메모리: 57856 KB
// 시간: 220 ms
// 분류: 재귀, 임의 정밀도/큰 수 연산
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
	// 원판이 20개 이하일 때만 실행
	if N <= 20 {
		Hanoi(N, 1, 3, 2)
	}

	// 2^N - 1
	cnt := big.NewInt(0).Sub(big.NewInt(1).Lsh(big.NewInt(1), uint(N)), big.NewInt(1))

	fmt.Fprintln(writer, cnt)
	for _, v := range answer {
		fmt.Fprintln(writer, v[0], v[1])
	}
}

func Hanoi(n, from, to, via int) {
	// 원판이 1개일 때 from -> to 로 이동
	if n == 1 {
		answer = append(answer, [2]int{from, to})
		return
	}

	Hanoi(n-1, from, via, to)                 // n-1개를 from에서 to를 거쳐 via로 이동
	answer = append(answer, [2]int{from, to}) // n번째 원판(가장 큰 원판)을 from에서 to로 이동
	Hanoi(n-1, via, to, from)                 // n-1개를 via에서 from을 거쳐 to로 이동
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
